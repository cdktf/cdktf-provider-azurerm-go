package siterecoveryreplicatedvm

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SiteRecoveryReplicatedVmConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#name SiteRecoveryReplicatedVm#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#recovery_replication_policy_id SiteRecoveryReplicatedVm#recovery_replication_policy_id}.
	RecoveryReplicationPolicyId *string `field:"required" json:"recoveryReplicationPolicyId" yaml:"recoveryReplicationPolicyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#recovery_vault_name SiteRecoveryReplicatedVm#recovery_vault_name}.
	RecoveryVaultName *string `field:"required" json:"recoveryVaultName" yaml:"recoveryVaultName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#resource_group_name SiteRecoveryReplicatedVm#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#source_recovery_fabric_name SiteRecoveryReplicatedVm#source_recovery_fabric_name}.
	SourceRecoveryFabricName *string `field:"required" json:"sourceRecoveryFabricName" yaml:"sourceRecoveryFabricName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#source_recovery_protection_container_name SiteRecoveryReplicatedVm#source_recovery_protection_container_name}.
	SourceRecoveryProtectionContainerName *string `field:"required" json:"sourceRecoveryProtectionContainerName" yaml:"sourceRecoveryProtectionContainerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#source_vm_id SiteRecoveryReplicatedVm#source_vm_id}.
	SourceVmId *string `field:"required" json:"sourceVmId" yaml:"sourceVmId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_recovery_fabric_id SiteRecoveryReplicatedVm#target_recovery_fabric_id}.
	TargetRecoveryFabricId *string `field:"required" json:"targetRecoveryFabricId" yaml:"targetRecoveryFabricId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_recovery_protection_container_id SiteRecoveryReplicatedVm#target_recovery_protection_container_id}.
	TargetRecoveryProtectionContainerId *string `field:"required" json:"targetRecoveryProtectionContainerId" yaml:"targetRecoveryProtectionContainerId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_resource_group_id SiteRecoveryReplicatedVm#target_resource_group_id}.
	TargetResourceGroupId *string `field:"required" json:"targetResourceGroupId" yaml:"targetResourceGroupId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#id SiteRecoveryReplicatedVm#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#managed_disk SiteRecoveryReplicatedVm#managed_disk}.
	ManagedDisk interface{} `field:"optional" json:"managedDisk" yaml:"managedDisk"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#network_interface SiteRecoveryReplicatedVm#network_interface}.
	NetworkInterface interface{} `field:"optional" json:"networkInterface" yaml:"networkInterface"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_availability_set_id SiteRecoveryReplicatedVm#target_availability_set_id}.
	TargetAvailabilitySetId *string `field:"optional" json:"targetAvailabilitySetId" yaml:"targetAvailabilitySetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_network_id SiteRecoveryReplicatedVm#target_network_id}.
	TargetNetworkId *string `field:"optional" json:"targetNetworkId" yaml:"targetNetworkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#target_zone SiteRecoveryReplicatedVm#target_zone}.
	TargetZone *string `field:"optional" json:"targetZone" yaml:"targetZone"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/site_recovery_replicated_vm#timeouts SiteRecoveryReplicatedVm#timeouts}
	Timeouts *SiteRecoveryReplicatedVmTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

