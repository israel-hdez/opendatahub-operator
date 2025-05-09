// Package modelmeshserving provides utility functions to config MoModelMesh, a general-purpose model serving management/routing layer
// +groupName=datasciencecluster.opendatahub.io
package modelmeshserving

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	operatorv1 "github.com/openshift/api/operator/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	dsciv1 "github.com/opendatahub-io/opendatahub-operator/v2/apis/dscinitialization/v1"
	"github.com/opendatahub-io/opendatahub-operator/v2/components"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/cluster"
	"github.com/opendatahub-io/opendatahub-operator/v2/pkg/deploy"
)

var (
	ComponentName          = "model-mesh"
	Path                   = deploy.DefaultManifestPath + "/" + ComponentName + "/overlays/odh"
	DependentComponentName = "odh-model-controller"
	DependentPath          = deploy.DefaultManifestPath + "/" + DependentComponentName + "/base"
)

// Verifies that Dashboard implements ComponentInterface.
var _ components.ComponentInterface = (*ModelMeshServing)(nil)

// ModelMeshServing struct holds the configuration for the ModelMeshServing component.
// +kubebuilder:object:generate=true
type ModelMeshServing struct {
	components.Component `json:""`
}

func (m *ModelMeshServing) Init(ctx context.Context, _ cluster.Platform) error {
	log := logf.FromContext(ctx).WithName(ComponentName)

	var imageParamMap = map[string]string{
		"odh-mm-rest-proxy":             "RELATED_IMAGE_ODH_MM_REST_PROXY_IMAGE",
		"odh-modelmesh-runtime-adapter": "RELATED_IMAGE_ODH_MODELMESH_RUNTIME_ADAPTER_IMAGE",
		"odh-modelmesh":                 "RELATED_IMAGE_ODH_MODELMESH_IMAGE",
		"odh-modelmesh-controller":      "RELATED_IMAGE_ODH_MODELMESH_CONTROLLER_IMAGE",
	}

	// odh-model-controller to use
	var dependentImageParamMap = map[string]string{
		"odh-model-controller": "RELATED_IMAGE_ODH_MODEL_CONTROLLER_IMAGE",
	}

	// Update image parameters
	if err := deploy.ApplyParams(Path, imageParamMap); err != nil {
		log.Error(err, "failed to update image", "path", Path)
	}

	// Update image parameters for odh-model-controller
	if err := deploy.ApplyParams(DependentPath, dependentImageParamMap); err != nil {
		log.Error(err, "failed to update image", "path", DependentPath)
	}

	return nil
}

func (m *ModelMeshServing) OverrideManifests(ctx context.Context, _ cluster.Platform) error {
	// Go through each manifest and set the overlays if defined
	for _, subcomponent := range m.DevFlags.Manifests {
		if strings.Contains(subcomponent.URI, DependentComponentName) {
			// Download subcomponent
			if err := deploy.DownloadManifests(ctx, DependentComponentName, subcomponent); err != nil {
				return err
			}
			// If overlay is defined, update paths
			defaultKustomizePath := "base"
			if subcomponent.SourcePath != "" {
				defaultKustomizePath = subcomponent.SourcePath
			}
			DependentPath = filepath.Join(deploy.DefaultManifestPath, DependentComponentName, defaultKustomizePath)
		}

		if strings.Contains(subcomponent.URI, ComponentName) {
			// Download subcomponent
			if err := deploy.DownloadManifests(ctx, ComponentName, subcomponent); err != nil {
				return err
			}
			// If overlay is defined, update paths
			defaultKustomizePath := "overlays/odh"
			if subcomponent.SourcePath != "" {
				defaultKustomizePath = subcomponent.SourcePath
			}
			Path = filepath.Join(deploy.DefaultManifestPath, ComponentName, defaultKustomizePath)
		}
	}
	return nil
}

func (m *ModelMeshServing) GetComponentName() string {
	return ComponentName
}

