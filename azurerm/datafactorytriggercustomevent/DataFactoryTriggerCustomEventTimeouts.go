// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggercustomevent


type DataFactoryTriggerCustomEventTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/data_factory_trigger_custom_event#create DataFactoryTriggerCustomEvent#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/data_factory_trigger_custom_event#delete DataFactoryTriggerCustomEvent#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/data_factory_trigger_custom_event#read DataFactoryTriggerCustomEvent#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/data_factory_trigger_custom_event#update DataFactoryTriggerCustomEvent#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

