package virtualmachine


type VirtualMachineOsProfileLinuxConfigSshKeys struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/virtual_machine#key_data VirtualMachine#key_data}.
	KeyData *string `field:"required" json:"keyData" yaml:"keyData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/virtual_machine#path VirtualMachine#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
}

