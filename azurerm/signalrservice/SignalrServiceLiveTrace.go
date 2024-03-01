// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservice


type SignalrServiceLiveTrace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/signalr_service#connectivity_logs_enabled SignalrService#connectivity_logs_enabled}.
	ConnectivityLogsEnabled interface{} `field:"optional" json:"connectivityLogsEnabled" yaml:"connectivityLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/signalr_service#enabled SignalrService#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/signalr_service#http_request_logs_enabled SignalrService#http_request_logs_enabled}.
	HttpRequestLogsEnabled interface{} `field:"optional" json:"httpRequestLogsEnabled" yaml:"httpRequestLogsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/signalr_service#messaging_logs_enabled SignalrService#messaging_logs_enabled}.
	MessagingLogsEnabled interface{} `field:"optional" json:"messagingLogsEnabled" yaml:"messagingLogsEnabled"`
}

