package virtualmachine


type VirtualMachineOsProfileWindowsConfigAdditionalUnattendConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#component VirtualMachine#component}.
	Component *string `field:"required" json:"component" yaml:"component"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#content VirtualMachine#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#pass VirtualMachine#pass}.
	Pass *string `field:"required" json:"pass" yaml:"pass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#setting_name VirtualMachine#setting_name}.
	SettingName *string `field:"required" json:"settingName" yaml:"settingName"`
}

