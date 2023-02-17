package storageaccount


type StorageAccountAzureFilesAuthentication struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#directory_type StorageAccount#directory_type}.
	DirectoryType *string `field:"required" json:"directoryType" yaml:"directoryType"`
	// active_directory block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#active_directory StorageAccount#active_directory}
	ActiveDirectory *StorageAccountAzureFilesAuthenticationActiveDirectory `field:"optional" json:"activeDirectory" yaml:"activeDirectory"`
}

