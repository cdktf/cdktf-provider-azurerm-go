package storageaccount


type StorageAccountSasPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#expiration_period StorageAccount#expiration_period}.
	ExpirationPeriod *string `field:"required" json:"expirationPeriod" yaml:"expirationPeriod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#expiration_action StorageAccount#expiration_action}.
	ExpirationAction *string `field:"optional" json:"expirationAction" yaml:"expirationAction"`
}

