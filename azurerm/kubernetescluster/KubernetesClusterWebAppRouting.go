package kubernetescluster


type KubernetesClusterWebAppRouting struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/kubernetes_cluster#dns_zone_id KubernetesCluster#dns_zone_id}.
	DnsZoneId *string `field:"required" json:"dnsZoneId" yaml:"dnsZoneId"`
}

