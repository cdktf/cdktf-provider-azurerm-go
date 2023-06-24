package storageaccount


type StorageAccountBlobPropertiesContainerDeleteRetentionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/storage_account#days StorageAccount#days}.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

