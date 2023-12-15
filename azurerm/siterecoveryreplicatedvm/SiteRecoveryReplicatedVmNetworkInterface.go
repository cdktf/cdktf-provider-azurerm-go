// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#failover_test_public_ip_address_id SiteRecoveryReplicatedVm#failover_test_public_ip_address_id}.
	FailoverTestPublicIpAddressId *string `field:"optional" json:"failoverTestPublicIpAddressId" yaml:"failoverTestPublicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#failover_test_static_ip SiteRecoveryReplicatedVm#failover_test_static_ip}.
	FailoverTestStaticIp *string `field:"optional" json:"failoverTestStaticIp" yaml:"failoverTestStaticIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#failover_test_subnet_name SiteRecoveryReplicatedVm#failover_test_subnet_name}.
	FailoverTestSubnetName *string `field:"optional" json:"failoverTestSubnetName" yaml:"failoverTestSubnetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#is_primary SiteRecoveryReplicatedVm#is_primary}.
	IsPrimary interface{} `field:"optional" json:"isPrimary" yaml:"isPrimary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#recovery_public_ip_address_id SiteRecoveryReplicatedVm#recovery_public_ip_address_id}.
	RecoveryPublicIpAddressId *string `field:"optional" json:"recoveryPublicIpAddressId" yaml:"recoveryPublicIpAddressId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#source_network_interface_id SiteRecoveryReplicatedVm#source_network_interface_id}.
	SourceNetworkInterfaceId *string `field:"optional" json:"sourceNetworkInterfaceId" yaml:"sourceNetworkInterfaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#target_static_ip SiteRecoveryReplicatedVm#target_static_ip}.
	TargetStaticIp *string `field:"optional" json:"targetStaticIp" yaml:"targetStaticIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/site_recovery_replicated_vm#target_subnet_name SiteRecoveryReplicatedVm#target_subnet_name}.
	TargetSubnetName *string `field:"optional" json:"targetSubnetName" yaml:"targetSubnetName"`
}

