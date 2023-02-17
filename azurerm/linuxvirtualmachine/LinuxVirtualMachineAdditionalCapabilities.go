package linuxvirtualmachine


type LinuxVirtualMachineAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#ultra_ssd_enabled LinuxVirtualMachine#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"optional" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

