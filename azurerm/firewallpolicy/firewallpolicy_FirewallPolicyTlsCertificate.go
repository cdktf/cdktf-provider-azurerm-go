package firewallpolicy


type FirewallPolicyTlsCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#key_vault_secret_id FirewallPolicy#key_vault_secret_id}.
	KeyVaultSecretId *string `field:"required" json:"keyVaultSecretId" yaml:"keyVaultSecretId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#name FirewallPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

