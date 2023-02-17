package storageaccountlocaluser


type StorageAccountLocalUserSshAuthorizedKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#key StorageAccountLocalUser#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#description StorageAccountLocalUser#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

