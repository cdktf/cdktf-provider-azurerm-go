// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqlvirtualmachine


type MssqlVirtualMachineWsfcDomainCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_virtual_machine#cluster_bootstrap_account_password MssqlVirtualMachine#cluster_bootstrap_account_password}.
	ClusterBootstrapAccountPassword *string `field:"required" json:"clusterBootstrapAccountPassword" yaml:"clusterBootstrapAccountPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_virtual_machine#cluster_operator_account_password MssqlVirtualMachine#cluster_operator_account_password}.
	ClusterOperatorAccountPassword *string `field:"required" json:"clusterOperatorAccountPassword" yaml:"clusterOperatorAccountPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/mssql_virtual_machine#sql_service_account_password MssqlVirtualMachine#sql_service_account_password}.
	SqlServiceAccountPassword *string `field:"required" json:"sqlServiceAccountPassword" yaml:"sqlServiceAccountPassword"`
}

