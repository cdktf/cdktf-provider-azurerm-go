package siterecoveryprotectioncontainermapping


type SiteRecoveryProtectionContainerMappingAutomaticUpdate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/site_recovery_protection_container_mapping#automation_account_id SiteRecoveryProtectionContainerMapping#automation_account_id}.
	AutomationAccountId *string `field:"optional" json:"automationAccountId" yaml:"automationAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/site_recovery_protection_container_mapping#enabled SiteRecoveryProtectionContainerMapping#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

