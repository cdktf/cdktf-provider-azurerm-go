package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetScaleIn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#force_deletion_enabled WindowsVirtualMachineScaleSet#force_deletion_enabled}.
	ForceDeletionEnabled interface{} `field:"optional" json:"forceDeletionEnabled" yaml:"forceDeletionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#rule WindowsVirtualMachineScaleSet#rule}.
	Rule *string `field:"optional" json:"rule" yaml:"rule"`
}
