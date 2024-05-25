// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mysqlserverkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MysqlServerKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/mysql_server_key#key_vault_key_id MysqlServerKey#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/mysql_server_key#server_id MysqlServerKey#server_id}.
	ServerId *string `field:"required" json:"serverId" yaml:"serverId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/mysql_server_key#id MysqlServerKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/mysql_server_key#timeouts MysqlServerKey#timeouts}
	Timeouts *MysqlServerKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

