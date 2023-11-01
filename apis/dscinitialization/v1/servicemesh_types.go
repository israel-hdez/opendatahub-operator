package v1

// ServiceMeshSpec configures Service Mesh.
type ServiceMeshSpec struct {
	// Mesh holds configuration of Service Mesh used by Opendatahub.
	Mesh MeshSpec `json:"mesh,omitempty"`
}

type MeshSpec struct {
	// Name is a name Service Mesh Control Plane. Defaults to "minimal".
	// +kubebuilder:default=minimal
	Name string `json:"name,omitempty"`
	// Namespace is a namespace where Service Mesh is deployed. Defaults to "istio-system".
	// +kubebuilder:default=istio-system
	Namespace string `json:"namespace,omitempty"`
}
