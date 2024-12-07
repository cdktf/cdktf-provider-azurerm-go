// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterKeyManagementService struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/kubernetes_cluster#key_vault_key_id KubernetesCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/kubernetes_cluster#key_vault_network_access KubernetesCluster#key_vault_network_access}.
	KeyVaultNetworkAccess *string `field:"optional" json:"keyVaultNetworkAccess" yaml:"keyVaultNetworkAccess"`
}

