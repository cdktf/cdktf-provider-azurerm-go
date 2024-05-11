// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappactiveslot

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppActiveSlotConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The ID of the Slot to swap with `Production`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/function_app_active_slot#slot_id FunctionAppActiveSlot#slot_id}
	SlotId *string `field:"required" json:"slotId" yaml:"slotId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/function_app_active_slot#id FunctionAppActiveSlot#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The swap action should overwrite the Production slot's network configuration with the configuration from this slot. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/function_app_active_slot#overwrite_network_config FunctionAppActiveSlot#overwrite_network_config}
	OverwriteNetworkConfig interface{} `field:"optional" json:"overwriteNetworkConfig" yaml:"overwriteNetworkConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/function_app_active_slot#timeouts FunctionAppActiveSlot#timeouts}
	Timeouts *FunctionAppActiveSlotTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

