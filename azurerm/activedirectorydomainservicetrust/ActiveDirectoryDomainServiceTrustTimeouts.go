// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomainservicetrust


type ActiveDirectoryDomainServiceTrustTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/active_directory_domain_service_trust#create ActiveDirectoryDomainServiceTrust#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/active_directory_domain_service_trust#delete ActiveDirectoryDomainServiceTrust#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/active_directory_domain_service_trust#read ActiveDirectoryDomainServiceTrust#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/active_directory_domain_service_trust#update ActiveDirectoryDomainServiceTrust#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

