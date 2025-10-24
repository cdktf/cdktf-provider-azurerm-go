// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databricksvirtualnetworkpeering

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabricksVirtualNetworkPeeringConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#name DatabricksVirtualNetworkPeering#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#remote_address_space_prefixes DatabricksVirtualNetworkPeering#remote_address_space_prefixes}.
	RemoteAddressSpacePrefixes *[]*string `field:"required" json:"remoteAddressSpacePrefixes" yaml:"remoteAddressSpacePrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#remote_virtual_network_id DatabricksVirtualNetworkPeering#remote_virtual_network_id}.
	RemoteVirtualNetworkId *string `field:"required" json:"remoteVirtualNetworkId" yaml:"remoteVirtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#resource_group_name DatabricksVirtualNetworkPeering#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#workspace_id DatabricksVirtualNetworkPeering#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#allow_forwarded_traffic DatabricksVirtualNetworkPeering#allow_forwarded_traffic}.
	AllowForwardedTraffic interface{} `field:"optional" json:"allowForwardedTraffic" yaml:"allowForwardedTraffic"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#allow_gateway_transit DatabricksVirtualNetworkPeering#allow_gateway_transit}.
	AllowGatewayTransit interface{} `field:"optional" json:"allowGatewayTransit" yaml:"allowGatewayTransit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#allow_virtual_network_access DatabricksVirtualNetworkPeering#allow_virtual_network_access}.
	AllowVirtualNetworkAccess interface{} `field:"optional" json:"allowVirtualNetworkAccess" yaml:"allowVirtualNetworkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#id DatabricksVirtualNetworkPeering#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#timeouts DatabricksVirtualNetworkPeering#timeouts}
	Timeouts *DatabricksVirtualNetworkPeeringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.50.0/docs/resources/databricks_virtual_network_peering#use_remote_gateways DatabricksVirtualNetworkPeering#use_remote_gateways}.
	UseRemoteGateways interface{} `field:"optional" json:"useRemoteGateways" yaml:"useRemoteGateways"`
}

