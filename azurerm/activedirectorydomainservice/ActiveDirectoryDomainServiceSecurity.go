// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package activedirectorydomainservice


type ActiveDirectoryDomainServiceSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#kerberos_armoring_enabled ActiveDirectoryDomainService#kerberos_armoring_enabled}.
	KerberosArmoringEnabled interface{} `field:"optional" json:"kerberosArmoringEnabled" yaml:"kerberosArmoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#kerberos_rc4_encryption_enabled ActiveDirectoryDomainService#kerberos_rc4_encryption_enabled}.
	KerberosRc4EncryptionEnabled interface{} `field:"optional" json:"kerberosRc4EncryptionEnabled" yaml:"kerberosRc4EncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#ntlm_v1_enabled ActiveDirectoryDomainService#ntlm_v1_enabled}.
	NtlmV1Enabled interface{} `field:"optional" json:"ntlmV1Enabled" yaml:"ntlmV1Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#sync_kerberos_passwords ActiveDirectoryDomainService#sync_kerberos_passwords}.
	SyncKerberosPasswords interface{} `field:"optional" json:"syncKerberosPasswords" yaml:"syncKerberosPasswords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#sync_ntlm_passwords ActiveDirectoryDomainService#sync_ntlm_passwords}.
	SyncNtlmPasswords interface{} `field:"optional" json:"syncNtlmPasswords" yaml:"syncNtlmPasswords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#sync_on_prem_passwords ActiveDirectoryDomainService#sync_on_prem_passwords}.
	SyncOnPremPasswords interface{} `field:"optional" json:"syncOnPremPasswords" yaml:"syncOnPremPasswords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/active_directory_domain_service#tls_v1_enabled ActiveDirectoryDomainService#tls_v1_enabled}.
	TlsV1Enabled interface{} `field:"optional" json:"tlsV1Enabled" yaml:"tlsV1Enabled"`
}

