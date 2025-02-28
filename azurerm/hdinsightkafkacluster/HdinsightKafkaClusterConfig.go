// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hdinsightkafkacluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightKafkaClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#cluster_version HdinsightKafkaCluster#cluster_version}.
	ClusterVersion *string `field:"required" json:"clusterVersion" yaml:"clusterVersion"`
	// component_version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#component_version HdinsightKafkaCluster#component_version}
	ComponentVersion *HdinsightKafkaClusterComponentVersion `field:"required" json:"componentVersion" yaml:"componentVersion"`
	// gateway block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#gateway HdinsightKafkaCluster#gateway}
	Gateway *HdinsightKafkaClusterGateway `field:"required" json:"gateway" yaml:"gateway"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#location HdinsightKafkaCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#name HdinsightKafkaCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#resource_group_name HdinsightKafkaCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// roles block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#roles HdinsightKafkaCluster#roles}
	Roles *HdinsightKafkaClusterRoles `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#tier HdinsightKafkaCluster#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// compute_isolation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#compute_isolation HdinsightKafkaCluster#compute_isolation}
	ComputeIsolation *HdinsightKafkaClusterComputeIsolation `field:"optional" json:"computeIsolation" yaml:"computeIsolation"`
	// disk_encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#disk_encryption HdinsightKafkaCluster#disk_encryption}
	DiskEncryption interface{} `field:"optional" json:"diskEncryption" yaml:"diskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#encryption_in_transit_enabled HdinsightKafkaCluster#encryption_in_transit_enabled}.
	EncryptionInTransitEnabled interface{} `field:"optional" json:"encryptionInTransitEnabled" yaml:"encryptionInTransitEnabled"`
	// extension block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#extension HdinsightKafkaCluster#extension}
	Extension *HdinsightKafkaClusterExtension `field:"optional" json:"extension" yaml:"extension"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#id HdinsightKafkaCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// metastores block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#metastores HdinsightKafkaCluster#metastores}
	Metastores *HdinsightKafkaClusterMetastores `field:"optional" json:"metastores" yaml:"metastores"`
	// monitor block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#monitor HdinsightKafkaCluster#monitor}
	Monitor *HdinsightKafkaClusterMonitor `field:"optional" json:"monitor" yaml:"monitor"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#network HdinsightKafkaCluster#network}
	Network *HdinsightKafkaClusterNetwork `field:"optional" json:"network" yaml:"network"`
	// private_link_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#private_link_configuration HdinsightKafkaCluster#private_link_configuration}
	PrivateLinkConfiguration *HdinsightKafkaClusterPrivateLinkConfiguration `field:"optional" json:"privateLinkConfiguration" yaml:"privateLinkConfiguration"`
	// rest_proxy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#rest_proxy HdinsightKafkaCluster#rest_proxy}
	RestProxy *HdinsightKafkaClusterRestProxy `field:"optional" json:"restProxy" yaml:"restProxy"`
	// security_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#security_profile HdinsightKafkaCluster#security_profile}
	SecurityProfile *HdinsightKafkaClusterSecurityProfile `field:"optional" json:"securityProfile" yaml:"securityProfile"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#storage_account HdinsightKafkaCluster#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// storage_account_gen2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#storage_account_gen2 HdinsightKafkaCluster#storage_account_gen2}
	StorageAccountGen2 *HdinsightKafkaClusterStorageAccountGen2 `field:"optional" json:"storageAccountGen2" yaml:"storageAccountGen2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#tags HdinsightKafkaCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#timeouts HdinsightKafkaCluster#timeouts}
	Timeouts *HdinsightKafkaClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/hdinsight_kafka_cluster#tls_min_version HdinsightKafkaCluster#tls_min_version}.
	TlsMinVersion *string `field:"optional" json:"tlsMinVersion" yaml:"tlsMinVersion"`
}

