// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticslinkedstorageaccount


type LogAnalyticsLinkedStorageAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/log_analytics_linked_storage_account#create LogAnalyticsLinkedStorageAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/log_analytics_linked_storage_account#delete LogAnalyticsLinkedStorageAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/log_analytics_linked_storage_account#read LogAnalyticsLinkedStorageAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/log_analytics_linked_storage_account#update LogAnalyticsLinkedStorageAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

