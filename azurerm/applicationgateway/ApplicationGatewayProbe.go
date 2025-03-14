// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway


type ApplicationGatewayProbe struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#interval ApplicationGateway#interval}.
	Interval *float64 `field:"required" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#name ApplicationGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#path ApplicationGateway#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#protocol ApplicationGateway#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#timeout ApplicationGateway#timeout}.
	Timeout *float64 `field:"required" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#unhealthy_threshold ApplicationGateway#unhealthy_threshold}.
	UnhealthyThreshold *float64 `field:"required" json:"unhealthyThreshold" yaml:"unhealthyThreshold"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#host ApplicationGateway#host}.
	Host *string `field:"optional" json:"host" yaml:"host"`
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#match ApplicationGateway#match}
	Match *ApplicationGatewayProbeMatch `field:"optional" json:"match" yaml:"match"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#minimum_servers ApplicationGateway#minimum_servers}.
	MinimumServers *float64 `field:"optional" json:"minimumServers" yaml:"minimumServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#pick_host_name_from_backend_http_settings ApplicationGateway#pick_host_name_from_backend_http_settings}.
	PickHostNameFromBackendHttpSettings interface{} `field:"optional" json:"pickHostNameFromBackendHttpSettings" yaml:"pickHostNameFromBackendHttpSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/application_gateway#port ApplicationGateway#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

