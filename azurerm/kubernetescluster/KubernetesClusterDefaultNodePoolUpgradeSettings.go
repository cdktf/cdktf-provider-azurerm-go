package kubernetescluster


type KubernetesClusterDefaultNodePoolUpgradeSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#max_surge KubernetesCluster#max_surge}.
	MaxSurge *string `field:"required" json:"maxSurge" yaml:"maxSurge"`
}

