// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#key_opts KeyVaultKey#key_opts}.
	KeyOpts *[]*string `field:"required" json:"keyOpts" yaml:"keyOpts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#key_type KeyVaultKey#key_type}.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#key_vault_id KeyVaultKey#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#name KeyVaultKey#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#curve KeyVaultKey#curve}.
	Curve *string `field:"optional" json:"curve" yaml:"curve"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#expiration_date KeyVaultKey#expiration_date}.
	ExpirationDate *string `field:"optional" json:"expirationDate" yaml:"expirationDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#id KeyVaultKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#key_size KeyVaultKey#key_size}.
	KeySize *float64 `field:"optional" json:"keySize" yaml:"keySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#not_before_date KeyVaultKey#not_before_date}.
	NotBeforeDate *string `field:"optional" json:"notBeforeDate" yaml:"notBeforeDate"`
	// rotation_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#rotation_policy KeyVaultKey#rotation_policy}
	RotationPolicy *KeyVaultKeyRotationPolicy `field:"optional" json:"rotationPolicy" yaml:"rotationPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#tags KeyVaultKey#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.98.0/docs/resources/key_vault_key#timeouts KeyVaultKey#timeouts}
	Timeouts *KeyVaultKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

