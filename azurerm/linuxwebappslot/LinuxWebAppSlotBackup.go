package linuxwebappslot


type LinuxWebAppSlotBackup struct {
	// The name which should be used for this Backup.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#name LinuxWebAppSlot#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedule block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#schedule LinuxWebAppSlot#schedule}
	Schedule *LinuxWebAppSlotBackupSchedule `field:"required" json:"schedule" yaml:"schedule"`
	// The SAS URL to the container.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#storage_account_url LinuxWebAppSlot#storage_account_url}
	StorageAccountUrl *string `field:"required" json:"storageAccountUrl" yaml:"storageAccountUrl"`
	// Should this backup job be enabled?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app_slot#enabled LinuxWebAppSlot#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
}
