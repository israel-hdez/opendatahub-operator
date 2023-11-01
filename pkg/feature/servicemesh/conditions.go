package servicemesh

import (
	"context"
	"github.com/hashicorp/go-multierror"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/feature"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/dynamic"
	"time"
)

const (
	interval = 2 * time.Second
	duration = 5 * time.Minute
)

func EnsureServiceMeshInstalled(f *feature.Feature) error {
	if err := feature.EnsureCRDIsInstalled("servicemeshcontrolplanes.maistra.io")(f); err != nil {
		log.Info("Failed to find the pre-requisite Service Mesh Control Plane CRD, please ensure Service Mesh Operator is installed.", "feature", f.Name)

		return err
	}

	smcp := f.Spec.Mesh.Name
	smcpNs := f.Spec.Mesh.Namespace

	if err := WaitForControlPlaneToBeReady(f); err != nil {
		log.Error(err, "failed waiting for control plane being ready", "feature", f.Name, "control-plane", smcp, "namespace", smcpNs)

		return multierror.Append(err, errors.New("service mesh control plane is not ready")).ErrorOrNil()
	}

	return nil
}

func WaitForControlPlaneToBeReady(feature *feature.Feature) error {
	return wait.PollUntilContextTimeout(context.TODO(), interval, duration, false, func(ctx context.Context) (bool, error) {
		smcp := feature.Spec.Mesh.Name
		smcpNs := feature.Spec.Mesh.Namespace

		log.Info("waiting for control plane components to be ready", "feature", feature.Name, "control-plane", smcp, "namespace", smcpNs, "duration (s)", duration.Seconds())
		ready, err := CheckControlPlaneComponentReadiness(feature.DynamicClient, smcp, smcpNs)

		if ready {
			log.Info("done waiting for control plane components to be ready", "feature", feature.Name, "control-plane", smcp, "namespace", smcpNs)
		}

		return ready, err
	})
}

func CheckControlPlaneComponentReadiness(dynamicClient dynamic.Interface, smcp, smcpNs string) (bool, error) {
	smcpgvr := schema.GroupVersionResource{
		Group:    "maistra.io",
		Version:  "v2",
		Resource: "servicemeshcontrolplanes",
	}

	unstructObj, err := dynamicClient.Resource(smcpgvr).Namespace(smcpNs).Get(context.TODO(), smcp, metav1.GetOptions{})
	if err != nil {
		log.Info("failed to find Service Mesh Control Plane", "control-plane", smcp, "namespace", smcpNs)

		return false, err
	}

	components, found, err := unstructured.NestedMap(unstructObj.Object, "status", "readiness", "components")
	if err != nil || !found {
		log.Info("status conditions not found or error in parsing of Service Mesh Control Plane")

		return false, err
	}

	readyComponents := len(components["ready"].([]interface{}))
	pendingComponents := len(components["pending"].([]interface{}))
	unreadyComponents := len(components["unready"].([]interface{}))

	return pendingComponents == 0 && unreadyComponents == 0 && readyComponents > 0, nil
}
