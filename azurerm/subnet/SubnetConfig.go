// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subnet

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubnetConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#address_prefixes Subnet#address_prefixes}.
	AddressPrefixes *[]*string `field:"required" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#name Subnet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#resource_group_name Subnet#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#virtual_network_name Subnet#virtual_network_name}.
	VirtualNetworkName *string `field:"required" json:"virtualNetworkName" yaml:"virtualNetworkName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#default_outbound_access_enabled Subnet#default_outbound_access_enabled}.
	DefaultOutboundAccessEnabled interface{} `field:"optional" json:"defaultOutboundAccessEnabled" yaml:"defaultOutboundAccessEnabled"`
	// delegation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#delegation Subnet#delegation}
	Delegation interface{} `field:"optional" json:"delegation" yaml:"delegation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#id Subnet#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#private_endpoint_network_policies Subnet#private_endpoint_network_policies}.
	PrivateEndpointNetworkPolicies *string `field:"optional" json:"privateEndpointNetworkPolicies" yaml:"privateEndpointNetworkPolicies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#private_link_service_network_policies_enabled Subnet#private_link_service_network_policies_enabled}.
	PrivateLinkServiceNetworkPoliciesEnabled interface{} `field:"optional" json:"privateLinkServiceNetworkPoliciesEnabled" yaml:"privateLinkServiceNetworkPoliciesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#service_endpoint_policy_ids Subnet#service_endpoint_policy_ids}.
	ServiceEndpointPolicyIds *[]*string `field:"optional" json:"serviceEndpointPolicyIds" yaml:"serviceEndpointPolicyIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#service_endpoints Subnet#service_endpoints}.
	ServiceEndpoints *[]*string `field:"optional" json:"serviceEndpoints" yaml:"serviceEndpoints"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/subnet#timeouts Subnet#timeouts}
	Timeouts *SubnetTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

