package eventhubnamespace


type EventhubNamespaceNetworkRulesetsVirtualNetworkRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/eventhub_namespace#ignore_missing_virtual_network_service_endpoint EventhubNamespace#ignore_missing_virtual_network_service_endpoint}.
	IgnoreMissingVirtualNetworkServiceEndpoint interface{} `field:"optional" json:"ignoreMissingVirtualNetworkServiceEndpoint" yaml:"ignoreMissingVirtualNetworkServiceEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/eventhub_namespace#subnet_id EventhubNamespace#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

