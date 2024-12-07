// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvault

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#location KeyVault#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#name KeyVault#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#resource_group_name KeyVault#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#sku_name KeyVault#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#tenant_id KeyVault#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#access_policy KeyVault#access_policy}.
	AccessPolicy interface{} `field:"optional" json:"accessPolicy" yaml:"accessPolicy"`
	// contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#contact KeyVault#contact}
	Contact interface{} `field:"optional" json:"contact" yaml:"contact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#enabled_for_deployment KeyVault#enabled_for_deployment}.
	EnabledForDeployment interface{} `field:"optional" json:"enabledForDeployment" yaml:"enabledForDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#enabled_for_disk_encryption KeyVault#enabled_for_disk_encryption}.
	EnabledForDiskEncryption interface{} `field:"optional" json:"enabledForDiskEncryption" yaml:"enabledForDiskEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#enabled_for_template_deployment KeyVault#enabled_for_template_deployment}.
	EnabledForTemplateDeployment interface{} `field:"optional" json:"enabledForTemplateDeployment" yaml:"enabledForTemplateDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#enable_rbac_authorization KeyVault#enable_rbac_authorization}.
	EnableRbacAuthorization interface{} `field:"optional" json:"enableRbacAuthorization" yaml:"enableRbacAuthorization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#id KeyVault#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// network_acls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#network_acls KeyVault#network_acls}
	NetworkAcls *KeyVaultNetworkAcls `field:"optional" json:"networkAcls" yaml:"networkAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#public_network_access_enabled KeyVault#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#purge_protection_enabled KeyVault#purge_protection_enabled}.
	PurgeProtectionEnabled interface{} `field:"optional" json:"purgeProtectionEnabled" yaml:"purgeProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#soft_delete_retention_days KeyVault#soft_delete_retention_days}.
	SoftDeleteRetentionDays *float64 `field:"optional" json:"softDeleteRetentionDays" yaml:"softDeleteRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#tags KeyVault#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.13.0/docs/resources/key_vault#timeouts KeyVault#timeouts}
	Timeouts *KeyVaultTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

