// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint


type CdnEndpointGeoFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/cdn_endpoint#action CdnEndpoint#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/cdn_endpoint#country_codes CdnEndpoint#country_codes}.
	CountryCodes *[]*string `field:"required" json:"countryCodes" yaml:"countryCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/cdn_endpoint#relative_path CdnEndpoint#relative_path}.
	RelativePath *string `field:"required" json:"relativePath" yaml:"relativePath"`
}

