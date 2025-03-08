// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorprivatelinkscopedservice


type MonitorPrivateLinkScopedServiceTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_private_link_scoped_service#create MonitorPrivateLinkScopedService#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_private_link_scoped_service#delete MonitorPrivateLinkScopedService#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/monitor_private_link_scoped_service#read MonitorPrivateLinkScopedService#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

