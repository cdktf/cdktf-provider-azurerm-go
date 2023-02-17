package policyvirtualmachineconfigurationassignment


type PolicyVirtualMachineConfigurationAssignmentConfigurationParameter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/policy_virtual_machine_configuration_assignment#name PolicyVirtualMachineConfigurationAssignment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/policy_virtual_machine_configuration_assignment#value PolicyVirtualMachineConfigurationAssignment#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

