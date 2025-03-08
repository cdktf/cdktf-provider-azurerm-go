// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountbatchconfiguration


type LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthly struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/logic_app_integration_account_batch_configuration#week LogicAppIntegrationAccountBatchConfiguration#week}.
	Week *float64 `field:"required" json:"week" yaml:"week"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/logic_app_integration_account_batch_configuration#weekday LogicAppIntegrationAccountBatchConfiguration#weekday}.
	Weekday *string `field:"required" json:"weekday" yaml:"weekday"`
}

