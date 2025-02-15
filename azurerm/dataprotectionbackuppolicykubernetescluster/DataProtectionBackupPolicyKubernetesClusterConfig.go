// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackuppolicykubernetescluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupPolicyKubernetesClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#backup_repeating_time_intervals DataProtectionBackupPolicyKubernetesCluster#backup_repeating_time_intervals}.
	BackupRepeatingTimeIntervals *[]*string `field:"required" json:"backupRepeatingTimeIntervals" yaml:"backupRepeatingTimeIntervals"`
	// default_retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#default_retention_rule DataProtectionBackupPolicyKubernetesCluster#default_retention_rule}
	DefaultRetentionRule *DataProtectionBackupPolicyKubernetesClusterDefaultRetentionRule `field:"required" json:"defaultRetentionRule" yaml:"defaultRetentionRule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#name DataProtectionBackupPolicyKubernetesCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#resource_group_name DataProtectionBackupPolicyKubernetesCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#vault_name DataProtectionBackupPolicyKubernetesCluster#vault_name}.
	VaultName *string `field:"required" json:"vaultName" yaml:"vaultName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#id DataProtectionBackupPolicyKubernetesCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// retention_rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#retention_rule DataProtectionBackupPolicyKubernetesCluster#retention_rule}
	RetentionRule interface{} `field:"optional" json:"retentionRule" yaml:"retentionRule"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#timeouts DataProtectionBackupPolicyKubernetesCluster#timeouts}
	Timeouts *DataProtectionBackupPolicyKubernetesClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/data_protection_backup_policy_kubernetes_cluster#time_zone DataProtectionBackupPolicyKubernetesCluster#time_zone}.
	TimeZone *string `field:"optional" json:"timeZone" yaml:"timeZone"`
}

