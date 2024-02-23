// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementlogger


type ApiManagementLoggerApplicationInsights struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/api_management_logger#instrumentation_key ApiManagementLogger#instrumentation_key}.
	InstrumentationKey *string `field:"required" json:"instrumentationKey" yaml:"instrumentationKey"`
}

