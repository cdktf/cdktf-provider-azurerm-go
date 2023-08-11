package linuxvirtualmachine


type LinuxVirtualMachineTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_virtual_machine#create LinuxVirtualMachine#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_virtual_machine#delete LinuxVirtualMachine#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_virtual_machine#read LinuxVirtualMachine#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/linux_virtual_machine#update LinuxVirtualMachine#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

