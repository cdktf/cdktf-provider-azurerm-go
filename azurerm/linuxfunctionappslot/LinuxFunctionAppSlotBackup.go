package linuxfunctionappslot


type LinuxFunctionAppSlotBackup struct {
	// The name which should be used for this Backup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_function_app_slot#name LinuxFunctionAppSlot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_function_app_slot#schedule LinuxFunctionAppSlot#schedule}
	Schedule *LinuxFunctionAppSlotBackupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The SAS URL to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_function_app_slot#storage_account_url LinuxFunctionAppSlot#storage_account_url}
	StorageAccountUrl *string `field:"required" json:"storageAccountUrl" yaml:"storageAccountUrl"`
	// Should this backup job be enabled?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_function_app_slot#enabled LinuxFunctionAppSlot#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}

