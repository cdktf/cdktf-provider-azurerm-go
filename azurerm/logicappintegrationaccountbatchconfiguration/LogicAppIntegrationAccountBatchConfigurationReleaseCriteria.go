package logicappintegrationaccountbatchconfiguration


type LogicAppIntegrationAccountBatchConfigurationReleaseCriteria struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_batch_configuration#batch_size LogicAppIntegrationAccountBatchConfiguration#batch_size}.
	BatchSize *float64 `field:"optional" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_batch_configuration#message_count LogicAppIntegrationAccountBatchConfiguration#message_count}.
	MessageCount *float64 `field:"optional" json:"messageCount" yaml:"messageCount"`
	// recurrence block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_batch_configuration#recurrence LogicAppIntegrationAccountBatchConfiguration#recurrence}
	Recurrence *LogicAppIntegrationAccountBatchConfigurationReleaseCriteriaRecurrence `field:"optional" json:"recurrence" yaml:"recurrence"`
}
