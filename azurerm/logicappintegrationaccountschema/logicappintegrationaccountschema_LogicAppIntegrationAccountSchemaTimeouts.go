package logicappintegrationaccountschema


type LogicAppIntegrationAccountSchemaTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_schema#create LogicAppIntegrationAccountSchema#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_schema#delete LogicAppIntegrationAccountSchema#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_schema#read LogicAppIntegrationAccountSchema#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_schema#update LogicAppIntegrationAccountSchema#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
