// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networksecuritygroup


type NetworkSecurityGroupSecurityRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#access NetworkSecurityGroup#access}.
	Access *string `field:"optional" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#description NetworkSecurityGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#destination_address_prefix NetworkSecurityGroup#destination_address_prefix}.
	DestinationAddressPrefix *string `field:"optional" json:"destinationAddressPrefix" yaml:"destinationAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#destination_address_prefixes NetworkSecurityGroup#destination_address_prefixes}.
	DestinationAddressPrefixes *[]*string `field:"optional" json:"destinationAddressPrefixes" yaml:"destinationAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#destination_application_security_group_ids NetworkSecurityGroup#destination_application_security_group_ids}.
	DestinationApplicationSecurityGroupIds *[]*string `field:"optional" json:"destinationApplicationSecurityGroupIds" yaml:"destinationApplicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#destination_port_range NetworkSecurityGroup#destination_port_range}.
	DestinationPortRange *string `field:"optional" json:"destinationPortRange" yaml:"destinationPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#destination_port_ranges NetworkSecurityGroup#destination_port_ranges}.
	DestinationPortRanges *[]*string `field:"optional" json:"destinationPortRanges" yaml:"destinationPortRanges"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#direction NetworkSecurityGroup#direction}.
	Direction *string `field:"optional" json:"direction" yaml:"direction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#name NetworkSecurityGroup#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#priority NetworkSecurityGroup#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#protocol NetworkSecurityGroup#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#source_address_prefix NetworkSecurityGroup#source_address_prefix}.
	SourceAddressPrefix *string `field:"optional" json:"sourceAddressPrefix" yaml:"sourceAddressPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#source_address_prefixes NetworkSecurityGroup#source_address_prefixes}.
	SourceAddressPrefixes *[]*string `field:"optional" json:"sourceAddressPrefixes" yaml:"sourceAddressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#source_application_security_group_ids NetworkSecurityGroup#source_application_security_group_ids}.
	SourceApplicationSecurityGroupIds *[]*string `field:"optional" json:"sourceApplicationSecurityGroupIds" yaml:"sourceApplicationSecurityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#source_port_range NetworkSecurityGroup#source_port_range}.
	SourcePortRange *string `field:"optional" json:"sourcePortRange" yaml:"sourcePortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/network_security_group#source_port_ranges NetworkSecurityGroup#source_port_ranges}.
	SourcePortRanges *[]*string `field:"optional" json:"sourcePortRanges" yaml:"sourcePortRanges"`
}

