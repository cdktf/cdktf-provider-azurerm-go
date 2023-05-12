package kubernetescluster


type KubernetesClusterWindowsProfileGmsa struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/kubernetes_cluster#dns_server KubernetesCluster#dns_server}.
	DnsServer *string `field:"required" json:"dnsServer" yaml:"dnsServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/kubernetes_cluster#root_domain KubernetesCluster#root_domain}.
	RootDomain *string `field:"required" json:"rootDomain" yaml:"rootDomain"`
}

