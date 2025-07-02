// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp


type LinuxFunctionAppSiteConfigAppServiceLogs struct {
	// The amount of disk space to use for logs. Valid values are between `25` and `100`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/linux_function_app#disk_quota_mb LinuxFunctionApp#disk_quota_mb}
	DiskQuotaMb *float64 `field:"optional" json:"diskQuotaMb" yaml:"diskQuotaMb"`
	// The retention period for logs in days. Valid values are between `0` and `99999`. Defaults to `0` (never delete).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/linux_function_app#retention_period_days LinuxFunctionApp#retention_period_days}
	RetentionPeriodDays *float64 `field:"optional" json:"retentionPeriodDays" yaml:"retentionPeriodDays"`
}

