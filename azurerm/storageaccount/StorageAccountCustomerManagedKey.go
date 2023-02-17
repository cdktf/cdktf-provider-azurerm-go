package storageaccount


type StorageAccountCustomerManagedKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#key_vault_key_id StorageAccount#key_vault_key_id}.
	KeyVaultKeyId *string `field:"required" json:"keyVaultKeyId" yaml:"keyVaultKeyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/storage_account#user_assigned_identity_id StorageAccount#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"required" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
}

