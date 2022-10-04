package storageaccount


type StorageAccountRouting struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#choice StorageAccount#choice}.
	Choice *string `field:"optional" json:"choice" yaml:"choice"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#publish_internet_endpoints StorageAccount#publish_internet_endpoints}.
	PublishInternetEndpoints interface{} `field:"optional" json:"publishInternetEndpoints" yaml:"publishInternetEndpoints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#publish_microsoft_endpoints StorageAccount#publish_microsoft_endpoints}.
	PublishMicrosoftEndpoints interface{} `field:"optional" json:"publishMicrosoftEndpoints" yaml:"publishMicrosoftEndpoints"`
}

