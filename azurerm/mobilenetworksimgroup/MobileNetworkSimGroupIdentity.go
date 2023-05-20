package mobilenetworksimgroup


type MobileNetworkSimGroupIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/mobile_network_sim_group#identity_ids MobileNetworkSimGroup#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/mobile_network_sim_group#type MobileNetworkSimGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

