package kubernetescluster


type KubernetesClusterLinuxProfileSshKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#key_data KubernetesCluster#key_data}.
	KeyData *string `field:"required" json:"keyData" yaml:"keyData"`
}

