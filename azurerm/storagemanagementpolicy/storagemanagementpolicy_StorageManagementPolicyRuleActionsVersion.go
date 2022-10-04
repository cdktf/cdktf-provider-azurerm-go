package storagemanagementpolicy


type StorageManagementPolicyRuleActionsVersion struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#change_tier_to_archive_after_days_since_creation StorageManagementPolicy#change_tier_to_archive_after_days_since_creation}.
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToArchiveAfterDaysSinceCreation" yaml:"changeTierToArchiveAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#change_tier_to_cool_after_days_since_creation StorageManagementPolicy#change_tier_to_cool_after_days_since_creation}.
	ChangeTierToCoolAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToCoolAfterDaysSinceCreation" yaml:"changeTierToCoolAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_creation StorageManagementPolicy#delete_after_days_since_creation}.
	DeleteAfterDaysSinceCreation *float64 `field:"optional" json:"deleteAfterDaysSinceCreation" yaml:"deleteAfterDaysSinceCreation"`
}

