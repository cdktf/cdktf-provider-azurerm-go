// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservice


type AppServiceLogs struct {
	// application_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service#application_logs AppService#application_logs}
	ApplicationLogs *AppServiceLogsApplicationLogs `field:"optional" json:"applicationLogs" yaml:"applicationLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service#detailed_error_messages_enabled AppService#detailed_error_messages_enabled}.
	DetailedErrorMessagesEnabled interface{} `field:"optional" json:"detailedErrorMessagesEnabled" yaml:"detailedErrorMessagesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service#failed_request_tracing_enabled AppService#failed_request_tracing_enabled}.
	FailedRequestTracingEnabled interface{} `field:"optional" json:"failedRequestTracingEnabled" yaml:"failedRequestTracingEnabled"`
	// http_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/app_service#http_logs AppService#http_logs}
	HttpLogs *AppServiceLogsHttpLogs `field:"optional" json:"httpLogs" yaml:"httpLogs"`
}

