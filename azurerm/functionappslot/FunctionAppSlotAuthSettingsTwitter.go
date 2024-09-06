// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot


type FunctionAppSlotAuthSettingsTwitter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/function_app_slot#consumer_key FunctionAppSlot#consumer_key}.
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/function_app_slot#consumer_secret FunctionAppSlot#consumer_secret}.
	ConsumerSecret *string `field:"required" json:"consumerSecret" yaml:"consumerSecret"`
}

