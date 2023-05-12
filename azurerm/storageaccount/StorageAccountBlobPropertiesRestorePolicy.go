package storageaccount


type StorageAccountBlobPropertiesRestorePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/storage_account#days StorageAccount#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

