// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapsinglenodevirtualinstance


type WorkloadsSapSingleNodeVirtualInstanceSingleServerConfigurationVirtualMachineConfigurationOsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/workloads_sap_single_node_virtual_instance#admin_username WorkloadsSapSingleNodeVirtualInstance#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/workloads_sap_single_node_virtual_instance#ssh_private_key WorkloadsSapSingleNodeVirtualInstance#ssh_private_key}.
	SshPrivateKey *string `field:"required" json:"sshPrivateKey" yaml:"sshPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/workloads_sap_single_node_virtual_instance#ssh_public_key WorkloadsSapSingleNodeVirtualInstance#ssh_public_key}.
	SshPublicKey *string `field:"required" json:"sshPublicKey" yaml:"sshPublicKey"`
}

