package virtualmachine


type VirtualMachineAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#ultra_ssd_enabled VirtualMachine#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"required" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

