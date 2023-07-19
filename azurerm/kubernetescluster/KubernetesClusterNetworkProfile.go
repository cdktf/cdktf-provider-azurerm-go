package kubernetescluster


type KubernetesClusterNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#network_plugin KubernetesCluster#network_plugin}.
	NetworkPlugin *string `field:"required" json:"networkPlugin" yaml:"networkPlugin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#dns_service_ip KubernetesCluster#dns_service_ip}.
	DnsServiceIp *string `field:"optional" json:"dnsServiceIp" yaml:"dnsServiceIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#docker_bridge_cidr KubernetesCluster#docker_bridge_cidr}.
	DockerBridgeCidr *string `field:"optional" json:"dockerBridgeCidr" yaml:"dockerBridgeCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#ebpf_data_plane KubernetesCluster#ebpf_data_plane}.
	EbpfDataPlane *string `field:"optional" json:"ebpfDataPlane" yaml:"ebpfDataPlane"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#ip_versions KubernetesCluster#ip_versions}.
	IpVersions *[]*string `field:"optional" json:"ipVersions" yaml:"ipVersions"`
	// load_balancer_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#load_balancer_profile KubernetesCluster#load_balancer_profile}
	LoadBalancerProfile *KubernetesClusterNetworkProfileLoadBalancerProfile `field:"optional" json:"loadBalancerProfile" yaml:"loadBalancerProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#load_balancer_sku KubernetesCluster#load_balancer_sku}.
	LoadBalancerSku *string `field:"optional" json:"loadBalancerSku" yaml:"loadBalancerSku"`
	// nat_gateway_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#nat_gateway_profile KubernetesCluster#nat_gateway_profile}
	NatGatewayProfile *KubernetesClusterNetworkProfileNatGatewayProfile `field:"optional" json:"natGatewayProfile" yaml:"natGatewayProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#network_mode KubernetesCluster#network_mode}.
	NetworkMode *string `field:"optional" json:"networkMode" yaml:"networkMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#network_plugin_mode KubernetesCluster#network_plugin_mode}.
	NetworkPluginMode *string `field:"optional" json:"networkPluginMode" yaml:"networkPluginMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#network_policy KubernetesCluster#network_policy}.
	NetworkPolicy *string `field:"optional" json:"networkPolicy" yaml:"networkPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#outbound_type KubernetesCluster#outbound_type}.
	OutboundType *string `field:"optional" json:"outboundType" yaml:"outboundType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#pod_cidr KubernetesCluster#pod_cidr}.
	PodCidr *string `field:"optional" json:"podCidr" yaml:"podCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#pod_cidrs KubernetesCluster#pod_cidrs}.
	PodCidrs *[]*string `field:"optional" json:"podCidrs" yaml:"podCidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#service_cidr KubernetesCluster#service_cidr}.
	ServiceCidr *string `field:"optional" json:"serviceCidr" yaml:"serviceCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/kubernetes_cluster#service_cidrs KubernetesCluster#service_cidrs}.
	ServiceCidrs *[]*string `field:"optional" json:"serviceCidrs" yaml:"serviceCidrs"`
}

