package storageaccountlocaluser


type StorageAccountLocalUserPermissionScopePermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#create StorageAccountLocalUser#create}.
	Create interface{} `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#delete StorageAccountLocalUser#delete}.
	Delete interface{} `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#list StorageAccountLocalUser#list}.
	List interface{} `field:"optional" json:"list" yaml:"list"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#read StorageAccountLocalUser#read}.
	Read interface{} `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#write StorageAccountLocalUser#write}.
	Write interface{} `field:"optional" json:"write" yaml:"write"`
}
