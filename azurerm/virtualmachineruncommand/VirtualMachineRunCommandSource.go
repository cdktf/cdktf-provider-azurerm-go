// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineruncommand


type VirtualMachineRunCommandSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/virtual_machine_run_command#command_id VirtualMachineRunCommand#command_id}.
	CommandId *string `field:"optional" json:"commandId" yaml:"commandId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/virtual_machine_run_command#script VirtualMachineRunCommand#script}.
	Script *string `field:"optional" json:"script" yaml:"script"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/virtual_machine_run_command#script_uri VirtualMachineRunCommand#script_uri}.
	ScriptUri *string `field:"optional" json:"scriptUri" yaml:"scriptUri"`
	// script_uri_managed_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/virtual_machine_run_command#script_uri_managed_identity VirtualMachineRunCommand#script_uri_managed_identity}
	ScriptUriManagedIdentity *VirtualMachineRunCommandSourceScriptUriManagedIdentity `field:"optional" json:"scriptUriManagedIdentity" yaml:"scriptUriManagedIdentity"`
}

