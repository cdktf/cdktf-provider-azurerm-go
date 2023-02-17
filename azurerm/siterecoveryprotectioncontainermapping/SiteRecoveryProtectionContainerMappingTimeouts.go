package siterecoveryprotectioncontainermapping


type SiteRecoveryProtectionContainerMappingTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_protection_container_mapping#create SiteRecoveryProtectionContainerMapping#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_protection_container_mapping#delete SiteRecoveryProtectionContainerMapping#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_protection_container_mapping#read SiteRecoveryProtectionContainerMapping#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_protection_container_mapping#update SiteRecoveryProtectionContainerMapping#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

