// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineautomanageconfigurationassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineAutomanageConfigurationAssignmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_automanage_configuration_assignment#configuration_id VirtualMachineAutomanageConfigurationAssignment#configuration_id}.
	ConfigurationId *string `field:"required" json:"configurationId" yaml:"configurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_automanage_configuration_assignment#virtual_machine_id VirtualMachineAutomanageConfigurationAssignment#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_automanage_configuration_assignment#id VirtualMachineAutomanageConfigurationAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/virtual_machine_automanage_configuration_assignment#timeouts VirtualMachineAutomanageConfigurationAssignment#timeouts}
	Timeouts *VirtualMachineAutomanageConfigurationAssignmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

