// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfiguration struct {
	// application_server_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#application_server_configuration WorkloadsSapThreeTierVirtualInstance#application_server_configuration}
	ApplicationServerConfiguration *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationApplicationServerConfiguration `field:"required" json:"applicationServerConfiguration" yaml:"applicationServerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#app_resource_group_name WorkloadsSapThreeTierVirtualInstance#app_resource_group_name}.
	AppResourceGroupName *string `field:"required" json:"appResourceGroupName" yaml:"appResourceGroupName"`
	// central_server_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#central_server_configuration WorkloadsSapThreeTierVirtualInstance#central_server_configuration}
	CentralServerConfiguration *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationCentralServerConfiguration `field:"required" json:"centralServerConfiguration" yaml:"centralServerConfiguration"`
	// database_server_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#database_server_configuration WorkloadsSapThreeTierVirtualInstance#database_server_configuration}
	DatabaseServerConfiguration *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationDatabaseServerConfiguration `field:"required" json:"databaseServerConfiguration" yaml:"databaseServerConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#high_availability_type WorkloadsSapThreeTierVirtualInstance#high_availability_type}.
	HighAvailabilityType *string `field:"optional" json:"highAvailabilityType" yaml:"highAvailabilityType"`
	// resource_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#resource_names WorkloadsSapThreeTierVirtualInstance#resource_names}
	ResourceNames *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNames `field:"optional" json:"resourceNames" yaml:"resourceNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#secondary_ip_enabled WorkloadsSapThreeTierVirtualInstance#secondary_ip_enabled}.
	SecondaryIpEnabled interface{} `field:"optional" json:"secondaryIpEnabled" yaml:"secondaryIpEnabled"`
	// transport_create_and_mount block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/workloads_sap_three_tier_virtual_instance#transport_create_and_mount WorkloadsSapThreeTierVirtualInstance#transport_create_and_mount}
	TransportCreateAndMount *WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationTransportCreateAndMount `field:"optional" json:"transportCreateAndMount" yaml:"transportCreateAndMount"`
}

