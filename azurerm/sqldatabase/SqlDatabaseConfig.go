// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqldatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SqlDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#location SqlDatabase#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#name SqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#resource_group_name SqlDatabase#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#server_name SqlDatabase#server_name}.
	ServerName *string `field:"required" json:"serverName" yaml:"serverName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#collation SqlDatabase#collation}.
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#create_mode SqlDatabase#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#edition SqlDatabase#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#elastic_pool_name SqlDatabase#elastic_pool_name}.
	ElasticPoolName *string `field:"optional" json:"elasticPoolName" yaml:"elasticPoolName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#id SqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// import block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#import SqlDatabase#import}
	Import *SqlDatabaseImport `field:"optional" json:"import" yaml:"import"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#max_size_bytes SqlDatabase#max_size_bytes}.
	MaxSizeBytes *string `field:"optional" json:"maxSizeBytes" yaml:"maxSizeBytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#max_size_gb SqlDatabase#max_size_gb}.
	MaxSizeGb *string `field:"optional" json:"maxSizeGb" yaml:"maxSizeGb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#read_scale SqlDatabase#read_scale}.
	ReadScale interface{} `field:"optional" json:"readScale" yaml:"readScale"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#requested_service_objective_id SqlDatabase#requested_service_objective_id}.
	RequestedServiceObjectiveId *string `field:"optional" json:"requestedServiceObjectiveId" yaml:"requestedServiceObjectiveId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#requested_service_objective_name SqlDatabase#requested_service_objective_name}.
	RequestedServiceObjectiveName *string `field:"optional" json:"requestedServiceObjectiveName" yaml:"requestedServiceObjectiveName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#restore_point_in_time SqlDatabase#restore_point_in_time}.
	RestorePointInTime *string `field:"optional" json:"restorePointInTime" yaml:"restorePointInTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#source_database_deletion_date SqlDatabase#source_database_deletion_date}.
	SourceDatabaseDeletionDate *string `field:"optional" json:"sourceDatabaseDeletionDate" yaml:"sourceDatabaseDeletionDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#source_database_id SqlDatabase#source_database_id}.
	SourceDatabaseId *string `field:"optional" json:"sourceDatabaseId" yaml:"sourceDatabaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#tags SqlDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// threat_detection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#threat_detection_policy SqlDatabase#threat_detection_policy}
	ThreatDetectionPolicy *SqlDatabaseThreatDetectionPolicy `field:"optional" json:"threatDetectionPolicy" yaml:"threatDetectionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#timeouts SqlDatabase#timeouts}
	Timeouts *SqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/sql_database#zone_redundant SqlDatabase#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

