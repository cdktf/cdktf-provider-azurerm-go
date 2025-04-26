// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestschedule


type DevTestScheduleHourlyRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.27.0/docs/resources/dev_test_schedule#minute DevTestSchedule#minute}.
	Minute *float64 `field:"required" json:"minute" yaml:"minute"`
}

