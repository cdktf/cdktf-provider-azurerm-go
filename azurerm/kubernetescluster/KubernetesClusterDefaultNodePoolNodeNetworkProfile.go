package kubernetescluster


type KubernetesClusterDefaultNodePoolNodeNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/kubernetes_cluster#node_public_ip_tags KubernetesCluster#node_public_ip_tags}.
	NodePublicIpTags *map[string]*string `field:"optional" json:"nodePublicIpTags" yaml:"nodePublicIpTags"`
}

