// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containergroup


type ContainerGroupDnsConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/container_group#nameservers ContainerGroup#nameservers}.
	Nameservers *[]*string `field:"required" json:"nameservers" yaml:"nameservers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/container_group#options ContainerGroup#options}.
	Options *[]*string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/container_group#search_domains ContainerGroup#search_domains}.
	SearchDomains *[]*string `field:"optional" json:"searchDomains" yaml:"searchDomains"`
}

