package siterecoveryreplicatedvm


type SiteRecoveryReplicatedVmNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#recovery_public_ip_address_id SiteRecoveryReplicatedVm#recovery_public_ip_address_id}.
	RecoveryPublicIpAddressId *string `field:"optional" json:"recoveryPublicIpAddressId" yaml:"recoveryPublicIpAddressId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#source_network_interface_id SiteRecoveryReplicatedVm#source_network_interface_id}.
	SourceNetworkInterfaceId *string `field:"optional" json:"sourceNetworkInterfaceId" yaml:"sourceNetworkInterfaceId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_static_ip SiteRecoveryReplicatedVm#target_static_ip}.
	TargetStaticIp *string `field:"optional" json:"targetStaticIp" yaml:"targetStaticIp"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_subnet_name SiteRecoveryReplicatedVm#target_subnet_name}.
	TargetSubnetName *string `field:"optional" json:"targetSubnetName" yaml:"targetSubnetName"`
}

