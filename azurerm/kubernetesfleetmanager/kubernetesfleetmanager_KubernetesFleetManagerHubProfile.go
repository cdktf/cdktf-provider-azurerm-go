package kubernetesfleetmanager


type KubernetesFleetManagerHubProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_fleet_manager#dns_prefix KubernetesFleetManager#dns_prefix}.
	DnsPrefix *string `field:"required" json:"dnsPrefix" yaml:"dnsPrefix"`
}

