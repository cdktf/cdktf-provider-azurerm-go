package storageaccount


type StorageAccountNetworkRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/storage_account#default_action StorageAccount#default_action}.
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/storage_account#bypass StorageAccount#bypass}.
	Bypass *[]*string `field:"optional" json:"bypass" yaml:"bypass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/storage_account#ip_rules StorageAccount#ip_rules}.
	IpRules *[]*string `field:"optional" json:"ipRules" yaml:"ipRules"`
	// private_link_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/storage_account#private_link_access StorageAccount#private_link_access}
	PrivateLinkAccess interface{} `field:"optional" json:"privateLinkAccess" yaml:"privateLinkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/storage_account#virtual_network_subnet_ids StorageAccount#virtual_network_subnet_ids}.
	VirtualNetworkSubnetIds *[]*string `field:"optional" json:"virtualNetworkSubnetIds" yaml:"virtualNetworkSubnetIds"`
}

