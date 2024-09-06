// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package recoveryservicesvault


type RecoveryServicesVaultEncryption struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/recovery_services_vault#infrastructure_encryption_enabled RecoveryServicesVault#infrastructure_encryption_enabled}.
	InfrastructureEncryptionEnabled interface{} `field:"required" json:"infrastructureEncryptionEnabled" yaml:"infrastructureEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/recovery_services_vault#key_id RecoveryServicesVault#key_id}.
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/recovery_services_vault#user_assigned_identity_id RecoveryServicesVault#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/recovery_services_vault#use_system_assigned_identity RecoveryServicesVault#use_system_assigned_identity}.
	UseSystemAssignedIdentity interface{} `field:"optional" json:"useSystemAssignedIdentity" yaml:"useSystemAssignedIdentity"`
}

