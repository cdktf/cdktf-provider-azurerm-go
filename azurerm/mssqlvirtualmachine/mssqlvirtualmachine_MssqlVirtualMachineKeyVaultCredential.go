package mssqlvirtualmachine


type MssqlVirtualMachineKeyVaultCredential struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#key_vault_url MssqlVirtualMachine#key_vault_url}.
	KeyVaultUrl *string `field:"required" json:"keyVaultUrl" yaml:"keyVaultUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#name MssqlVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#service_principal_name MssqlVirtualMachine#service_principal_name}.
	ServicePrincipalName *string `field:"required" json:"servicePrincipalName" yaml:"servicePrincipalName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/mssql_virtual_machine#service_principal_secret MssqlVirtualMachine#service_principal_secret}.
	ServicePrincipalSecret *string `field:"required" json:"servicePrincipalSecret" yaml:"servicePrincipalSecret"`
}

