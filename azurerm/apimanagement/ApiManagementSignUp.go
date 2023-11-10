// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagement


type ApiManagementSignUp struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/api_management#enabled ApiManagement#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// terms_of_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/api_management#terms_of_service ApiManagement#terms_of_service}
	TermsOfService *ApiManagementSignUpTermsOfService `field:"required" json:"termsOfService" yaml:"termsOfService"`
}

