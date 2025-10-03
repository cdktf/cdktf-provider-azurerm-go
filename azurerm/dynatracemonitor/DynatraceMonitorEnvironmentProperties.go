// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dynatracemonitor


type DynatraceMonitorEnvironmentProperties struct {
	// environment_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/dynatrace_monitor#environment_info DynatraceMonitor#environment_info}
	EnvironmentInfo interface{} `field:"required" json:"environmentInfo" yaml:"environmentInfo"`
}

