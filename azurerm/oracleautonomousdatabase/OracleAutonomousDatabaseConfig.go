// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package oracleautonomousdatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OracleAutonomousDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#admin_password OracleAutonomousDatabase#admin_password}.
	AdminPassword *string `field:"required" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#auto_scaling_enabled OracleAutonomousDatabase#auto_scaling_enabled}.
	AutoScalingEnabled interface{} `field:"required" json:"autoScalingEnabled" yaml:"autoScalingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#auto_scaling_for_storage_enabled OracleAutonomousDatabase#auto_scaling_for_storage_enabled}.
	AutoScalingForStorageEnabled interface{} `field:"required" json:"autoScalingForStorageEnabled" yaml:"autoScalingForStorageEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#backup_retention_period_in_days OracleAutonomousDatabase#backup_retention_period_in_days}.
	BackupRetentionPeriodInDays *float64 `field:"required" json:"backupRetentionPeriodInDays" yaml:"backupRetentionPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#character_set OracleAutonomousDatabase#character_set}.
	CharacterSet *string `field:"required" json:"characterSet" yaml:"characterSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#compute_count OracleAutonomousDatabase#compute_count}.
	ComputeCount *float64 `field:"required" json:"computeCount" yaml:"computeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#compute_model OracleAutonomousDatabase#compute_model}.
	ComputeModel *string `field:"required" json:"computeModel" yaml:"computeModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#data_storage_size_in_tbs OracleAutonomousDatabase#data_storage_size_in_tbs}.
	DataStorageSizeInTbs *float64 `field:"required" json:"dataStorageSizeInTbs" yaml:"dataStorageSizeInTbs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#db_version OracleAutonomousDatabase#db_version}.
	DbVersion *string `field:"required" json:"dbVersion" yaml:"dbVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#db_workload OracleAutonomousDatabase#db_workload}.
	DbWorkload *string `field:"required" json:"dbWorkload" yaml:"dbWorkload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#display_name OracleAutonomousDatabase#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#license_model OracleAutonomousDatabase#license_model}.
	LicenseModel *string `field:"required" json:"licenseModel" yaml:"licenseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#location OracleAutonomousDatabase#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#mtls_connection_required OracleAutonomousDatabase#mtls_connection_required}.
	MtlsConnectionRequired interface{} `field:"required" json:"mtlsConnectionRequired" yaml:"mtlsConnectionRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#name OracleAutonomousDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#national_character_set OracleAutonomousDatabase#national_character_set}.
	NationalCharacterSet *string `field:"required" json:"nationalCharacterSet" yaml:"nationalCharacterSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#resource_group_name OracleAutonomousDatabase#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#subnet_id OracleAutonomousDatabase#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#virtual_network_id OracleAutonomousDatabase#virtual_network_id}.
	VirtualNetworkId *string `field:"required" json:"virtualNetworkId" yaml:"virtualNetworkId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#customer_contacts OracleAutonomousDatabase#customer_contacts}.
	CustomerContacts *[]*string `field:"optional" json:"customerContacts" yaml:"customerContacts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#id OracleAutonomousDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#tags OracleAutonomousDatabase#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/oracle_autonomous_database#timeouts OracleAutonomousDatabase#timeouts}
	Timeouts *OracleAutonomousDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

