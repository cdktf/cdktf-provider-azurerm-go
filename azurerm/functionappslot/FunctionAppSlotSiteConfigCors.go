// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot


type FunctionAppSlotSiteConfigCors struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/function_app_slot#allowed_origins FunctionAppSlot#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/function_app_slot#support_credentials FunctionAppSlot#support_credentials}.
	SupportCredentials interface{} `field:"optional" json:"supportCredentials" yaml:"supportCredentials"`
}

