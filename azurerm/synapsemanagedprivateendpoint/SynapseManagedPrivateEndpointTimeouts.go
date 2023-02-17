package synapsemanagedprivateendpoint


type SynapseManagedPrivateEndpointTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_managed_private_endpoint#create SynapseManagedPrivateEndpoint#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_managed_private_endpoint#delete SynapseManagedPrivateEndpoint#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_managed_private_endpoint#read SynapseManagedPrivateEndpoint#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

