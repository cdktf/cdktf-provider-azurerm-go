package storageaccountlocaluser


type StorageAccountLocalUserPermissionScope struct {
	// permissions block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#permissions StorageAccountLocalUser#permissions}
	Permissions *StorageAccountLocalUserPermissionScopePermissions `field:"required" json:"permissions" yaml:"permissions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#resource_name StorageAccountLocalUser#resource_name}.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account_local_user#service StorageAccountLocalUser#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}
