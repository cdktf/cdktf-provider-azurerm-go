// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountbatchconfiguration


type LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceSchedule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/logic_app_integration_account_batch_configuration#hours LogicAppIntegrationAccountBatchConfiguration#hours}.
	Hours *[]*float64 `field:"optional" json:"hours" yaml:"hours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/logic_app_integration_account_batch_configuration#minutes LogicAppIntegrationAccountBatchConfiguration#minutes}.
	Minutes *[]*float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/logic_app_integration_account_batch_configuration#month_days LogicAppIntegrationAccountBatchConfiguration#month_days}.
	MonthDays *[]*float64 `field:"optional" json:"monthDays" yaml:"monthDays"`
	// monthly block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/logic_app_integration_account_batch_configuration#monthly LogicAppIntegrationAccountBatchConfiguration#monthly}
	Monthly interface{} `field:"optional" json:"monthly" yaml:"monthly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.28.0/docs/resources/logic_app_integration_account_batch_configuration#week_days LogicAppIntegrationAccountBatchConfiguration#week_days}.
	WeekDays *[]*string `field:"optional" json:"weekDays" yaml:"weekDays"`
}

