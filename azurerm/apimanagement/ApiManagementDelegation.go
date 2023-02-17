package apimanagement


type ApiManagementDelegation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#subscriptions_enabled ApiManagement#subscriptions_enabled}.
	SubscriptionsEnabled interface{} `field:"optional" json:"subscriptionsEnabled" yaml:"subscriptionsEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#url ApiManagement#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#user_registration_enabled ApiManagement#user_registration_enabled}.
	UserRegistrationEnabled interface{} `field:"optional" json:"userRegistrationEnabled" yaml:"userRegistrationEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/api_management#validation_key ApiManagement#validation_key}.
	ValidationKey *string `field:"optional" json:"validationKey" yaml:"validationKey"`
}

