package linuxvirtualmachine


type LinuxVirtualMachineOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/linux_virtual_machine#option LinuxVirtualMachine#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/linux_virtual_machine#placement LinuxVirtualMachine#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}

