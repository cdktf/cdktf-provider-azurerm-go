package fluidrelayserver


type FluidRelayServerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/fluid_relay_server#create FluidRelayServer#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/fluid_relay_server#delete FluidRelayServer#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/fluid_relay_server#read FluidRelayServer#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/fluid_relay_server#update FluidRelayServer#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

