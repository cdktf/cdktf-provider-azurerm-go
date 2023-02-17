package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#disable_automatic_rollback LinuxVirtualMachineScaleSet#disable_automatic_rollback}.
	DisableAutomaticRollback interface{} `field:"required" json:"disableAutomaticRollback" yaml:"disableAutomaticRollback"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#enable_automatic_os_upgrade LinuxVirtualMachineScaleSet#enable_automatic_os_upgrade}.
	EnableAutomaticOsUpgrade interface{} `field:"required" json:"enableAutomaticOsUpgrade" yaml:"enableAutomaticOsUpgrade"`
}

