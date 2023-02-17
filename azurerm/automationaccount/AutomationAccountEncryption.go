package automationaccount


type AutomationAccountEncryption struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_account#key_vault_key_id AutomationAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_account#key_source AutomationAccount#key_source}.
	KeySource *string `field:"optional" json:"keySource" yaml:"keySource"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/automation_account#user_assigned_identity_id AutomationAccount#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

