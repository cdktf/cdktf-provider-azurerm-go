package logicappintegrationaccountsession


type LogicAppIntegrationAccountSessionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_session#create LogicAppIntegrationAccountSession#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_session#delete LogicAppIntegrationAccountSession#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_session#read LogicAppIntegrationAccountSession#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_session#update LogicAppIntegrationAccountSession#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

