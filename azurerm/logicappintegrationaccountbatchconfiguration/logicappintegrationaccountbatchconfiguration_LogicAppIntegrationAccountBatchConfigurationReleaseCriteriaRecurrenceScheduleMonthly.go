package logicappintegrationaccountbatchconfiguration


type LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrenceScheduleMonthly struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_batch_configuration#week LogicAppIntegrationAccountBatchConfiguration#week}.
	Week *float64 `field:"required" json:"week" yaml:"week"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_batch_configuration#weekday LogicAppIntegrationAccountBatchConfiguration#weekday}.
	Weekday *string `field:"required" json:"weekday" yaml:"weekday"`
}

