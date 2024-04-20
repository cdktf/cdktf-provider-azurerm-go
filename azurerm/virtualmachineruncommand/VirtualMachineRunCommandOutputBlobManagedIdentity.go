// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualmachineruncommand


type VirtualMachineRunCommandOutputBlobManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_run_command#client_id VirtualMachineRunCommand#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/virtual_machine_run_command#object_id VirtualMachineRunCommand#object_id}.
	ObjectId *string `field:"optional" json:"objectId" yaml:"objectId"`
}

