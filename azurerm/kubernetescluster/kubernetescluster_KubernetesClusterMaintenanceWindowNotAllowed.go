package kubernetescluster


type KubernetesClusterMaintenanceWindowNotAllowed struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#end KubernetesCluster#end}.
	End *string `field:"required" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#start KubernetesCluster#start}.
	Start *string `field:"required" json:"start" yaml:"start"`
}

