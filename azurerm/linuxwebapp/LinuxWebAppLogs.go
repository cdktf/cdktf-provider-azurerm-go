// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/linux_web_app#application_logs LinuxWebApp#application_logs}
	ApplicationLogs *LinuxWebAppLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/linux_web_app#detailed_error_messages LinuxWebApp#detailed_error_messages}.
	DetailedErrorMessages interface{} `field:"optional" json:"detailedErrorMessages" yaml:"detailedErrorMessages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/linux_web_app#failed_request_tracing LinuxWebApp#failed_request_tracing}.
	FailedRequestTracing interface{} `field:"optional" json:"failedRequestTracing" yaml:"failedRequestTracing"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/linux_web_app#http_logs LinuxWebApp#http_logs}
	HttpLogs *LinuxWebAppLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

