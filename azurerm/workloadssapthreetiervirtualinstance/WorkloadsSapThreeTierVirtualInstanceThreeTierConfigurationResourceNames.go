// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames struct {
	// application_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/workloads_sap_three_tier_virtual_instance#application_server WorkloadsSapThreeTierVirtualInstance#application_server}
	ApplicationServer *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesApplicationServer `field:"optional" json:"applicationServer" yaml:"applicationServer"`
	// central_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/workloads_sap_three_tier_virtual_instance#central_server WorkloadsSapThreeTierVirtualInstance#central_server}
	CentralServer *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServer `field:"optional" json:"centralServer" yaml:"centralServer"`
	// database_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/workloads_sap_three_tier_virtual_instance#database_server WorkloadsSapThreeTierVirtualInstance#database_server}
	DatabaseServer *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesDatabaseServer `field:"optional" json:"databaseServer" yaml:"databaseServer"`
	// shared_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/workloads_sap_three_tier_virtual_instance#shared_storage WorkloadsSapThreeTierVirtualInstance#shared_storage}
	SharedStorage *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesSharedStorage `field:"optional" json:"sharedStorage" yaml:"sharedStorage"`
}

