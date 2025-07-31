// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workloadssapthreetiervirtualinstance


type WorkloadsSapThreeTierVirtualInstanceThreeTierConfigurationResourceNamesCentralServerLoadBalancer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/workloads_sap_three_tier_virtual_instance#backend_pool_names WorkloadsSapThreeTierVirtualInstance#backend_pool_names}.
	BackendPoolNames *[]*string `field:"optional" json:"backendPoolNames" yaml:"backendPoolNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/workloads_sap_three_tier_virtual_instance#frontend_ip_configuration_names WorkloadsSapThreeTierVirtualInstance#frontend_ip_configuration_names}.
	FrontendIpConfigurationNames *[]*string `field:"optional" json:"frontendIpConfigurationNames" yaml:"frontendIpConfigurationNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/workloads_sap_three_tier_virtual_instance#health_probe_names WorkloadsSapThreeTierVirtualInstance#health_probe_names}.
	HealthProbeNames *[]*string `field:"optional" json:"healthProbeNames" yaml:"healthProbeNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/workloads_sap_three_tier_virtual_instance#name WorkloadsSapThreeTierVirtualInstance#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

