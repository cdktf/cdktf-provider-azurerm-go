package kubernetescluster


type KubernetesClusterAciConnectorLinux struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_name KubernetesCluster#subnet_name}.
	SubnetName *string `field:"required" json:"subnetName" yaml:"subnetName"`
}

