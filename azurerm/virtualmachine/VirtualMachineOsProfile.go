package virtualmachine


type VirtualMachineOsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/virtual_machine#admin_username VirtualMachine#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/virtual_machine#computer_name VirtualMachine#computer_name}.
	ComputerName *string `field:"required" json:"computerName" yaml:"computerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/virtual_machine#admin_password VirtualMachine#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/virtual_machine#custom_data VirtualMachine#custom_data}.
	CustomData *string `field:"optional" json:"customData" yaml:"customData"`
}

