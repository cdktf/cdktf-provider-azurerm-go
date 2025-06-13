// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabricmanagedcluster


type ServiceFabricManagedClusterNodeTypeVmSecrets struct {
	// certificates block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/service_fabric_managed_cluster#certificates ServiceFabricManagedCluster#certificates}
	Certificates interface{} `field:"required" json:"certificates" yaml:"certificates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/service_fabric_managed_cluster#vault_id ServiceFabricManagedCluster#vault_id}.
	VaultId *string `field:"required" json:"vaultId" yaml:"vaultId"`
}

