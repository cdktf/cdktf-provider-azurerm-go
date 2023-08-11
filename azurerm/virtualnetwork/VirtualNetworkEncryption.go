package virtualnetwork


type VirtualNetworkEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_network#enforcement VirtualNetwork#enforcement}.
	Enforcement *string `field:"required" json:"enforcement" yaml:"enforcement"`
}

