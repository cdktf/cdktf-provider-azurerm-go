package kubernetescluster


type KubernetesClusterDefaultNodePoolNodeNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#node_public_ip_tags KubernetesCluster#node_public_ip_tags}.
	NodePublicIpTags *map[string]*string `field:"optional" json:"nodePublicIpTags" yaml:"nodePublicIpTags"`
}

