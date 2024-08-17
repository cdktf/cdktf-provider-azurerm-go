// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecurityrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkSecurityRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#access NetworkSecurityRule#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#direction NetworkSecurityRule#direction}.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#name NetworkSecurityRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#network_security_group_name NetworkSecurityRule#network_security_group_name}.
	NetworkSecurityGroupName *string `field:"required" json:"networkSecurityGroupName" yaml:"networkSecurityGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#priority NetworkSecurityRule#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#protocol NetworkSecurityRule#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#resource_group_name NetworkSecurityRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#description NetworkSecurityRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#destination_address_prefix NetworkSecurityRule#destination_address_prefix}.
	DestinationAddressPrefix *string `field:"optional" json:"destinationAddressPrefix" yaml:"destinationAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#destination_address_prefixes NetworkSecurityRule#destination_address_prefixes}.
	DestinationAddressPrefixes *[]*string `field:"optional" json:"destinationAddressPrefixes" yaml:"destinationAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#destination_application_security_group_ids NetworkSecurityRule#destination_application_security_group_ids}.
	DestinationApplicationSecurityGroupIds *[]*string `field:"optional" json:"destinationApplicationSecurityGroupIds" yaml:"destinationApplicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#destination_port_range NetworkSecurityRule#destination_port_range}.
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#destination_port_ranges NetworkSecurityRule#destination_port_ranges}.
	DestinationPortRanges *[]*string `field:"optional" json:"destinationPortRanges" yaml:"destinationPortRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#id NetworkSecurityRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#source_address_prefix NetworkSecurityRule#source_address_prefix}.
	SourceAddressPrefix *string `field:"optional" json:"sourceAddressPrefix" yaml:"sourceAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#source_address_prefixes NetworkSecurityRule#source_address_prefixes}.
	SourceAddressPrefixes *[]*string `field:"optional" json:"sourceAddressPrefixes" yaml:"sourceAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#source_application_security_group_ids NetworkSecurityRule#source_application_security_group_ids}.
	SourceApplicationSecurityGroupIds *[]*string `field:"optional" json:"sourceApplicationSecurityGroupIds" yaml:"sourceApplicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#source_port_range NetworkSecurityRule#source_port_range}.
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#source_port_ranges NetworkSecurityRule#source_port_ranges}.
	SourcePortRanges *[]*string `field:"optional" json:"sourcePortRanges" yaml:"sourcePortRanges"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/network_security_rule#timeouts NetworkSecurityRule#timeouts}
	Timeouts *NetworkSecurityRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

