package windowsfunctionapp


type WindowsFunctionAppBackup struct {
	// The name which should be used for this Backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/windows_function_app#name WindowsFunctionApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/windows_function_app#schedule WindowsFunctionApp#schedule}
	Schedule *WindowsFunctionAppBackupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The SAS URL to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/windows_function_app#storage_account_url WindowsFunctionApp#storage_account_url}
	StorageAccountUrl *string `field:"required" json:"storageAccountUrl" yaml:"storageAccountUrl"`
	// Should this backup job be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/windows_function_app#enabled WindowsFunctionApp#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

