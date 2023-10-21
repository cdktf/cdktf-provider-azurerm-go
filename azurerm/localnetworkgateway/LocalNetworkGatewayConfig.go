// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package localnetworkgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LocalNetworkGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#location LocalNetworkGateway#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#name LocalNetworkGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#resource_group_name LocalNetworkGateway#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#address_space LocalNetworkGateway#address_space}.
	AddressSpace *[]*string `field:"optional" json:"addressSpace" yaml:"addressSpace"`
	// bgp_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#bgp_settings LocalNetworkGateway#bgp_settings}
	BgpSettings *LocalNetworkGatewayBgpSettings `field:"optional" json:"bgpSettings" yaml:"bgpSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#gateway_address LocalNetworkGateway#gateway_address}.
	GatewayAddress *string `field:"optional" json:"gatewayAddress" yaml:"gatewayAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#gateway_fqdn LocalNetworkGateway#gateway_fqdn}.
	GatewayFqdn *string `field:"optional" json:"gatewayFqdn" yaml:"gatewayFqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#id LocalNetworkGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#tags LocalNetworkGateway#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/local_network_gateway#timeouts LocalNetworkGateway#timeouts}
	Timeouts *LocalNetworkGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

