// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containergroup


type ContainerGroupDiagnostics struct {
	// log_analytics block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_group#log_analytics ContainerGroup#log_analytics}
	LogAnalytics *ContainerGroupDiagnosticsLogAnalytics `field:"required" json:"logAnalytics" yaml:"logAnalytics"`
}

