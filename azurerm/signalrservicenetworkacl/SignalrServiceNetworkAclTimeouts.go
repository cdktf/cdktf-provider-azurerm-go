package signalrservicenetworkacl


type SignalrServiceNetworkAclTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/signalr_service_network_acl#create SignalrServiceNetworkAcl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/signalr_service_network_acl#delete SignalrServiceNetworkAcl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/signalr_service_network_acl#read SignalrServiceNetworkAcl#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/signalr_service_network_acl#update SignalrServiceNetworkAcl#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

