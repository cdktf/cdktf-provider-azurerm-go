package kubernetesclusternodepool


type KubernetesClusterNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#max_surge KubernetesClusterNodePool#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

