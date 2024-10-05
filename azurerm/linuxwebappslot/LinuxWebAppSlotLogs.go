// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebappslot


type LinuxWebAppSlotLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/linux_web_app_slot#application_logs LinuxWebAppSlot#application_logs}
	ApplicationLogs *LinuxWebAppSlotLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/linux_web_app_slot#detailed_error_messages LinuxWebAppSlot#detailed_error_messages}.
	DetailedErrorMessages interface{} `field:"optional" json:"detailedErrorMessages" yaml:"detailedErrorMessages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/linux_web_app_slot#failed_request_tracing LinuxWebAppSlot#failed_request_tracing}.
	FailedRequestTracing interface{} `field:"optional" json:"failedRequestTracing" yaml:"failedRequestTracing"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.4.0/docs/resources/linux_web_app_slot#http_logs LinuxWebAppSlot#http_logs}
	HttpLogs *LinuxWebAppSlotLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

