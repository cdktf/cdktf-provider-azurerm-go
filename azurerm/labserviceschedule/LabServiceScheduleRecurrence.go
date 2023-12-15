// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labserviceschedule


type LabServiceScheduleRecurrence struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/lab_service_schedule#expiration_date LabServiceSchedule#expiration_date}.
	ExpirationDate *string `field:"required" json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/lab_service_schedule#frequency LabServiceSchedule#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/lab_service_schedule#interval LabServiceSchedule#interval}.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/lab_service_schedule#week_days LabServiceSchedule#week_days}.
	WeekDays *[]*string `field:"optional" json:"weekDays" yaml:"weekDays"`
}

