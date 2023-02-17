package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetScaleIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#force_deletion_enabled LinuxVirtualMachineScaleSet#force_deletion_enabled}.
	ForceDeletionEnabled interface{} `field:"optional" json:"forceDeletionEnabled" yaml:"forceDeletionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#rule LinuxVirtualMachineScaleSet#rule}.
	Rule *string `field:"optional" json:"rule" yaml:"rule"`
}

