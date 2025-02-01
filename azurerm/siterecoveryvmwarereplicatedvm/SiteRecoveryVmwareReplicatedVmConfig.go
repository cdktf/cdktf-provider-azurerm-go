// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package siterecoveryvmwarereplicatedvm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryVmwareReplicatedVmConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#appliance_name SiteRecoveryVmwareReplicatedVm#appliance_name}.
	ApplianceName *string `field:"required" json:"applianceName" yaml:"applianceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#name SiteRecoveryVmwareReplicatedVm#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#physical_server_credential_name SiteRecoveryVmwareReplicatedVm#physical_server_credential_name}.
	PhysicalServerCredentialName *string `field:"required" json:"physicalServerCredentialName" yaml:"physicalServerCredentialName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#recovery_replication_policy_id SiteRecoveryVmwareReplicatedVm#recovery_replication_policy_id}.
	RecoveryReplicationPolicyId *string `field:"required" json:"recoveryReplicationPolicyId" yaml:"recoveryReplicationPolicyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#recovery_vault_id SiteRecoveryVmwareReplicatedVm#recovery_vault_id}.
	RecoveryVaultId *string `field:"required" json:"recoveryVaultId" yaml:"recoveryVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#source_vm_name SiteRecoveryVmwareReplicatedVm#source_vm_name}.
	SourceVmName *string `field:"required" json:"sourceVmName" yaml:"sourceVmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_resource_group_id SiteRecoveryVmwareReplicatedVm#target_resource_group_id}.
	TargetResourceGroupId *string `field:"required" json:"targetResourceGroupId" yaml:"targetResourceGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_vm_name SiteRecoveryVmwareReplicatedVm#target_vm_name}.
	TargetVmName *string `field:"required" json:"targetVmName" yaml:"targetVmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#default_log_storage_account_id SiteRecoveryVmwareReplicatedVm#default_log_storage_account_id}.
	DefaultLogStorageAccountId *string `field:"optional" json:"defaultLogStorageAccountId" yaml:"defaultLogStorageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#default_recovery_disk_type SiteRecoveryVmwareReplicatedVm#default_recovery_disk_type}.
	DefaultRecoveryDiskType *string `field:"optional" json:"defaultRecoveryDiskType" yaml:"defaultRecoveryDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#default_target_disk_encryption_set_id SiteRecoveryVmwareReplicatedVm#default_target_disk_encryption_set_id}.
	DefaultTargetDiskEncryptionSetId *string `field:"optional" json:"defaultTargetDiskEncryptionSetId" yaml:"defaultTargetDiskEncryptionSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#id SiteRecoveryVmwareReplicatedVm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#license_type SiteRecoveryVmwareReplicatedVm#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// managed_disk block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#managed_disk SiteRecoveryVmwareReplicatedVm#managed_disk}
	ManagedDisk interface{} `field:"optional" json:"managedDisk" yaml:"managedDisk"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#multi_vm_group_name SiteRecoveryVmwareReplicatedVm#multi_vm_group_name}.
	MultiVmGroupName *string `field:"optional" json:"multiVmGroupName" yaml:"multiVmGroupName"`
	// network_interface block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#network_interface SiteRecoveryVmwareReplicatedVm#network_interface}
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_availability_set_id SiteRecoveryVmwareReplicatedVm#target_availability_set_id}.
	TargetAvailabilitySetId *string `field:"optional" json:"targetAvailabilitySetId" yaml:"targetAvailabilitySetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_boot_diagnostics_storage_account_id SiteRecoveryVmwareReplicatedVm#target_boot_diagnostics_storage_account_id}.
	TargetBootDiagnosticsStorageAccountId *string `field:"optional" json:"targetBootDiagnosticsStorageAccountId" yaml:"targetBootDiagnosticsStorageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_network_id SiteRecoveryVmwareReplicatedVm#target_network_id}.
	TargetNetworkId *string `field:"optional" json:"targetNetworkId" yaml:"targetNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_proximity_placement_group_id SiteRecoveryVmwareReplicatedVm#target_proximity_placement_group_id}.
	TargetProximityPlacementGroupId *string `field:"optional" json:"targetProximityPlacementGroupId" yaml:"targetProximityPlacementGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_vm_size SiteRecoveryVmwareReplicatedVm#target_vm_size}.
	TargetVmSize *string `field:"optional" json:"targetVmSize" yaml:"targetVmSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#target_zone SiteRecoveryVmwareReplicatedVm#target_zone}.
	TargetZone *string `field:"optional" json:"targetZone" yaml:"targetZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#test_network_id SiteRecoveryVmwareReplicatedVm#test_network_id}.
	TestNetworkId *string `field:"optional" json:"testNetworkId" yaml:"testNetworkId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/site_recovery_vmware_replicated_vm#timeouts SiteRecoveryVmwareReplicatedVm#timeouts}
	Timeouts *SiteRecoveryVmwareReplicatedVmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

