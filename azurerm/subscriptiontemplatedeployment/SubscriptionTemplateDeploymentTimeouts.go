package subscriptiontemplatedeployment


type SubscriptionTemplateDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_template_deployment#create SubscriptionTemplateDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_template_deployment#delete SubscriptionTemplateDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_template_deployment#read SubscriptionTemplateDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/subscription_template_deployment#update SubscriptionTemplateDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

