/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/apimachinery/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kmapi "kmodules.xyz/client-go/api/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type Connection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ConnectionSpec   `json:"spec,omitempty"`
	Status            ConnectionStatus `json:"status,omitempty"`
}

type ConnectionSpec struct {
	ConnectionSpec2 `json:",inline"`
	// +optional
	KubeformOutput ConnectionSpec2 `json:"kubeformOutput,omitempty" tf:"-"`
}

type ConnectionSpecRoutes struct {
	// +optional
	DestinationCIDRBlock *string `json:"destinationCIDRBlock,omitempty" tf:"destination_cidr_block"`
	// +optional
	Source *string `json:"source,omitempty" tf:"source"`
	// +optional
	State *string `json:"state,omitempty" tf:"state"`
}

type ConnectionSpecVgwTelemetry struct {
	// +optional
	AcceptedRouteCount *int64 `json:"acceptedRouteCount,omitempty" tf:"accepted_route_count"`
	// +optional
	LastStatusChange *string `json:"lastStatusChange,omitempty" tf:"last_status_change"`
	// +optional
	OutsideIPAddress *string `json:"outsideIPAddress,omitempty" tf:"outside_ip_address"`
	// +optional
	Status *string `json:"status,omitempty" tf:"status"`
	// +optional
	StatusMessage *string `json:"statusMessage,omitempty" tf:"status_message"`
}

