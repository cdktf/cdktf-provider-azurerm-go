package postgresqlflexibleserver


type PostgresqlFlexibleServerCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#key_vault_key_id PostgresqlFlexibleServer#key_vault_key_id}.
	KeyVaultKeyId *string `field:"optional" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/postgresql_flexible_server#primary_user_assigned_identity_id PostgresqlFlexibleServer#primary_user_assigned_identity_id}.
	PrimaryUserAssignedIdentityId *string `field:"optional" json:"primaryUserAssignedIdentityId" yaml:"primaryUserAssignedIdentityId"`
}