func (m *ModelMeshServing) ReconcileComponent(ctx context.Context,
	cli client.Client,
	owner metav1.Object,
	dscispec *dsciv1.DSCInitializationSpec,
	platform cluster.Platform,
	_ bool,
) error {
	l := logf.FromContext(ctx)
	enabled := m.GetManagementState() == operatorv1.Managed
	monitoringEnabled := dscispec.Monitoring.ManagementState == operatorv1.Managed

	// Update Default rolebinding
	if enabled {
		if m.DevFlags != nil {
			// Download manifests and update paths
			if err := m.OverrideManifests(ctx, platform); err != nil {
				return err
			}
		}

		if err := cluster.UpdatePodSecurityRolebinding(ctx, cli, dscispec.ApplicationsNamespace,
			"modelmesh",
			"modelmesh-controller"); err != nil {
			return err
		}
	}

	extraParamsMap := map[string]string{
		"nim-state": getNimManagementFlag(owner),
	}
	if err := deploy.ApplyParams(DependentPath, nil, extraParamsMap); err != nil {
		return fmt.Errorf("failed to update image from %s : %w", Path, err)
	}

	if err := deploy.DeployManifestsFromPath(ctx, cli, owner, Path, dscispec.ApplicationsNamespace, ComponentName, enabled); err != nil {
		return fmt.Errorf("failed to apply manifests from %s : %w", Path, err)
	}
	l.WithValues("Path", Path).Info("apply manifests done for modelmesh")
	// For odh-model-controller
	if enabled {
		if err := cluster.UpdatePodSecurityRolebinding(ctx, cli, dscispec.ApplicationsNamespace,
			"odh-model-controller"); err != nil {
			return err
		}
	}
	if err := deploy.DeployManifestsFromPath(ctx, cli, owner, DependentPath, dscispec.ApplicationsNamespace, m.GetComponentName(), enabled); err != nil {
		// explicitly ignore error if error contains keywords "spec.selector" and "field is immutable" and return all other error.
		if !strings.Contains(err.Error(), "spec.selector") || !strings.Contains(err.Error(), "field is immutable") {
			return err
		}
	}

	l.WithValues("Path", DependentPath).Info("apply manifests done for odh-model-controller")

	if enabled {
		if err := cluster.WaitForDeploymentAvailable(ctx, cli, ComponentName, dscispec.ApplicationsNamespace, 20, 2); err != nil {
			return fmt.Errorf("deployment for %s is not ready to server: %w", ComponentName, err)
		}
	}

	// CloudService Monitoring handling
	if platform == cluster.ManagedRhoai {
		// first model-mesh rules
		if err := m.UpdatePrometheusConfig(cli, l, enabled && monitoringEnabled, ComponentName); err != nil {
			return err
		}
		// then odh-model-controller rules
		if err := m.UpdatePrometheusConfig(cli, l, enabled && monitoringEnabled, DependentComponentName); err != nil {
			return err
		}
		if err := deploy.DeployManifestsFromPath(ctx, cli, owner,
			filepath.Join(deploy.DefaultManifestPath, "monitoring", "prometheus", "apps"),
			dscispec.Monitoring.Namespace,
			"prometheus", true); err != nil {
			return err
		}
		l.Info("updating SRE monitoring done")
	}

	return nil
}

func getNimManagementFlag(obj metav1.Object) string {
	removed := string(operatorv1.Removed)
	un, err := runtime.DefaultUnstructuredConverter.ToUnstructured(obj)
	if err != nil {
		return removed
	}
	kserve, foundKserve, _ := unstructured.NestedString(un, "spec", "components", "kserve", "managementState")
	if foundKserve && kserve != removed {
		nim, foundNim, _ := unstructured.NestedString(un, "spec", "components", "kserve", "nim", "managementState")
		if foundNim {
			return nim
		}
	}
	return removed
}
