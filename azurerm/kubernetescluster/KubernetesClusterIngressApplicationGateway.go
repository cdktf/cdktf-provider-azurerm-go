package kubernetescluster


type KubernetesClusterIngressApplicationGateway struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#gateway_id KubernetesCluster#gateway_id}.
	GatewayId *string `field:"optional" json:"gatewayId" yaml:"gatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#gateway_name KubernetesCluster#gateway_name}.
	GatewayName *string `field:"optional" json:"gatewayName" yaml:"gatewayName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_cidr KubernetesCluster#subnet_cidr}.
	SubnetCidr *string `field:"optional" json:"subnetCidr" yaml:"subnetCidr"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_cluster#subnet_id KubernetesCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

