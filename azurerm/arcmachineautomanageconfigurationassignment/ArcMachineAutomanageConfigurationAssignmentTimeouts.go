// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arcmachineautomanageconfigurationassignment


type ArcMachineAutomanageConfigurationAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/arc_machine_automanage_configuration_assignment#create ArcMachineAutomanageConfigurationAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/arc_machine_automanage_configuration_assignment#delete ArcMachineAutomanageConfigurationAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/arc_machine_automanage_configuration_assignment#read ArcMachineAutomanageConfigurationAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

