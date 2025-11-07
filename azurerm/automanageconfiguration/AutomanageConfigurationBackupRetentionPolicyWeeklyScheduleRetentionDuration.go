// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automanageconfiguration


type AutomanageConfigurationBackupRetentionPolicyWeeklyScheduleRetentionDuration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/automanage_configuration#count AutomanageConfiguration#count}.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/automanage_configuration#duration_type AutomanageConfiguration#duration_type}.
	DurationType *string `field:"optional" json:"durationType" yaml:"durationType"`
}

