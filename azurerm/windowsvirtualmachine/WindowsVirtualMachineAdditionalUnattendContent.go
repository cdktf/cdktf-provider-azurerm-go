package windowsvirtualmachine


type WindowsVirtualMachineAdditionalUnattendContent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/windows_virtual_machine#content WindowsVirtualMachine#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/windows_virtual_machine#setting WindowsVirtualMachine#setting}.
	Setting *string `field:"required" json:"setting" yaml:"setting"`
}

