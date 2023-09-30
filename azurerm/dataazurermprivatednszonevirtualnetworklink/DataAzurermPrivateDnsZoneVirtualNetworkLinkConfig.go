// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermprivatednszonevirtualnetworklink

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermPrivateDnsZoneVirtualNetworkLinkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/data-sources/private_dns_zone_virtual_network_link#name DataAzurermPrivateDnsZoneVirtualNetworkLink#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/data-sources/private_dns_zone_virtual_network_link#private_dns_zone_name DataAzurermPrivateDnsZoneVirtualNetworkLink#private_dns_zone_name}.
	PrivateDnsZoneName *string `field:"required" json:"privateDnsZoneName" yaml:"privateDnsZoneName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/data-sources/private_dns_zone_virtual_network_link#resource_group_name DataAzurermPrivateDnsZoneVirtualNetworkLink#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/data-sources/private_dns_zone_virtual_network_link#id DataAzurermPrivateDnsZoneVirtualNetworkLink#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/data-sources/private_dns_zone_virtual_network_link#timeouts DataAzurermPrivateDnsZoneVirtualNetworkLink#timeouts}
	Timeouts *DataAzurermPrivateDnsZoneVirtualNetworkLinkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

