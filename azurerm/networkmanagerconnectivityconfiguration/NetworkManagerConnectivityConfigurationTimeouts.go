// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectivityconfiguration


type NetworkManagerConnectivityConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/network_manager_connectivity_configuration#create NetworkManagerConnectivityConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/network_manager_connectivity_configuration#delete NetworkManagerConnectivityConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/network_manager_connectivity_configuration#read NetworkManagerConnectivityConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/network_manager_connectivity_configuration#update NetworkManagerConnectivityConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

