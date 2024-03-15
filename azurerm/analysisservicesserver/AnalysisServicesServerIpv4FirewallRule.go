// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package analysisservicesserver


type AnalysisServicesServerIpv4FirewallRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/analysis_services_server#name AnalysisServicesServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/analysis_services_server#range_end AnalysisServicesServer#range_end}.
	RangeEnd *string `field:"required" json:"rangeEnd" yaml:"rangeEnd"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/analysis_services_server#range_start AnalysisServicesServer#range_start}.
	RangeStart *string `field:"required" json:"rangeStart" yaml:"rangeStart"`
}

