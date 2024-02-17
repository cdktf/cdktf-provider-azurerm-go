// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserver


type PostgresqlFlexibleServerCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/postgresql_flexible_server#geo_backup_key_vault_key_id PostgresqlFlexibleServer#geo_backup_key_vault_key_id}.
	GeoBackupKeyVaultKeyId *string `field:"optional" json:"geoBackupKeyVaultKeyId" yaml:"geoBackupKeyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/postgresql_flexible_server#geo_backup_user_assigned_identity_id PostgresqlFlexibleServer#geo_backup_user_assigned_identity_id}.
	GeoBackupUserAssignedIdentityId *string `field:"optional" json:"geoBackupUserAssignedIdentityId" yaml:"geoBackupUserAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/postgresql_flexible_server#key_vault_key_id PostgresqlFlexibleServer#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/postgresql_flexible_server#primary_user_assigned_identity_id PostgresqlFlexibleServer#primary_user_assigned_identity_id}.
	PrimaryUserAssignedIdentityId *string `field:"optional" json:"primaryUserAssignedIdentityId" yaml:"primaryUserAssignedIdentityId"`
}

