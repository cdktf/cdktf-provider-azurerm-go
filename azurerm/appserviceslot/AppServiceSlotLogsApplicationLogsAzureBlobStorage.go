// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceslot


type AppServiceSlotLogsApplicationLogsAzureBlobStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/app_service_slot#level AppServiceSlot#level}.
	Level *string `field:"required" json:"level" yaml:"level"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/app_service_slot#retention_in_days AppServiceSlot#retention_in_days}.
	RetentionInDays *float64 `field:"required" json:"retentionInDays" yaml:"retentionInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/app_service_slot#sas_url AppServiceSlot#sas_url}.
	SasUrl *string `field:"required" json:"sasUrl" yaml:"sasUrl"`
}

