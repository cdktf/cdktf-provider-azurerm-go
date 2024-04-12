// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservice


type AppServiceLogsHttpLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/app_service#retention_in_days AppService#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/app_service#sas_url AppService#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

