package logicappintegrationaccountagreement


type LogicAppIntegrationAccountAgreementGuestIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_agreement#qualifier LogicAppIntegrationAccountAgreement#qualifier}.
	Qualifier *string `field:"required" json:"qualifier" yaml:"qualifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_integration_account_agreement#value LogicAppIntegrationAccountAgreement#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

