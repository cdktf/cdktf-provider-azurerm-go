// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceslot


type AppServiceSlotLogsHttpLogsFileSystem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/app_service_slot#retention_in_days AppServiceSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/app_service_slot#retention_in_mb AppServiceSlot#retention_in_mb}.
	RetentionInMb *float64 `field:"required" json:"retentionInMb" yaml:"retentionInMb"`
}

