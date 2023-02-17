package kubernetescluster


type KubernetesClusterKeyManagementService struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#key_vault_key_id KubernetesCluster#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#key_vault_network_access KubernetesCluster#key_vault_network_access}.
	KeyVaultNetworkAccess *string `field:"optional" json:"keyVaultNetworkAccess" yaml:"keyVaultNetworkAccess"`
}

