// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscache


type RedisCacheRedisConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#active_directory_authentication_enabled RedisCache#active_directory_authentication_enabled}.
	ActiveDirectoryAuthenticationEnabled interface{} `field:"optional" json:"activeDirectoryAuthenticationEnabled" yaml:"activeDirectoryAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#aof_backup_enabled RedisCache#aof_backup_enabled}.
	AofBackupEnabled interface{} `field:"optional" json:"aofBackupEnabled" yaml:"aofBackupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#aof_storage_connection_string_0 RedisCache#aof_storage_connection_string_0}.
	AofStorageConnectionString0 *string `field:"optional" json:"aofStorageConnectionString0" yaml:"aofStorageConnectionString0"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#aof_storage_connection_string_1 RedisCache#aof_storage_connection_string_1}.
	AofStorageConnectionString1 *string `field:"optional" json:"aofStorageConnectionString1" yaml:"aofStorageConnectionString1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#authentication_enabled RedisCache#authentication_enabled}.
	AuthenticationEnabled interface{} `field:"optional" json:"authenticationEnabled" yaml:"authenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#data_persistence_authentication_method RedisCache#data_persistence_authentication_method}.
	DataPersistenceAuthenticationMethod *string `field:"optional" json:"dataPersistenceAuthenticationMethod" yaml:"dataPersistenceAuthenticationMethod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#maxfragmentationmemory_reserved RedisCache#maxfragmentationmemory_reserved}.
	MaxfragmentationmemoryReserved *float64 `field:"optional" json:"maxfragmentationmemoryReserved" yaml:"maxfragmentationmemoryReserved"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#maxmemory_delta RedisCache#maxmemory_delta}.
	MaxmemoryDelta *float64 `field:"optional" json:"maxmemoryDelta" yaml:"maxmemoryDelta"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#maxmemory_policy RedisCache#maxmemory_policy}.
	MaxmemoryPolicy *string `field:"optional" json:"maxmemoryPolicy" yaml:"maxmemoryPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#maxmemory_reserved RedisCache#maxmemory_reserved}.
	MaxmemoryReserved *float64 `field:"optional" json:"maxmemoryReserved" yaml:"maxmemoryReserved"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#notify_keyspace_events RedisCache#notify_keyspace_events}.
	NotifyKeyspaceEvents *string `field:"optional" json:"notifyKeyspaceEvents" yaml:"notifyKeyspaceEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#rdb_backup_enabled RedisCache#rdb_backup_enabled}.
	RdbBackupEnabled interface{} `field:"optional" json:"rdbBackupEnabled" yaml:"rdbBackupEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#rdb_backup_frequency RedisCache#rdb_backup_frequency}.
	RdbBackupFrequency *float64 `field:"optional" json:"rdbBackupFrequency" yaml:"rdbBackupFrequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#rdb_backup_max_snapshot_count RedisCache#rdb_backup_max_snapshot_count}.
	RdbBackupMaxSnapshotCount *float64 `field:"optional" json:"rdbBackupMaxSnapshotCount" yaml:"rdbBackupMaxSnapshotCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#rdb_storage_connection_string RedisCache#rdb_storage_connection_string}.
	RdbStorageConnectionString *string `field:"optional" json:"rdbStorageConnectionString" yaml:"rdbStorageConnectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/redis_cache#storage_account_subscription_id RedisCache#storage_account_subscription_id}.
	StorageAccountSubscriptionId *string `field:"optional" json:"storageAccountSubscriptionId" yaml:"storageAccountSubscriptionId"`
}

