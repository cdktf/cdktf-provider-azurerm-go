// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworkpacketcoredataplane

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkPacketCoreDataPlaneConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#location MobileNetworkPacketCoreDataPlane#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#mobile_network_packet_core_control_plane_id MobileNetworkPacketCoreDataPlane#mobile_network_packet_core_control_plane_id}.
	MobileNetworkPacketCoreControlPlaneId *string `field:"required" json:"mobileNetworkPacketCoreControlPlaneId" yaml:"mobileNetworkPacketCoreControlPlaneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#name MobileNetworkPacketCoreDataPlane#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#id MobileNetworkPacketCoreDataPlane#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#tags MobileNetworkPacketCoreDataPlane#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#timeouts MobileNetworkPacketCoreDataPlane#timeouts}
	Timeouts *MobileNetworkPacketCoreDataPlaneTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#user_plane_access_ipv4_address MobileNetworkPacketCoreDataPlane#user_plane_access_ipv4_address}.
	UserPlaneAccessIpv4Address *string `field:"optional" json:"userPlaneAccessIpv4Address" yaml:"userPlaneAccessIpv4Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#user_plane_access_ipv4_gateway MobileNetworkPacketCoreDataPlane#user_plane_access_ipv4_gateway}.
	UserPlaneAccessIpv4Gateway *string `field:"optional" json:"userPlaneAccessIpv4Gateway" yaml:"userPlaneAccessIpv4Gateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#user_plane_access_ipv4_subnet MobileNetworkPacketCoreDataPlane#user_plane_access_ipv4_subnet}.
	UserPlaneAccessIpv4Subnet *string `field:"optional" json:"userPlaneAccessIpv4Subnet" yaml:"userPlaneAccessIpv4Subnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/mobile_network_packet_core_data_plane#user_plane_access_name MobileNetworkPacketCoreDataPlane#user_plane_access_name}.
	UserPlaneAccessName *string `field:"optional" json:"userPlaneAccessName" yaml:"userPlaneAccessName"`
}

