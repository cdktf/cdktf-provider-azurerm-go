// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayBackendHttpSettingsConnectionDraining struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/application_gateway#drain_timeout_sec ApplicationGateway#drain_timeout_sec}.
	DrainTimeoutSec *float64 `field:"required" json:"drainTimeoutSec" yaml:"drainTimeoutSec"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.29.0/docs/resources/application_gateway#enabled ApplicationGateway#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
}

