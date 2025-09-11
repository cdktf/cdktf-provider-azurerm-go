// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvault


type RecoveryServicesVaultMonitoring struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/recovery_services_vault#alerts_for_all_job_failures_enabled RecoveryServicesVault#alerts_for_all_job_failures_enabled}.
	AlertsForAllJobFailuresEnabled interface{} `field:"optional" json:"alertsForAllJobFailuresEnabled" yaml:"alertsForAllJobFailuresEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/recovery_services_vault#alerts_for_critical_operation_failures_enabled RecoveryServicesVault#alerts_for_critical_operation_failures_enabled}.
	AlertsForCriticalOperationFailuresEnabled interface{} `field:"optional" json:"alertsForCriticalOperationFailuresEnabled" yaml:"alertsForCriticalOperationFailuresEnabled"`
}

