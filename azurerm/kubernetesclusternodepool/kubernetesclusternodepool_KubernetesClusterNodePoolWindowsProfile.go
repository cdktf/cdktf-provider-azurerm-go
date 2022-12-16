package kubernetesclusternodepool


type KubernetesClusterNodePoolWindowsProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster_node_pool#outbound_nat_enabled KubernetesClusterNodePool#outbound_nat_enabled}.
	OutboundNatEnabled interface{} `field:"optional" json:"outboundNatEnabled" yaml:"outboundNatEnabled"`
}

