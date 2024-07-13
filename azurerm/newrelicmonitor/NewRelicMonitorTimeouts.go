// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package newrelicmonitor


type NewRelicMonitorTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/new_relic_monitor#create NewRelicMonitor#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/new_relic_monitor#delete NewRelicMonitor#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/new_relic_monitor#read NewRelicMonitor#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

