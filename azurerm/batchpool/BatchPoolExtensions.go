package batchpool


type BatchPoolExtensions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#name BatchPool#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#publisher BatchPool#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#type BatchPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#auto_upgrade_minor_version BatchPool#auto_upgrade_minor_version}.
	AutoUpgradeMinorVersion interface{} `field:"optional" json:"autoUpgradeMinorVersion" yaml:"autoUpgradeMinorVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#protected_settings BatchPool#protected_settings}.
	ProtectedSettings *string `field:"optional" json:"protectedSettings" yaml:"protectedSettings"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#provision_after_extensions BatchPool#provision_after_extensions}.
	ProvisionAfterExtensions *[]*string `field:"optional" json:"provisionAfterExtensions" yaml:"provisionAfterExtensions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#settings_json BatchPool#settings_json}.
	SettingsJson *string `field:"optional" json:"settingsJson" yaml:"settingsJson"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#type_handler_version BatchPool#type_handler_version}.
	TypeHandlerVersion *string `field:"optional" json:"typeHandlerVersion" yaml:"typeHandlerVersion"`
}
