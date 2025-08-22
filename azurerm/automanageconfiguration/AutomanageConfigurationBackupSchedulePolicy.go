// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration


type AutomanageConfigurationBackupSchedulePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/automanage_configuration#schedule_policy_type AutomanageConfiguration#schedule_policy_type}.
	SchedulePolicyType *string `field:"optional" json:"schedulePolicyType" yaml:"schedulePolicyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/automanage_configuration#schedule_run_days AutomanageConfiguration#schedule_run_days}.
	ScheduleRunDays *[]*string `field:"optional" json:"scheduleRunDays" yaml:"scheduleRunDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/automanage_configuration#schedule_run_frequency AutomanageConfiguration#schedule_run_frequency}.
	ScheduleRunFrequency *string `field:"optional" json:"scheduleRunFrequency" yaml:"scheduleRunFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/automanage_configuration#schedule_run_times AutomanageConfiguration#schedule_run_times}.
	ScheduleRunTimes *[]*string `field:"optional" json:"scheduleRunTimes" yaml:"scheduleRunTimes"`
}

