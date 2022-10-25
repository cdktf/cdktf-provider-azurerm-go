package storagemanagementpolicy


type StorageManagementPolicyRuleActionsBaseBlob struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_last_access_time_greater_than StorageManagementPolicy#delete_after_days_since_last_access_time_greater_than}.
	DeleteAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceLastAccessTimeGreaterThan" yaml:"deleteAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#delete_after_days_since_modification_greater_than StorageManagementPolicy#delete_after_days_since_modification_greater_than}.
	DeleteAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"deleteAfterDaysSinceModificationGreaterThan" yaml:"deleteAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_archive_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_access_time_greater_than}.
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_archive_after_days_since_last_tier_change_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_tier_change_greater_than}.
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_archive_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_modification_greater_than}.
	TierToArchiveAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceModificationGreaterThan" yaml:"tierToArchiveAfterDaysSinceModificationGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_cool_after_days_since_last_access_time_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_last_access_time_greater_than}.
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan" yaml:"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_management_policy#tier_to_cool_after_days_since_modification_greater_than StorageManagementPolicy#tier_to_cool_after_days_since_modification_greater_than}.
	TierToCoolAfterDaysSinceModificationGreaterThan *float64 `field:"optional" json:"tierToCoolAfterDaysSinceModificationGreaterThan" yaml:"tierToCoolAfterDaysSinceModificationGreaterThan"`
}

