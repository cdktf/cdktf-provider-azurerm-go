// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkconnectionmonitor


type NetworkConnectionMonitorTestConfigurationHttpConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/network_connection_monitor#method NetworkConnectionMonitor#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/network_connection_monitor#path NetworkConnectionMonitor#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/network_connection_monitor#port NetworkConnectionMonitor#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/network_connection_monitor#prefer_https NetworkConnectionMonitor#prefer_https}.
	PreferHttps interface{} `field:"optional" json:"preferHttps" yaml:"preferHttps"`
	// request_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/network_connection_monitor#request_header NetworkConnectionMonitor#request_header}
	RequestHeader interface{} `field:"optional" json:"requestHeader" yaml:"requestHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/network_connection_monitor#valid_status_code_ranges NetworkConnectionMonitor#valid_status_code_ranges}.
	ValidStatusCodeRanges *[]*string `field:"optional" json:"validStatusCodeRanges" yaml:"validStatusCodeRanges"`
}

