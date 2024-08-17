// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlserver


type MysqlServerThreatDetectionPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#disabled_alerts MysqlServer#disabled_alerts}.
	DisabledAlerts *[]*string `field:"optional" json:"disabledAlerts" yaml:"disabledAlerts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#email_account_admins MysqlServer#email_account_admins}.
	EmailAccountAdmins interface{} `field:"optional" json:"emailAccountAdmins" yaml:"emailAccountAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#email_addresses MysqlServer#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#enabled MysqlServer#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#retention_days MysqlServer#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#storage_account_access_key MysqlServer#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/mysql_server#storage_endpoint MysqlServer#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
}

