package networkmanagersubscriptionconnection


type NetworkManagerSubscriptionConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_subscription_connection#create NetworkManagerSubscriptionConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_subscription_connection#delete NetworkManagerSubscriptionConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_subscription_connection#read NetworkManagerSubscriptionConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_subscription_connection#update NetworkManagerSubscriptionConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
