package v1

import operatorv1 "github.com/openshift/api/operator/v1"

// ServerlessSpec configures Service Mesh.
type ServerlessSpec struct {
	// +kubebuilder:validation:Enum=Managed;Removed
	// +kubebuilder:default=Removed
	ManagementState operatorv1.ManagementState `json:"managementState,omitempty"`

	Serving Serving `json:"serving,omitempty"`
}

type Serving struct {
	// +kubebuilder:default=knative-serving
	Name string `json:"name,omitempty"`
	// +kubebuilder:default=knative-serving
	Namespace string `json:"namespace,omitempty"`
	// +kubebuilder:default=knative-local-gateway
	LocalGatewayServiceName string         `json:"localGatewayServiceName,omitempty"`
	IngressGateway          IngressGateway `json:"ingressGateway,omitempty"`
}

type IngressGateway struct {
	// GatewaySelector map<string, string> `json:"selector,omitempty"`
	// Domain string `json:"domain,omitempty"`

	Certificate Certificate `json:"certificate,omitempty"`
}

type Certificate struct {
	// +kubebuilder:default=knative-serving-cert
	SecretName string `json:"secretName,omitempty"`
	// +kubebuilder:default=true
	Generate bool `json:"generate,omitempty"`
}
