// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlmanagedinstancetransparentdataencryption

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlManagedInstanceTransparentDataEncryptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/mssql_managed_instance_transparent_data_encryption#managed_instance_id MssqlManagedInstanceTransparentDataEncryption#managed_instance_id}.
	ManagedInstanceId *string `field:"required" json:"managedInstanceId" yaml:"managedInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/mssql_managed_instance_transparent_data_encryption#auto_rotation_enabled MssqlManagedInstanceTransparentDataEncryption#auto_rotation_enabled}.
	AutoRotationEnabled interface{} `field:"optional" json:"autoRotationEnabled" yaml:"autoRotationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/mssql_managed_instance_transparent_data_encryption#id MssqlManagedInstanceTransparentDataEncryption#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/mssql_managed_instance_transparent_data_encryption#key_vault_key_id MssqlManagedInstanceTransparentDataEncryption#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/mssql_managed_instance_transparent_data_encryption#timeouts MssqlManagedInstanceTransparentDataEncryption#timeouts}
	Timeouts *MssqlManagedInstanceTransparentDataEncryptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

