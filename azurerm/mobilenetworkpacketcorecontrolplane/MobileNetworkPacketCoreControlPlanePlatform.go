// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkpacketcorecontrolplane


type MobileNetworkPacketCoreControlPlanePlatform struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/mobile_network_packet_core_control_plane#type MobileNetworkPacketCoreControlPlane#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/mobile_network_packet_core_control_plane#arc_kubernetes_cluster_id MobileNetworkPacketCoreControlPlane#arc_kubernetes_cluster_id}.
	ArcKubernetesClusterId *string `field:"optional" json:"arcKubernetesClusterId" yaml:"arcKubernetesClusterId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/mobile_network_packet_core_control_plane#custom_location_id MobileNetworkPacketCoreControlPlane#custom_location_id}.
	CustomLocationId *string `field:"optional" json:"customLocationId" yaml:"customLocationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/mobile_network_packet_core_control_plane#edge_device_id MobileNetworkPacketCoreControlPlane#edge_device_id}.
	EdgeDeviceId *string `field:"optional" json:"edgeDeviceId" yaml:"edgeDeviceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/mobile_network_packet_core_control_plane#stack_hci_cluster_id MobileNetworkPacketCoreControlPlane#stack_hci_cluster_id}.
	StackHciClusterId *string `field:"optional" json:"stackHciClusterId" yaml:"stackHciClusterId"`
}

