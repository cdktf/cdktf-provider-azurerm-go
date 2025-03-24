// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor


type FrontdoorBackendPoolBackend struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#address Frontdoor#address}.
	Address *string `field:"required" json:"address" yaml:"address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#host_header Frontdoor#host_header}.
	HostHeader *string `field:"required" json:"hostHeader" yaml:"hostHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#http_port Frontdoor#http_port}.
	HttpPort *float64 `field:"required" json:"httpPort" yaml:"httpPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#https_port Frontdoor#https_port}.
	HttpsPort *float64 `field:"required" json:"httpsPort" yaml:"httpsPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#priority Frontdoor#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/frontdoor#weight Frontdoor#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

