// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KeyVaultManagedHardwareSecurityModuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#admin_object_ids KeyVaultManagedHardwareSecurityModule#admin_object_ids}.
	AdminObjectIds *[]*string `field:"required" json:"adminObjectIds" yaml:"adminObjectIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#location KeyVaultManagedHardwareSecurityModule#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#name KeyVaultManagedHardwareSecurityModule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#resource_group_name KeyVaultManagedHardwareSecurityModule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#sku_name KeyVaultManagedHardwareSecurityModule#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#tenant_id KeyVaultManagedHardwareSecurityModule#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#id KeyVaultManagedHardwareSecurityModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// network_acls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#network_acls KeyVaultManagedHardwareSecurityModule#network_acls}
	NetworkAcls *KeyVaultManagedHardwareSecurityModuleNetworkAcls `field:"optional" json:"networkAcls" yaml:"networkAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#public_network_access_enabled KeyVaultManagedHardwareSecurityModule#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#purge_protection_enabled KeyVaultManagedHardwareSecurityModule#purge_protection_enabled}.
	PurgeProtectionEnabled interface{} `field:"optional" json:"purgeProtectionEnabled" yaml:"purgeProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#security_domain_key_vault_certificate_ids KeyVaultManagedHardwareSecurityModule#security_domain_key_vault_certificate_ids}.
	SecurityDomainKeyVaultCertificateIds *[]*string `field:"optional" json:"securityDomainKeyVaultCertificateIds" yaml:"securityDomainKeyVaultCertificateIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#security_domain_quorum KeyVaultManagedHardwareSecurityModule#security_domain_quorum}.
	SecurityDomainQuorum *float64 `field:"optional" json:"securityDomainQuorum" yaml:"securityDomainQuorum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#soft_delete_retention_days KeyVaultManagedHardwareSecurityModule#soft_delete_retention_days}.
	SoftDeleteRetentionDays *float64 `field:"optional" json:"softDeleteRetentionDays" yaml:"softDeleteRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#tags KeyVaultManagedHardwareSecurityModule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.72.0/docs/resources/key_vault_managed_hardware_security_module#timeouts KeyVaultManagedHardwareSecurityModule#timeouts}
	Timeouts *KeyVaultManagedHardwareSecurityModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

