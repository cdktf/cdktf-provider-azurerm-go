package policyvirtualmachineconfigurationassignment


type PolicyVirtualMachineConfigurationAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/policy_virtual_machine_configuration_assignment#create PolicyVirtualMachineConfigurationAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/policy_virtual_machine_configuration_assignment#delete PolicyVirtualMachineConfigurationAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/policy_virtual_machine_configuration_assignment#read PolicyVirtualMachineConfigurationAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/policy_virtual_machine_configuration_assignment#update PolicyVirtualMachineConfigurationAssignment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

