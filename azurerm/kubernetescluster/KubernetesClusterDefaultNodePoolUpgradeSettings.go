package kubernetescluster


type KubernetesClusterDefaultNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/kubernetes_cluster#max_surge KubernetesCluster#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

