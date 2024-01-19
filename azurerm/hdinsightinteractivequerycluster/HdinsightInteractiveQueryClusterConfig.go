// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightinteractivequerycluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightInteractiveQueryClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#cluster_version HdinsightInteractiveQueryCluster#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// component_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#component_version HdinsightInteractiveQueryCluster#component_version}
	ComponentVersion *HdinsightInteractiveQueryClusterComponentVersion `field:"required" json:"componentVersion" yaml:"componentVersion"`
	// gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#gateway HdinsightInteractiveQueryCluster#gateway}
	Gateway *HdinsightInteractiveQueryClusterGateway `field:"required" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#location HdinsightInteractiveQueryCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#name HdinsightInteractiveQueryCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#resource_group_name HdinsightInteractiveQueryCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#roles HdinsightInteractiveQueryCluster#roles}
	Roles *HdinsightInteractiveQueryClusterRoles `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#tier HdinsightInteractiveQueryCluster#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// compute_isolation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#compute_isolation HdinsightInteractiveQueryCluster#compute_isolation}
	ComputeIsolation *HdinsightInteractiveQueryClusterComputeIsolation `field:"optional" json:"computeIsolation" yaml:"computeIsolation"`
	// disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#disk_encryption HdinsightInteractiveQueryCluster#disk_encryption}
	DiskEncryption interface{} `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#encryption_in_transit_enabled HdinsightInteractiveQueryCluster#encryption_in_transit_enabled}.
	EncryptionInTransitEnabled interface{} `field:"optional" json:"encryptionInTransitEnabled" yaml:"encryptionInTransitEnabled"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#extension HdinsightInteractiveQueryCluster#extension}
	Extension *HdinsightInteractiveQueryClusterExtension `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#id HdinsightInteractiveQueryCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metastores block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#metastores HdinsightInteractiveQueryCluster#metastores}
	Metastores *HdinsightInteractiveQueryClusterMetastores `field:"optional" json:"metastores" yaml:"metastores"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#monitor HdinsightInteractiveQueryCluster#monitor}
	Monitor *HdinsightInteractiveQueryClusterMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#network HdinsightInteractiveQueryCluster#network}
	Network *HdinsightInteractiveQueryClusterNetwork `field:"optional" json:"network" yaml:"network"`
	// security_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#security_profile HdinsightInteractiveQueryCluster#security_profile}
	SecurityProfile *HdinsightInteractiveQueryClusterSecurityProfile `field:"optional" json:"securityProfile" yaml:"securityProfile"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#storage_account HdinsightInteractiveQueryCluster#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// storage_account_gen2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#storage_account_gen2 HdinsightInteractiveQueryCluster#storage_account_gen2}
	StorageAccountGen2 *HdinsightInteractiveQueryClusterStorageAccountGen2 `field:"optional" json:"storageAccountGen2" yaml:"storageAccountGen2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#tags HdinsightInteractiveQueryCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#timeouts HdinsightInteractiveQueryCluster#timeouts}
	Timeouts *HdinsightInteractiveQueryClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/hdinsight_interactive_query_cluster#tls_min_version HdinsightInteractiveQueryCluster#tls_min_version}.
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
}

