// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotSiteConfigAutoHealSettingTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/windows_web_app_slot#private_memory_kb WindowsWebAppSlot#private_memory_kb}.
	PrivateMemoryKb *float64 `field:"optional" json:"privateMemoryKb" yaml:"privateMemoryKb"`
	// requests block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/windows_web_app_slot#requests WindowsWebAppSlot#requests}
	Requests *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerRequests `field:"optional" json:"requests" yaml:"requests"`
	// slow_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/windows_web_app_slot#slow_request WindowsWebAppSlot#slow_request}
	SlowRequest *WindowsWebAppSlotSiteConfigAutoHealSettingTriggerSlowRequest `field:"optional" json:"slowRequest" yaml:"slowRequest"`
	// slow_request_with_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/windows_web_app_slot#slow_request_with_path WindowsWebAppSlot#slow_request_with_path}
	SlowRequestWithPath interface{} `field:"optional" json:"slowRequestWithPath" yaml:"slowRequestWithPath"`
	// status_code block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/windows_web_app_slot#status_code WindowsWebAppSlot#status_code}
	StatusCode interface{} `field:"optional" json:"statusCode" yaml:"statusCode"`
}

