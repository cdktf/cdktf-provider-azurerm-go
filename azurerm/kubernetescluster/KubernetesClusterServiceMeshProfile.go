package kubernetescluster


type KubernetesClusterServiceMeshProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/kubernetes_cluster#mode KubernetesCluster#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
}

