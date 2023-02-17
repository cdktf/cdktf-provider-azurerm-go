package mobilenetworksimgroup


type MobileNetworkSimGroupIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_sim_group#type MobileNetworkSimGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mobile_network_sim_group#identity_ids MobileNetworkSimGroup#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

