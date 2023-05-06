package virtualmachine


type VirtualMachineAdditionalCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/virtual_machine#ultra_ssd_enabled VirtualMachine#ultra_ssd_enabled}.
	UltraSsdEnabled interface{} `field:"required" json:"ultraSsdEnabled" yaml:"ultraSsdEnabled"`
}

