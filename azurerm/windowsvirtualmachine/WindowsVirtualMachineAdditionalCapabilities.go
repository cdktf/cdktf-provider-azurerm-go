package windowsvirtualmachine


type WindowsVirtualMachineAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#ultra_ssd_enabled WindowsVirtualMachine#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

