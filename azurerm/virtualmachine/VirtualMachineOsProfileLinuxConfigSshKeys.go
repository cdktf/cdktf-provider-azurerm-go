package virtualmachine


type VirtualMachineOsProfileLinuxConfigSshKeys struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#key_data VirtualMachine#key_data}.
	KeyData *string `field:"required" json:"keyData" yaml:"keyData"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#path VirtualMachine#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

