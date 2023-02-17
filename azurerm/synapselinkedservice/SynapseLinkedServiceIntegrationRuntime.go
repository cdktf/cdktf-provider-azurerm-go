package synapselinkedservice


type SynapseLinkedServiceIntegrationRuntime struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_linked_service#name SynapseLinkedService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_linked_service#parameters SynapseLinkedService#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
}

