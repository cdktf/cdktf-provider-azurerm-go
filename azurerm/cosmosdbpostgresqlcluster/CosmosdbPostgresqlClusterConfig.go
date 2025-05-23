// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbpostgresqlcluster

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbPostgresqlClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#location CosmosdbPostgresqlCluster#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#name CosmosdbPostgresqlCluster#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#node_count CosmosdbPostgresqlCluster#node_count}.
	NodeCount *float64 `field:"required" json:"nodeCount" yaml:"nodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#resource_group_name CosmosdbPostgresqlCluster#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#administrator_login_password CosmosdbPostgresqlCluster#administrator_login_password}.
	AdministratorLoginPassword *string `field:"optional" json:"administratorLoginPassword" yaml:"administratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#citus_version CosmosdbPostgresqlCluster#citus_version}.
	CitusVersion *string `field:"optional" json:"citusVersion" yaml:"citusVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#coordinator_public_ip_access_enabled CosmosdbPostgresqlCluster#coordinator_public_ip_access_enabled}.
	CoordinatorPublicIpAccessEnabled interface{} `field:"optional" json:"coordinatorPublicIpAccessEnabled" yaml:"coordinatorPublicIpAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#coordinator_server_edition CosmosdbPostgresqlCluster#coordinator_server_edition}.
	CoordinatorServerEdition *string `field:"optional" json:"coordinatorServerEdition" yaml:"coordinatorServerEdition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#coordinator_storage_quota_in_mb CosmosdbPostgresqlCluster#coordinator_storage_quota_in_mb}.
	CoordinatorStorageQuotaInMb *float64 `field:"optional" json:"coordinatorStorageQuotaInMb" yaml:"coordinatorStorageQuotaInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#coordinator_vcore_count CosmosdbPostgresqlCluster#coordinator_vcore_count}.
	CoordinatorVcoreCount *float64 `field:"optional" json:"coordinatorVcoreCount" yaml:"coordinatorVcoreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#ha_enabled CosmosdbPostgresqlCluster#ha_enabled}.
	HaEnabled interface{} `field:"optional" json:"haEnabled" yaml:"haEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#id CosmosdbPostgresqlCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#maintenance_window CosmosdbPostgresqlCluster#maintenance_window}
	MaintenanceWindow *CosmosdbPostgresqlClusterMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#node_public_ip_access_enabled CosmosdbPostgresqlCluster#node_public_ip_access_enabled}.
	NodePublicIpAccessEnabled interface{} `field:"optional" json:"nodePublicIpAccessEnabled" yaml:"nodePublicIpAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#node_server_edition CosmosdbPostgresqlCluster#node_server_edition}.
	NodeServerEdition *string `field:"optional" json:"nodeServerEdition" yaml:"nodeServerEdition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#node_storage_quota_in_mb CosmosdbPostgresqlCluster#node_storage_quota_in_mb}.
	NodeStorageQuotaInMb *float64 `field:"optional" json:"nodeStorageQuotaInMb" yaml:"nodeStorageQuotaInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#node_vcores CosmosdbPostgresqlCluster#node_vcores}.
	NodeVcores *float64 `field:"optional" json:"nodeVcores" yaml:"nodeVcores"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#point_in_time_in_utc CosmosdbPostgresqlCluster#point_in_time_in_utc}.
	PointInTimeInUtc *string `field:"optional" json:"pointInTimeInUtc" yaml:"pointInTimeInUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#preferred_primary_zone CosmosdbPostgresqlCluster#preferred_primary_zone}.
	PreferredPrimaryZone *string `field:"optional" json:"preferredPrimaryZone" yaml:"preferredPrimaryZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#shards_on_coordinator_enabled CosmosdbPostgresqlCluster#shards_on_coordinator_enabled}.
	ShardsOnCoordinatorEnabled interface{} `field:"optional" json:"shardsOnCoordinatorEnabled" yaml:"shardsOnCoordinatorEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#source_location CosmosdbPostgresqlCluster#source_location}.
	SourceLocation *string `field:"optional" json:"sourceLocation" yaml:"sourceLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#source_resource_id CosmosdbPostgresqlCluster#source_resource_id}.
	SourceResourceId *string `field:"optional" json:"sourceResourceId" yaml:"sourceResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#sql_version CosmosdbPostgresqlCluster#sql_version}.
	SqlVersion *string `field:"optional" json:"sqlVersion" yaml:"sqlVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#tags CosmosdbPostgresqlCluster#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/cosmosdb_postgresql_cluster#timeouts CosmosdbPostgresqlCluster#timeouts}
	Timeouts *CosmosdbPostgresqlClusterTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

