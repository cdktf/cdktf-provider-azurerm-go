package synapseworkspace


type SynapseWorkspaceCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/synapse_workspace#key_versionless_id SynapseWorkspace#key_versionless_id}.
	KeyVersionlessId *string `field:"required" json:"keyVersionlessId" yaml:"keyVersionlessId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/synapse_workspace#key_name SynapseWorkspace#key_name}.
	KeyName *string `field:"optional" json:"keyName" yaml:"keyName"`
}

