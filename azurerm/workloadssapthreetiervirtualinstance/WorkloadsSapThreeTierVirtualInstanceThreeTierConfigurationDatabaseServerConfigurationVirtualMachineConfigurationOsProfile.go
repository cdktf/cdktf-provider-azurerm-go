// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfigurationVirtualMachineConfigurationOsProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/workloads_sap_three_tier_virtual_instance#admin_username WorkloadsSapThreeTierVirtualInstance#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/workloads_sap_three_tier_virtual_instance#ssh_private_key WorkloadsSapThreeTierVirtualInstance#ssh_private_key}.
	SshPrivateKey *string `field:"required" json:"sshPrivateKey" yaml:"sshPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/workloads_sap_three_tier_virtual_instance#ssh_public_key WorkloadsSapThreeTierVirtualInstance#ssh_public_key}.
	SshPublicKey *string `field:"required" json:"sshPublicKey" yaml:"sshPublicKey"`
}

