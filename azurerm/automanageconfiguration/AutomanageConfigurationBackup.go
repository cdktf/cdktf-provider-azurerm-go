// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration


type AutomanageConfigurationBackup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/automanage_configuration#instant_rp_retention_range_in_days AutomanageConfiguration#instant_rp_retention_range_in_days}.
	InstantRpRetentionRangeInDays *float64 `field:"optional" json:"instantRpRetentionRangeInDays" yaml:"instantRpRetentionRangeInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/automanage_configuration#policy_name AutomanageConfiguration#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/automanage_configuration#retention_policy AutomanageConfiguration#retention_policy}
	RetentionPolicy *AutomanageConfigurationBackupRetentionPolicy `field:"optional" json:"retentionPolicy" yaml:"retentionPolicy"`
	// schedule_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/automanage_configuration#schedule_policy AutomanageConfiguration#schedule_policy}
	SchedulePolicy *AutomanageConfigurationBackupSchedulePolicy `field:"optional" json:"schedulePolicy" yaml:"schedulePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/automanage_configuration#time_zone AutomanageConfiguration#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

