// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestschedule


type DevTestScheduleDailyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dev_test_schedule#time DevTestSchedule#time}.
	Time *string `field:"required" json:"time" yaml:"time"`
}

