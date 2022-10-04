package mssqldatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#name MssqlDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#server_id MssqlDatabase#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#auto_pause_delay_in_minutes MssqlDatabase#auto_pause_delay_in_minutes}.
	AutoPauseDelayInMinutes *float64 `field:"optional" json:"autoPauseDelayInMinutes" yaml:"autoPauseDelayInMinutes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#collation MssqlDatabase#collation}.
	Collation *string `field:"optional" json:"collation" yaml:"collation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#create_mode MssqlDatabase#create_mode}.
	CreateMode *string `field:"optional" json:"createMode" yaml:"createMode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#creation_source_database_id MssqlDatabase#creation_source_database_id}.
	CreationSourceDatabaseId *string `field:"optional" json:"creationSourceDatabaseId" yaml:"creationSourceDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#elastic_pool_id MssqlDatabase#elastic_pool_id}.
	ElasticPoolId *string `field:"optional" json:"elasticPoolId" yaml:"elasticPoolId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#geo_backup_enabled MssqlDatabase#geo_backup_enabled}.
	GeoBackupEnabled interface{} `field:"optional" json:"geoBackupEnabled" yaml:"geoBackupEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#id MssqlDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#ledger_enabled MssqlDatabase#ledger_enabled}.
	LedgerEnabled interface{} `field:"optional" json:"ledgerEnabled" yaml:"ledgerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#license_type MssqlDatabase#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// long_term_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#long_term_retention_policy MssqlDatabase#long_term_retention_policy}
	LongTermRetentionPolicy *MssqlDatabaseLongTermRetentionPolicy `field:"optional" json:"longTermRetentionPolicy" yaml:"longTermRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#maintenance_configuration_name MssqlDatabase#maintenance_configuration_name}.
	MaintenanceConfigurationName *string `field:"optional" json:"maintenanceConfigurationName" yaml:"maintenanceConfigurationName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#max_size_gb MssqlDatabase#max_size_gb}.
	MaxSizeGb *float64 `field:"optional" json:"maxSizeGb" yaml:"maxSizeGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#min_capacity MssqlDatabase#min_capacity}.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#read_replica_count MssqlDatabase#read_replica_count}.
	ReadReplicaCount *float64 `field:"optional" json:"readReplicaCount" yaml:"readReplicaCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#read_scale MssqlDatabase#read_scale}.
	ReadScale interface{} `field:"optional" json:"readScale" yaml:"readScale"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#recover_database_id MssqlDatabase#recover_database_id}.
	RecoverDatabaseId *string `field:"optional" json:"recoverDatabaseId" yaml:"recoverDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#restore_dropped_database_id MssqlDatabase#restore_dropped_database_id}.
	RestoreDroppedDatabaseId *string `field:"optional" json:"restoreDroppedDatabaseId" yaml:"restoreDroppedDatabaseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#restore_point_in_time MssqlDatabase#restore_point_in_time}.
	RestorePointInTime *string `field:"optional" json:"restorePointInTime" yaml:"restorePointInTime"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#sample_name MssqlDatabase#sample_name}.
	SampleName *string `field:"optional" json:"sampleName" yaml:"sampleName"`
	// short_term_retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#short_term_retention_policy MssqlDatabase#short_term_retention_policy}
	ShortTermRetentionPolicy *MssqlDatabaseShortTermRetentionPolicy `field:"optional" json:"shortTermRetentionPolicy" yaml:"shortTermRetentionPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#sku_name MssqlDatabase#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#storage_account_type MssqlDatabase#storage_account_type}.
	StorageAccountType *string `field:"optional" json:"storageAccountType" yaml:"storageAccountType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#tags MssqlDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// threat_detection_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#threat_detection_policy MssqlDatabase#threat_detection_policy}
	ThreatDetectionPolicy *MssqlDatabaseThreatDetectionPolicy `field:"optional" json:"threatDetectionPolicy" yaml:"threatDetectionPolicy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#timeouts MssqlDatabase#timeouts}
	Timeouts *MssqlDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#transparent_data_encryption_enabled MssqlDatabase#transparent_data_encryption_enabled}.
	TransparentDataEncryptionEnabled interface{} `field:"optional" json:"transparentDataEncryptionEnabled" yaml:"transparentDataEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_database#zone_redundant MssqlDatabase#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

