package storageaccount


type StorageAccountBlobPropertiesRestorePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#days StorageAccount#days}.
	Days *float64 `field:"required" json:"days" yaml:"days"`
}

