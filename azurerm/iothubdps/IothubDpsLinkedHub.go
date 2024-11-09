// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iothubdps


type IothubDpsLinkedHub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/iothub_dps#connection_string IothubDps#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/iothub_dps#location IothubDps#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/iothub_dps#allocation_weight IothubDps#allocation_weight}.
	AllocationWeight *float64 `field:"optional" json:"allocationWeight" yaml:"allocationWeight"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/iothub_dps#apply_allocation_policy IothubDps#apply_allocation_policy}.
	ApplyAllocationPolicy interface{} `field:"optional" json:"applyAllocationPolicy" yaml:"applyAllocationPolicy"`
}

