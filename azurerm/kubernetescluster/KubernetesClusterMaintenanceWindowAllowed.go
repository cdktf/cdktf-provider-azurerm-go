package kubernetescluster


type KubernetesClusterMaintenanceWindowAllowed struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#day KubernetesCluster#day}.
	Day *string `field:"required" json:"day" yaml:"day"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#hours KubernetesCluster#hours}.
	Hours *[]*float64 `field:"required" json:"hours" yaml:"hours"`
}

