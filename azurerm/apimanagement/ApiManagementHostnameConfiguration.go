// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementHostnameConfiguration struct {
	// developer_portal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/api_management#developer_portal ApiManagement#developer_portal}
	DeveloperPortal interface{} `field:"optional" json:"developerPortal" yaml:"developerPortal"`
	// management block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/api_management#management ApiManagement#management}
	Management interface{} `field:"optional" json:"management" yaml:"management"`
	// portal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/api_management#portal ApiManagement#portal}
	Portal interface{} `field:"optional" json:"portal" yaml:"portal"`
	// proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/api_management#proxy ApiManagement#proxy}
	Proxy interface{} `field:"optional" json:"proxy" yaml:"proxy"`
	// scm block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/api_management#scm ApiManagement#scm}
	Scm interface{} `field:"optional" json:"scm" yaml:"scm"`
}

