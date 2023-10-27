// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabricmanagedcluster


type ServiceFabricManagedClusterNodeTypeVmSecretsCertificates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/service_fabric_managed_cluster#store ServiceFabricManagedCluster#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/service_fabric_managed_cluster#url ServiceFabricManagedCluster#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}

