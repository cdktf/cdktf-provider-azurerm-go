package linuxvirtualmachine


type LinuxVirtualMachineTerminationNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#enabled LinuxVirtualMachine#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#timeout LinuxVirtualMachine#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

