// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresqlFlexibleServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#location PostgresqlFlexibleServer#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#name PostgresqlFlexibleServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#resource_group_name PostgresqlFlexibleServer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#administrator_login PostgresqlFlexibleServer#administrator_login}.
	AdministratorLogin *string `field:"optional" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#administrator_password PostgresqlFlexibleServer#administrator_password}.
	AdministratorPassword *string `field:"optional" json:"administratorPassword" yaml:"administratorPassword"`
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#authentication PostgresqlFlexibleServer#authentication}
	Authentication *PostgresqlFlexibleServerAuthentication `field:"optional" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#auto_grow_enabled PostgresqlFlexibleServer#auto_grow_enabled}.
	AutoGrowEnabled interface{} `field:"optional" json:"autoGrowEnabled" yaml:"autoGrowEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#backup_retention_days PostgresqlFlexibleServer#backup_retention_days}.
	BackupRetentionDays *float64 `field:"optional" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#create_mode PostgresqlFlexibleServer#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// customer_managed_key block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#customer_managed_key PostgresqlFlexibleServer#customer_managed_key}
	CustomerManagedKey *PostgresqlFlexibleServerCustomerManagedKey `field:"optional" json:"customerManagedKey" yaml:"customerManagedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#delegated_subnet_id PostgresqlFlexibleServer#delegated_subnet_id}.
	DelegatedSubnetId *string `field:"optional" json:"delegatedSubnetId" yaml:"delegatedSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#geo_redundant_backup_enabled PostgresqlFlexibleServer#geo_redundant_backup_enabled}.
	GeoRedundantBackupEnabled interface{} `field:"optional" json:"geoRedundantBackupEnabled" yaml:"geoRedundantBackupEnabled"`
	// high_availability block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#high_availability PostgresqlFlexibleServer#high_availability}
	HighAvailability *PostgresqlFlexibleServerHighAvailability `field:"optional" json:"highAvailability" yaml:"highAvailability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#id PostgresqlFlexibleServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#identity PostgresqlFlexibleServer#identity}
	Identity *PostgresqlFlexibleServerIdentity `field:"optional" json:"identity" yaml:"identity"`
	// maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#maintenance_window PostgresqlFlexibleServer#maintenance_window}
	MaintenanceWindow *PostgresqlFlexibleServerMaintenanceWindow `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#point_in_time_restore_time_in_utc PostgresqlFlexibleServer#point_in_time_restore_time_in_utc}.
	PointInTimeRestoreTimeInUtc *string `field:"optional" json:"pointInTimeRestoreTimeInUtc" yaml:"pointInTimeRestoreTimeInUtc"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#private_dns_zone_id PostgresqlFlexibleServer#private_dns_zone_id}.
	PrivateDnsZoneId *string `field:"optional" json:"privateDnsZoneId" yaml:"privateDnsZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#public_network_access_enabled PostgresqlFlexibleServer#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#replication_role PostgresqlFlexibleServer#replication_role}.
	ReplicationRole *string `field:"optional" json:"replicationRole" yaml:"replicationRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#sku_name PostgresqlFlexibleServer#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#source_server_id PostgresqlFlexibleServer#source_server_id}.
	SourceServerId *string `field:"optional" json:"sourceServerId" yaml:"sourceServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#storage_mb PostgresqlFlexibleServer#storage_mb}.
	StorageMb *float64 `field:"optional" json:"storageMb" yaml:"storageMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#storage_tier PostgresqlFlexibleServer#storage_tier}.
	StorageTier *string `field:"optional" json:"storageTier" yaml:"storageTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#tags PostgresqlFlexibleServer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#timeouts PostgresqlFlexibleServer#timeouts}
	Timeouts *PostgresqlFlexibleServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#version PostgresqlFlexibleServer#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/postgresql_flexible_server#zone PostgresqlFlexibleServer#zone}.
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

