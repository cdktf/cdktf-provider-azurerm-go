// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlserver

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PostgresqlServerConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#location PostgresqlServer#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#name PostgresqlServer#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#resource_group_name PostgresqlServer#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#sku_name PostgresqlServer#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#ssl_enforcement_enabled PostgresqlServer#ssl_enforcement_enabled}.
	SslEnforcementEnabled interface{} `field:"required" json:"sslEnforcementEnabled" yaml:"sslEnforcementEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#version PostgresqlServer#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#administrator_login PostgresqlServer#administrator_login}.
	AdministratorLogin *string `field:"optional" json:"administratorLogin" yaml:"administratorLogin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#administrator_login_password PostgresqlServer#administrator_login_password}.
	AdministratorLoginPassword *string `field:"optional" json:"administratorLoginPassword" yaml:"administratorLoginPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#auto_grow_enabled PostgresqlServer#auto_grow_enabled}.
	AutoGrowEnabled interface{} `field:"optional" json:"autoGrowEnabled" yaml:"autoGrowEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#backup_retention_days PostgresqlServer#backup_retention_days}.
	BackupRetentionDays *float64 `field:"optional" json:"backupRetentionDays" yaml:"backupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#create_mode PostgresqlServer#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#creation_source_server_id PostgresqlServer#creation_source_server_id}.
	CreationSourceServerId *string `field:"optional" json:"creationSourceServerId" yaml:"creationSourceServerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#geo_redundant_backup_enabled PostgresqlServer#geo_redundant_backup_enabled}.
	GeoRedundantBackupEnabled interface{} `field:"optional" json:"geoRedundantBackupEnabled" yaml:"geoRedundantBackupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#id PostgresqlServer#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#identity PostgresqlServer#identity}
	Identity *PostgresqlServerIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#infrastructure_encryption_enabled PostgresqlServer#infrastructure_encryption_enabled}.
	InfrastructureEncryptionEnabled interface{} `field:"optional" json:"infrastructureEncryptionEnabled" yaml:"infrastructureEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#public_network_access_enabled PostgresqlServer#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#restore_point_in_time PostgresqlServer#restore_point_in_time}.
	RestorePointInTime *string `field:"optional" json:"restorePointInTime" yaml:"restorePointInTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#ssl_minimal_tls_version_enforced PostgresqlServer#ssl_minimal_tls_version_enforced}.
	SslMinimalTlsVersionEnforced *string `field:"optional" json:"sslMinimalTlsVersionEnforced" yaml:"sslMinimalTlsVersionEnforced"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#storage_mb PostgresqlServer#storage_mb}.
	StorageMb *float64 `field:"optional" json:"storageMb" yaml:"storageMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#tags PostgresqlServer#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// threat_detection_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#threat_detection_policy PostgresqlServer#threat_detection_policy}
	ThreatDetectionPolicy *PostgresqlServerThreatDetectionPolicy `field:"optional" json:"threatDetectionPolicy" yaml:"threatDetectionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/postgresql_server#timeouts PostgresqlServer#timeouts}
	Timeouts *PostgresqlServerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

