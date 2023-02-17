package virtualmachine


type VirtualMachineOsProfileWindowsConfig struct {
	// additional_unattend_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#additional_unattend_config VirtualMachine#additional_unattend_config}
	AdditionalUnattendConfig interface{} `field:"optional" json:"additionalUnattendConfig" yaml:"additionalUnattendConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#enable_automatic_upgrades VirtualMachine#enable_automatic_upgrades}.
	EnableAutomaticUpgrades interface{} `field:"optional" json:"enableAutomaticUpgrades" yaml:"enableAutomaticUpgrades"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#provision_vm_agent VirtualMachine#provision_vm_agent}.
	ProvisionVmAgent interface{} `field:"optional" json:"provisionVmAgent" yaml:"provisionVmAgent"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#timezone VirtualMachine#timezone}.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// winrm block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#winrm VirtualMachine#winrm}
	Winrm interface{} `field:"optional" json:"winrm" yaml:"winrm"`
}

