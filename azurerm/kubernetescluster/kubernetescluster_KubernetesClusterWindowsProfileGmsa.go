package kubernetescluster


type KubernetesClusterWindowsProfileGmsa struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#dns_server KubernetesCluster#dns_server}.
	DnsServer *string `field:"required" json:"dnsServer" yaml:"dnsServer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#root_domain KubernetesCluster#root_domain}.
	RootDomain *string `field:"required" json:"rootDomain" yaml:"rootDomain"`
}

