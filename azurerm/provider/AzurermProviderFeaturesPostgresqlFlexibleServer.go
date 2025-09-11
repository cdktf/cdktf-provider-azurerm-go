// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type AzurermProviderFeaturesPostgresqlFlexibleServer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs#restart_server_on_configuration_value_change AzurermProvider#restart_server_on_configuration_value_change}.
	RestartServerOnConfigurationValueChange interface{} `field:"optional" json:"restartServerOnConfigurationValueChange" yaml:"restartServerOnConfigurationValueChange"`
}

