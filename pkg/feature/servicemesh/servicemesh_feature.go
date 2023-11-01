package servicemesh

import (
	"path"
	"path/filepath"

	ctrlLog "sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/feature"
)

var log = ctrlLog.Log.WithName("features")

func ConfigureServiceMeshFeatures(s *feature.FeaturesInitializer) error {
	var rootDir = filepath.Join("/tmp/odh-operator", s.DSCInitializationSpec.ApplicationsNamespace)
	if err := feature.CopyEmbeddedFiles("templates/servicemesh", rootDir); err != nil {
		return err
	}

	serviceMeshSpec := s.ServiceMesh

	if oauth, err := feature.CreateFeature("control-plane-creation").
		For(s.DSCInitializationSpec).
		Manifests(
			path.Join(rootDir, "templates/servicemesh/base"),
		).
		PreConditions(
			EnsureServiceMeshOperatorInstalled,
			feature.CreateNamespace(serviceMeshSpec.Mesh.Namespace),
		).
		PostConditions(
			feature.WaitForPodsToBeReady(serviceMeshSpec.Mesh.Namespace),
		).
		Load(); err != nil {
		return err
	} else {
		s.Features = append(s.Features, oauth)
	}

	return nil
}

func EnsureServiceMeshOperatorInstalled(f *feature.Feature) error {
	if err := feature.EnsureCRDIsInstalled("servicemeshcontrolplanes.maistra.io")(f); err != nil {
		log.Info("Failed to find the pre-requisite Service Mesh Control Plane CRD, please ensure Service Mesh Operator is installed.", "feature", f.Name)

		return err
	}

	return nil
}
