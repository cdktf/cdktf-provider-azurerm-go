// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementlogger


type ApiManagementLoggerEventhub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/api_management_logger#connection_string ApiManagementLogger#connection_string}.
	ConnectionString *string `field:"required" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/api_management_logger#name ApiManagementLogger#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

