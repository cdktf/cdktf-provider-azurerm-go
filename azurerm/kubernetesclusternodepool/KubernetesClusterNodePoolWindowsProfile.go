package kubernetesclusternodepool


type KubernetesClusterNodePoolWindowsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/kubernetes_cluster_node_pool#outbound_nat_enabled KubernetesClusterNodePool#outbound_nat_enabled}.
	OutboundNatEnabled interface{} `field:"optional" json:"outboundNatEnabled" yaml:"outboundNatEnabled"`
}

