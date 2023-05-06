package eventhubnamespace


type EventhubNamespaceNetworkRulesetsIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/eventhub_namespace#action EventhubNamespace#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/eventhub_namespace#ip_mask EventhubNamespace#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}

