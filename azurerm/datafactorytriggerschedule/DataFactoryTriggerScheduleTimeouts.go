// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggerschedule


type DataFactoryTriggerScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_trigger_schedule#create DataFactoryTriggerSchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_trigger_schedule#delete DataFactoryTriggerSchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_trigger_schedule#read DataFactoryTriggerSchedule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/data_factory_trigger_schedule#update DataFactoryTriggerSchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

