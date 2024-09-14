// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot


type WindowsWebAppSlotLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app_slot#application_logs WindowsWebAppSlot#application_logs}
	ApplicationLogs *WindowsWebAppSlotLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app_slot#detailed_error_messages WindowsWebAppSlot#detailed_error_messages}.
	DetailedErrorMessages interface{} `field:"optional" json:"detailedErrorMessages" yaml:"detailedErrorMessages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app_slot#failed_request_tracing WindowsWebAppSlot#failed_request_tracing}.
	FailedRequestTracing interface{} `field:"optional" json:"failedRequestTracing" yaml:"failedRequestTracing"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/windows_web_app_slot#http_logs WindowsWebAppSlot#http_logs}
	HttpLogs *WindowsWebAppSlotLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

