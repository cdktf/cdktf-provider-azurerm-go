// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualdesktophostpool


type VirtualDesktopHostPoolScheduledAgentUpdates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_desktop_host_pool#enabled VirtualDesktopHostPool#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_desktop_host_pool#schedule VirtualDesktopHostPool#schedule}
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_desktop_host_pool#timezone VirtualDesktopHostPool#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/virtual_desktop_host_pool#use_session_host_timezone VirtualDesktopHostPool#use_session_host_timezone}.
	UseSessionHostTimezone interface{} `field:"optional" json:"useSessionHostTimezone" yaml:"useSessionHostTimezone"`
}

