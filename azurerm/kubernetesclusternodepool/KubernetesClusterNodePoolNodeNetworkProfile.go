package kubernetesclusternodepool


type KubernetesClusterNodePoolNodeNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/kubernetes_cluster_node_pool#node_public_ip_tags KubernetesClusterNodePool#node_public_ip_tags}.
	NodePublicIpTags *map[string]*string `field:"optional" json:"nodePublicIpTags" yaml:"nodePublicIpTags"`
}

