package windowsvirtualmachine


type WindowsVirtualMachineTerminationNotification struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#enabled WindowsVirtualMachine#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#timeout WindowsVirtualMachine#timeout}.
	Timeout *string `field:"optional" json:"timeout" yaml:"timeout"`
}

