package networkmanager


type NetworkManagerScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_manager#management_group_ids NetworkManager#management_group_ids}.
	ManagementGroupIds *[]*string `field:"optional" json:"managementGroupIds" yaml:"managementGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/network_manager#subscription_ids NetworkManager#subscription_ids}.
	SubscriptionIds *[]*string `field:"optional" json:"subscriptionIds" yaml:"subscriptionIds"`
}

