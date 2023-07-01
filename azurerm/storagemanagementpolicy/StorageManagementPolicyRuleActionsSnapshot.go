package storagemanagementpolicy


type StorageManagementPolicyRuleActionsSnapshot struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/storage_management_policy#change_tier_to_archive_after_days_since_creation StorageManagementPolicy#change_tier_to_archive_after_days_since_creation}.
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToArchiveAfterDaysSinceCreation" yaml:"changeTierToArchiveAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/storage_management_policy#change_tier_to_cool_after_days_since_creation StorageManagementPolicy#change_tier_to_cool_after_days_since_creation}.
	ChangeTierToCoolAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToCoolAfterDaysSinceCreation" yaml:"changeTierToCoolAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/storage_management_policy#delete_after_days_since_creation_greater_than StorageManagementPolicy#delete_after_days_since_creation_greater_than}.
	DeleteAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceCreationGreaterThan" yaml:"deleteAfterDaysSinceCreationGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/storage_management_policy#tier_to_archive_after_days_since_last_tier_change_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_tier_change_greater_than}.
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan"`
}

