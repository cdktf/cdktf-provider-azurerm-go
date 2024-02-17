// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksimpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkSimPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#default_slice_id MobileNetworkSimPolicy#default_slice_id}.
	DefaultSliceId *string `field:"required" json:"defaultSliceId" yaml:"defaultSliceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#location MobileNetworkSimPolicy#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#mobile_network_id MobileNetworkSimPolicy#mobile_network_id}.
	MobileNetworkId *string `field:"required" json:"mobileNetworkId" yaml:"mobileNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#name MobileNetworkSimPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// slice block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#slice MobileNetworkSimPolicy#slice}
	Slice interface{} `field:"required" json:"slice" yaml:"slice"`
	// user_equipment_aggregate_maximum_bit_rate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#user_equipment_aggregate_maximum_bit_rate MobileNetworkSimPolicy#user_equipment_aggregate_maximum_bit_rate}
	UserEquipmentAggregateMaximumBitRate *MobileNetworkSimPolicyUserEquipmentAggregateMaximumBitRate `field:"required" json:"userEquipmentAggregateMaximumBitRate" yaml:"userEquipmentAggregateMaximumBitRate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#id MobileNetworkSimPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#rat_frequency_selection_priority_index MobileNetworkSimPolicy#rat_frequency_selection_priority_index}.
	RatFrequencySelectionPriorityIndex *float64 `field:"optional" json:"ratFrequencySelectionPriorityIndex" yaml:"ratFrequencySelectionPriorityIndex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#registration_timer_in_seconds MobileNetworkSimPolicy#registration_timer_in_seconds}.
	RegistrationTimerInSeconds *float64 `field:"optional" json:"registrationTimerInSeconds" yaml:"registrationTimerInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#tags MobileNetworkSimPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/mobile_network_sim_policy#timeouts MobileNetworkSimPolicy#timeouts}
	Timeouts *MobileNetworkSimPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

