// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedlustrefilesystem


type ManagedLustreFileSystemMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/managed_lustre_file_system#day_of_week ManagedLustreFileSystem#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/managed_lustre_file_system#time_of_day_in_utc ManagedLustreFileSystem#time_of_day_in_utc}.
	TimeOfDayInUtc *string `field:"required" json:"timeOfDayInUtc" yaml:"timeOfDayInUtc"`
}

