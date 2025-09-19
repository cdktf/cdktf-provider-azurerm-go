// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot


type LinuxWebAppSlotLogsHttpLogsFileSystem struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_web_app_slot#retention_in_days LinuxWebAppSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_web_app_slot#retention_in_mb LinuxWebAppSlot#retention_in_mb}.
	RetentionInMb *float64 `field:"required" json:"retentionInMb" yaml:"retentionInMb"`
}

