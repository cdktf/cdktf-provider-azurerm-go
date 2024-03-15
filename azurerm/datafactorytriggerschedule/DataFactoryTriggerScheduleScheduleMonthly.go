// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggerschedule


type DataFactoryTriggerScheduleScheduleMonthly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/data_factory_trigger_schedule#weekday DataFactoryTriggerSchedule#weekday}.
	Weekday *string `field:"required" json:"weekday" yaml:"weekday"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/data_factory_trigger_schedule#week DataFactoryTriggerSchedule#week}.
	Week *float64 `field:"optional" json:"week" yaml:"week"`
}

