package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#tag LinuxVirtualMachineScaleSet#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#type LinuxVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

