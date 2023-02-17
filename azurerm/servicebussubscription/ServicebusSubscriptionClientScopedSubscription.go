package servicebussubscription


type ServicebusSubscriptionClientScopedSubscription struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#client_id ServicebusSubscription#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/servicebus_subscription#is_client_scoped_subscription_shareable ServicebusSubscription#is_client_scoped_subscription_shareable}.
	IsClientScopedSubscriptionShareable interface{} `field:"optional" json:"isClientScopedSubscriptionShareable" yaml:"isClientScopedSubscriptionShareable"`
}

