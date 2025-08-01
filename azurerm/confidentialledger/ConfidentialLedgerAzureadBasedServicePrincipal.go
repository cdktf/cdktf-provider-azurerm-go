// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package confidentialledger


type ConfidentialLedgerAzureadBasedServicePrincipal struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/confidential_ledger#ledger_role_name ConfidentialLedger#ledger_role_name}.
	LedgerRoleName *string `field:"required" json:"ledgerRoleName" yaml:"ledgerRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/confidential_ledger#principal_id ConfidentialLedger#principal_id}.
	PrincipalId *string `field:"required" json:"principalId" yaml:"principalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/confidential_ledger#tenant_id ConfidentialLedger#tenant_id}.
	TenantId *string `field:"required" json:"tenantId" yaml:"tenantId"`
}

