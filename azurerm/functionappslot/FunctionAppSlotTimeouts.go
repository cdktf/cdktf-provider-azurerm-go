// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot


type FunctionAppSlotTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_slot#create FunctionAppSlot#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_slot#delete FunctionAppSlot#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_slot#read FunctionAppSlot#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/function_app_slot#update FunctionAppSlot#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

