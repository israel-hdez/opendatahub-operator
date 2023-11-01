package serverless

import (
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/feature/servicemesh"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"path"
	"path/filepath"

	ctrlLog "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/feature"
)

var log = ctrlLog.Log.WithName("features")

func ConfigureServerlessFeatures(s *feature.FeaturesInitializer) error {
	var rootDir = filepath.Join( /*feature.BaseOutputDir*/ "/tmp/odh-operator", s.DSCInitializationSpec.ApplicationsNamespace)
	if err := feature.CopyEmbeddedFiles("templates/serverless", rootDir); err != nil {
		return err
	}

	serverlessSpec := s.Serverless

	servingDeployment, err := feature.CreateFeature("serverless-serving-deployment").
		For(s.DSCInitializationSpec).
		Manifests(
			path.Join(rootDir, "templates/serverless/serving-install"),
		).
		PreConditions(
			EnsureServerlessOperatorInstalled,
			EnsureServerlessAbsent,
			servicemesh.EnsureServiceMeshInstalled,
			feature.CreateNamespace(serverlessSpec.Serving.Namespace),
		).
		PostConditions(
			feature.WaitForPodsToBeReady(serverlessSpec.Serving.Namespace),
		).
		Load()
	if err != nil {
		return err
	}
	s.Features = append(s.Features, servingDeployment)

	servingIstioGateways, err := feature.CreateFeature("serverless-serving-gateways").
		For(s.DSCInitializationSpec).
		PreConditions(
			// Check serverless is installed
			feature.WaitForResourceToBeCreated(serverlessSpec.Serving.Namespace, schema.GroupVersionResource{
				Group:    "operator.knative.dev",
				Version:  "v1beta1",
				Resource: "knativeservings",
			}),
		).
		Manifests(
			path.Join(rootDir, "templates/serverless/serving-istio-gateways"),
		).
		Load()
	if err != nil {
		return err
	}
	s.Features = append(s.Features, servingIstioGateways)

	return nil
}

// TODO
func EnsureServerlessAbsent(f *feature.Feature) error {
	return nil
}

func EnsureServerlessOperatorInstalled(f *feature.Feature) error {
	if err := feature.EnsureCRDIsInstalled("knativeservings.operator.knative.dev")(f); err != nil {
		log.Info("Failed to find the pre-requisite KNative Serving Operator CRD, please ensure Serverless Operator is installed.", "feature", f.Name)

		return err
	}

	return nil
}
