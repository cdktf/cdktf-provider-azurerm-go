// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiImport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_api#content_format ApiManagementApi#content_format}.
	ContentFormat *string `field:"required" json:"contentFormat" yaml:"contentFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_api#content_value ApiManagementApi#content_value}.
	ContentValue *string `field:"required" json:"contentValue" yaml:"contentValue"`
	// wsdl_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/api_management_api#wsdl_selector ApiManagementApi#wsdl_selector}
	WsdlSelector *ApiManagementApiImportWsdlSelector `field:"optional" json:"wsdlSelector" yaml:"wsdlSelector"`
}