type ConnectionSpec2 struct {
	TerminationPolicy base.TerminationPolicy `json:"terminationPolicy,omitempty" tf:"-"`

	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Arn *string `json:"arn,omitempty" tf:"arn"`
	// +optional
	CustomerGatewayConfiguration *string `json:"customerGatewayConfiguration,omitempty" tf:"customer_gateway_configuration"`
	CustomerGatewayID            *string `json:"customerGatewayID" tf:"customer_gateway_id"`
	// +optional
	EnableAcceleration *bool `json:"enableAcceleration,omitempty" tf:"enable_acceleration"`
	// +optional
	LocalIpv4NetworkCIDR *string `json:"localIpv4NetworkCIDR,omitempty" tf:"local_ipv4_network_cidr"`
	// +optional
	LocalIpv6NetworkCIDR *string `json:"localIpv6NetworkCIDR,omitempty" tf:"local_ipv6_network_cidr"`
	// +optional
	RemoteIpv4NetworkCIDR *string `json:"remoteIpv4NetworkCIDR,omitempty" tf:"remote_ipv4_network_cidr"`
	// +optional
	RemoteIpv6NetworkCIDR *string `json:"remoteIpv6NetworkCIDR,omitempty" tf:"remote_ipv6_network_cidr"`
	// +optional
	Routes []ConnectionSpecRoutes `json:"routes,omitempty" tf:"routes"`
	// +optional
	StaticRoutesOnly *bool `json:"staticRoutesOnly,omitempty" tf:"static_routes_only"`
	// +optional
	Tags *map[string]string `json:"tags,omitempty" tf:"tags"`
	// +optional
	TransitGatewayAttachmentID *string `json:"transitGatewayAttachmentID,omitempty" tf:"transit_gateway_attachment_id"`
	// +optional
	TransitGatewayID *string `json:"transitGatewayID,omitempty" tf:"transit_gateway_id"`
	// +optional
	Tunnel1Address *string `json:"tunnel1Address,omitempty" tf:"tunnel1_address"`
	// +optional
	Tunnel1BGPAsn *string `json:"tunnel1BGPAsn,omitempty" tf:"tunnel1_bgp_asn"`
	// +optional
	Tunnel1BGPHoldtime *int64 `json:"tunnel1BGPHoldtime,omitempty" tf:"tunnel1_bgp_holdtime"`
	// +optional
	Tunnel1CgwInsideAddress *string `json:"tunnel1CgwInsideAddress,omitempty" tf:"tunnel1_cgw_inside_address"`
	// +optional
	Tunnel1DpdTimeoutAction *string `json:"tunnel1DpdTimeoutAction,omitempty" tf:"tunnel1_dpd_timeout_action"`
	// +optional
	Tunnel1DpdTimeoutSeconds *int64 `json:"tunnel1DpdTimeoutSeconds,omitempty" tf:"tunnel1_dpd_timeout_seconds"`
	// +optional
	Tunnel1IkeVersions []string `json:"tunnel1IkeVersions,omitempty" tf:"tunnel1_ike_versions"`
	// +optional
	Tunnel1InsideCIDR *string `json:"tunnel1InsideCIDR,omitempty" tf:"tunnel1_inside_cidr"`
	// +optional
	Tunnel1InsideIpv6CIDR *string `json:"tunnel1InsideIpv6CIDR,omitempty" tf:"tunnel1_inside_ipv6_cidr"`
	// +optional
	Tunnel1Phase1DhGroupNumbers []int64 `json:"tunnel1Phase1DhGroupNumbers,omitempty" tf:"tunnel1_phase1_dh_group_numbers"`
	// +optional
	Tunnel1Phase1EncryptionAlgorithms []string `json:"tunnel1Phase1EncryptionAlgorithms,omitempty" tf:"tunnel1_phase1_encryption_algorithms"`
	// +optional
	Tunnel1Phase1IntegrityAlgorithms []string `json:"tunnel1Phase1IntegrityAlgorithms,omitempty" tf:"tunnel1_phase1_integrity_algorithms"`
	// +optional
	Tunnel1Phase1LifetimeSeconds *int64 `json:"tunnel1Phase1LifetimeSeconds,omitempty" tf:"tunnel1_phase1_lifetime_seconds"`
	// +optional
	Tunnel1Phase2DhGroupNumbers []int64 `json:"tunnel1Phase2DhGroupNumbers,omitempty" tf:"tunnel1_phase2_dh_group_numbers"`
	// +optional
	Tunnel1Phase2EncryptionAlgorithms []string `json:"tunnel1Phase2EncryptionAlgorithms,omitempty" tf:"tunnel1_phase2_encryption_algorithms"`
	// +optional
	Tunnel1Phase2IntegrityAlgorithms []string `json:"tunnel1Phase2IntegrityAlgorithms,omitempty" tf:"tunnel1_phase2_integrity_algorithms"`
	// +optional
	Tunnel1Phase2LifetimeSeconds *int64 `json:"tunnel1Phase2LifetimeSeconds,omitempty" tf:"tunnel1_phase2_lifetime_seconds"`
	// +optional
	Tunnel1PresharedKey *string `json:"-" sensitive:"true" tf:"tunnel1_preshared_key"`
	// +optional
	Tunnel1RekeyFuzzPercentage *int64 `json:"tunnel1RekeyFuzzPercentage,omitempty" tf:"tunnel1_rekey_fuzz_percentage"`
	// +optional
	Tunnel1RekeyMarginTimeSeconds *int64 `json:"tunnel1RekeyMarginTimeSeconds,omitempty" tf:"tunnel1_rekey_margin_time_seconds"`
	// +optional
	Tunnel1ReplayWindowSize *int64 `json:"tunnel1ReplayWindowSize,omitempty" tf:"tunnel1_replay_window_size"`
	// +optional
	Tunnel1StartupAction *string `json:"tunnel1StartupAction,omitempty" tf:"tunnel1_startup_action"`
	// +optional
	Tunnel1VgwInsideAddress *string `json:"tunnel1VgwInsideAddress,omitempty" tf:"tunnel1_vgw_inside_address"`
	// +optional
	Tunnel2Address *string `json:"tunnel2Address,omitempty" tf:"tunnel2_address"`
	// +optional
	Tunnel2BGPAsn *string `json:"tunnel2BGPAsn,omitempty" tf:"tunnel2_bgp_asn"`
	// +optional
	Tunnel2BGPHoldtime *int64 `json:"tunnel2BGPHoldtime,omitempty" tf:"tunnel2_bgp_holdtime"`
	// +optional
	Tunnel2CgwInsideAddress *string `json:"tunnel2CgwInsideAddress,omitempty" tf:"tunnel2_cgw_inside_address"`
	// +optional
	Tunnel2DpdTimeoutAction *string `json:"tunnel2DpdTimeoutAction,omitempty" tf:"tunnel2_dpd_timeout_action"`
	// +optional
	Tunnel2DpdTimeoutSeconds *int64 `json:"tunnel2DpdTimeoutSeconds,omitempty" tf:"tunnel2_dpd_timeout_seconds"`
	// +optional
	Tunnel2IkeVersions []string `json:"tunnel2IkeVersions,omitempty" tf:"tunnel2_ike_versions"`
	// +optional
	Tunnel2InsideCIDR *string `json:"tunnel2InsideCIDR,omitempty" tf:"tunnel2_inside_cidr"`
	// +optional
	Tunnel2InsideIpv6CIDR *string `json:"tunnel2InsideIpv6CIDR,omitempty" tf:"tunnel2_inside_ipv6_cidr"`
	// +optional
	Tunnel2Phase1DhGroupNumbers []int64 `json:"tunnel2Phase1DhGroupNumbers,omitempty" tf:"tunnel2_phase1_dh_group_numbers"`
	// +optional
	Tunnel2Phase1EncryptionAlgorithms []string `json:"tunnel2Phase1EncryptionAlgorithms,omitempty" tf:"tunnel2_phase1_encryption_algorithms"`
	// +optional
	Tunnel2Phase1IntegrityAlgorithms []string `json:"tunnel2Phase1IntegrityAlgorithms,omitempty" tf:"tunnel2_phase1_integrity_algorithms"`
	// +optional
	Tunnel2Phase1LifetimeSeconds *int64 `json:"tunnel2Phase1LifetimeSeconds,omitempty" tf:"tunnel2_phase1_lifetime_seconds"`
	// +optional
	Tunnel2Phase2DhGroupNumbers []int64 `json:"tunnel2Phase2DhGroupNumbers,omitempty" tf:"tunnel2_phase2_dh_group_numbers"`
	// +optional
	Tunnel2Phase2EncryptionAlgorithms []string `json:"tunnel2Phase2EncryptionAlgorithms,omitempty" tf:"tunnel2_phase2_encryption_algorithms"`
	// +optional
	Tunnel2Phase2IntegrityAlgorithms []string `json:"tunnel2Phase2IntegrityAlgorithms,omitempty" tf:"tunnel2_phase2_integrity_algorithms"`
	// +optional
	Tunnel2Phase2LifetimeSeconds *int64 `json:"tunnel2Phase2LifetimeSeconds,omitempty" tf:"tunnel2_phase2_lifetime_seconds"`
	// +optional
	Tunnel2PresharedKey *string `json:"-" sensitive:"true" tf:"tunnel2_preshared_key"`
	// +optional
	Tunnel2RekeyFuzzPercentage *int64 `json:"tunnel2RekeyFuzzPercentage,omitempty" tf:"tunnel2_rekey_fuzz_percentage"`
	// +optional
	Tunnel2RekeyMarginTimeSeconds *int64 `json:"tunnel2RekeyMarginTimeSeconds,omitempty" tf:"tunnel2_rekey_margin_time_seconds"`
	// +optional
	Tunnel2ReplayWindowSize *int64 `json:"tunnel2ReplayWindowSize,omitempty" tf:"tunnel2_replay_window_size"`
	// +optional
	Tunnel2StartupAction *string `json:"tunnel2StartupAction,omitempty" tf:"tunnel2_startup_action"`
	// +optional
	Tunnel2VgwInsideAddress *string `json:"tunnel2VgwInsideAddress,omitempty" tf:"tunnel2_vgw_inside_address"`
	// +optional
	TunnelInsideIPVersion *string `json:"tunnelInsideIPVersion,omitempty" tf:"tunnel_inside_ip_version"`
	Type                  *string `json:"type" tf:"type"`
	// +optional
	VgwTelemetry []ConnectionSpecVgwTelemetry `json:"vgwTelemetry,omitempty" tf:"vgw_telemetry"`
	// +optional
	VpnGatewayID *string `json:"vpnGatewayID,omitempty" tf:"vpn_gateway_id"`
}

type ConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
	// +optional
	Conditions []kmapi.Condition `json:"conditions,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ConnectionList is a list of Connections
type ConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Connection CRD objects
	Items []Connection `json:"items,omitempty"`
}
