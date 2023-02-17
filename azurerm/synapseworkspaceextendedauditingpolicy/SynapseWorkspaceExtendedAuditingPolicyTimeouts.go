package synapseworkspaceextendedauditingpolicy


type SynapseWorkspaceExtendedAuditingPolicyTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace_extended_auditing_policy#create SynapseWorkspaceExtendedAuditingPolicy#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace_extended_auditing_policy#delete SynapseWorkspaceExtendedAuditingPolicy#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace_extended_auditing_policy#read SynapseWorkspaceExtendedAuditingPolicy#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_workspace_extended_auditing_policy#update SynapseWorkspaceExtendedAuditingPolicy#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

