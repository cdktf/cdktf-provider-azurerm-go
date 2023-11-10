// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanagerconnectivityconfiguration


type NetworkManagerConnectivityConfigurationHub struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/network_manager_connectivity_configuration#resource_id NetworkManagerConnectivityConfiguration#resource_id}.
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/network_manager_connectivity_configuration#resource_type NetworkManagerConnectivityConfiguration#resource_type}.
	ResourceType *string `field:"required" json:"resourceType" yaml:"resourceType"`
}

