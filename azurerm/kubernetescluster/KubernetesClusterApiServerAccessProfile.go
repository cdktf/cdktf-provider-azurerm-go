package kubernetescluster


type KubernetesClusterApiServerAccessProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#authorized_ip_ranges KubernetesCluster#authorized_ip_ranges}.
	AuthorizedIpRanges *[]*string `field:"optional" json:"authorizedIpRanges" yaml:"authorizedIpRanges"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_id KubernetesCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#vnet_integration_enabled KubernetesCluster#vnet_integration_enabled}.
	VnetIntegrationEnabled interface{} `field:"optional" json:"vnetIntegrationEnabled" yaml:"vnetIntegrationEnabled"`
}

