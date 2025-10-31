// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineruncommand

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineRunCommandConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#location VirtualMachineRunCommand#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#name VirtualMachineRunCommand#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#source VirtualMachineRunCommand#source}
	Source *VirtualMachineRunCommandSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#virtual_machine_id VirtualMachineRunCommand#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// error_blob_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#error_blob_managed_identity VirtualMachineRunCommand#error_blob_managed_identity}
	ErrorBlobManagedIdentity *VirtualMachineRunCommandErrorBlobManagedIdentity `field:"optional" json:"errorBlobManagedIdentity" yaml:"errorBlobManagedIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#error_blob_uri VirtualMachineRunCommand#error_blob_uri}.
	ErrorBlobUri *string `field:"optional" json:"errorBlobUri" yaml:"errorBlobUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#id VirtualMachineRunCommand#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// output_blob_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#output_blob_managed_identity VirtualMachineRunCommand#output_blob_managed_identity}
	OutputBlobManagedIdentity *VirtualMachineRunCommandOutputBlobManagedIdentity `field:"optional" json:"outputBlobManagedIdentity" yaml:"outputBlobManagedIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#output_blob_uri VirtualMachineRunCommand#output_blob_uri}.
	OutputBlobUri *string `field:"optional" json:"outputBlobUri" yaml:"outputBlobUri"`
	// parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#parameter VirtualMachineRunCommand#parameter}
	Parameter interface{} `field:"optional" json:"parameter" yaml:"parameter"`
	// protected_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#protected_parameter VirtualMachineRunCommand#protected_parameter}
	ProtectedParameter interface{} `field:"optional" json:"protectedParameter" yaml:"protectedParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#run_as_password VirtualMachineRunCommand#run_as_password}.
	RunAsPassword *string `field:"optional" json:"runAsPassword" yaml:"runAsPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#run_as_user VirtualMachineRunCommand#run_as_user}.
	RunAsUser *string `field:"optional" json:"runAsUser" yaml:"runAsUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#tags VirtualMachineRunCommand#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/virtual_machine_run_command#timeouts VirtualMachineRunCommand#timeouts}
	Timeouts *VirtualMachineRunCommandTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

