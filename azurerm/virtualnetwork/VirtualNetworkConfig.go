// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetwork

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#address_space VirtualNetwork#address_space}.
	AddressSpace *[]*string `field:"required" json:"addressSpace" yaml:"addressSpace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#location VirtualNetwork#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#name VirtualNetwork#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#resource_group_name VirtualNetwork#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#bgp_community VirtualNetwork#bgp_community}.
	BgpCommunity *string `field:"optional" json:"bgpCommunity" yaml:"bgpCommunity"`
	// ddos_protection_plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#ddos_protection_plan VirtualNetwork#ddos_protection_plan}
	DdosProtectionPlan *VirtualNetworkDdosProtectionPlan `field:"optional" json:"ddosProtectionPlan" yaml:"ddosProtectionPlan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#dns_servers VirtualNetwork#dns_servers}.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#edge_zone VirtualNetwork#edge_zone}.
	EdgeZone *string `field:"optional" json:"edgeZone" yaml:"edgeZone"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#encryption VirtualNetwork#encryption}
	Encryption *VirtualNetworkEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#flow_timeout_in_minutes VirtualNetwork#flow_timeout_in_minutes}.
	FlowTimeoutInMinutes *float64 `field:"optional" json:"flowTimeoutInMinutes" yaml:"flowTimeoutInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#id VirtualNetwork#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#subnet VirtualNetwork#subnet}.
	Subnet interface{} `field:"optional" json:"subnet" yaml:"subnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#tags VirtualNetwork#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/virtual_network#timeouts VirtualNetwork#timeouts}
	Timeouts *VirtualNetworkTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

