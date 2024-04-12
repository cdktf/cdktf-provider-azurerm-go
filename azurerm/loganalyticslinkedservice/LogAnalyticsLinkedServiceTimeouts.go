// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package loganalyticslinkedservice


type LogAnalyticsLinkedServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/log_analytics_linked_service#create LogAnalyticsLinkedService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/log_analytics_linked_service#delete LogAnalyticsLinkedService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/log_analytics_linked_service#read LogAnalyticsLinkedService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/log_analytics_linked_service#update LogAnalyticsLinkedService#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

