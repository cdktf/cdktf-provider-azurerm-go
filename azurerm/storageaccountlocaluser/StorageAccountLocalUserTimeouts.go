package storageaccountlocaluser


type StorageAccountLocalUserTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#create StorageAccountLocalUser#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#delete StorageAccountLocalUser#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#read StorageAccountLocalUser#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#update StorageAccountLocalUser#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
