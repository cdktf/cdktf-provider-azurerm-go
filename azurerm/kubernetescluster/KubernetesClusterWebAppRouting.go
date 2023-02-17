package kubernetescluster


type KubernetesClusterWebAppRouting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#dns_zone_id KubernetesCluster#dns_zone_id}.
	DnsZoneId *string `field:"required" json:"dnsZoneId" yaml:"dnsZoneId"`
}

