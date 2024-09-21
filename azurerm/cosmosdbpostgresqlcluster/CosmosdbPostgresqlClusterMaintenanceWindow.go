// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbpostgresqlcluster


type CosmosdbPostgresqlClusterMaintenanceWindow struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/cosmosdb_postgresql_cluster#day_of_week CosmosdbPostgresqlCluster#day_of_week}.
	DayOfWeek *float64 `field:"optional" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/cosmosdb_postgresql_cluster#start_hour CosmosdbPostgresqlCluster#start_hour}.
	StartHour *float64 `field:"optional" json:"startHour" yaml:"startHour"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.3.0/docs/resources/cosmosdb_postgresql_cluster#start_minute CosmosdbPostgresqlCluster#start_minute}.
	StartMinute *float64 `field:"optional" json:"startMinute" yaml:"startMinute"`
}

