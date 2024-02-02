// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagemanagementpolicy


type StorageManagementPolicyRuleActionsVersion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/storage_management_policy#change_tier_to_archive_after_days_since_creation StorageManagementPolicy#change_tier_to_archive_after_days_since_creation}.
	ChangeTierToArchiveAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToArchiveAfterDaysSinceCreation" yaml:"changeTierToArchiveAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/storage_management_policy#change_tier_to_cool_after_days_since_creation StorageManagementPolicy#change_tier_to_cool_after_days_since_creation}.
	ChangeTierToCoolAfterDaysSinceCreation *float64 `field:"optional" json:"changeTierToCoolAfterDaysSinceCreation" yaml:"changeTierToCoolAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/storage_management_policy#delete_after_days_since_creation StorageManagementPolicy#delete_after_days_since_creation}.
	DeleteAfterDaysSinceCreation *float64 `field:"optional" json:"deleteAfterDaysSinceCreation" yaml:"deleteAfterDaysSinceCreation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/storage_management_policy#tier_to_archive_after_days_since_last_tier_change_greater_than StorageManagementPolicy#tier_to_archive_after_days_since_last_tier_change_greater_than}.
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan *float64 `field:"optional" json:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan" yaml:"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/storage_management_policy#tier_to_cold_after_days_since_creation_greater_than StorageManagementPolicy#tier_to_cold_after_days_since_creation_greater_than}.
	TierToColdAfterDaysSinceCreationGreaterThan *float64 `field:"optional" json:"tierToColdAfterDaysSinceCreationGreaterThan" yaml:"tierToColdAfterDaysSinceCreationGreaterThan"`
}

