package mssqlvirtualmachine

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#virtual_machine_id MssqlVirtualMachine#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// assessment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#assessment MssqlVirtualMachine#assessment}
	Assessment *MssqlVirtualMachineAssessment `field:"optional" json:"assessment" yaml:"assessment"`
	// auto_backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#auto_backup MssqlVirtualMachine#auto_backup}
	AutoBackup *MssqlVirtualMachineAutoBackup `field:"optional" json:"autoBackup" yaml:"autoBackup"`
	// auto_patching block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#auto_patching MssqlVirtualMachine#auto_patching}
	AutoPatching *MssqlVirtualMachineAutoPatching `field:"optional" json:"autoPatching" yaml:"autoPatching"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#id MssqlVirtualMachine#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// key_vault_credential block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#key_vault_credential MssqlVirtualMachine#key_vault_credential}
	KeyVaultCredential *MssqlVirtualMachineKeyVaultCredential `field:"optional" json:"keyVaultCredential" yaml:"keyVaultCredential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#r_services_enabled MssqlVirtualMachine#r_services_enabled}.
	RServicesEnabled interface{} `field:"optional" json:"rServicesEnabled" yaml:"rServicesEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#sql_connectivity_port MssqlVirtualMachine#sql_connectivity_port}.
	SqlConnectivityPort *float64 `field:"optional" json:"sqlConnectivityPort" yaml:"sqlConnectivityPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#sql_connectivity_type MssqlVirtualMachine#sql_connectivity_type}.
	SqlConnectivityType *string `field:"optional" json:"sqlConnectivityType" yaml:"sqlConnectivityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#sql_connectivity_update_password MssqlVirtualMachine#sql_connectivity_update_password}.
	SqlConnectivityUpdatePassword *string `field:"optional" json:"sqlConnectivityUpdatePassword" yaml:"sqlConnectivityUpdatePassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#sql_connectivity_update_username MssqlVirtualMachine#sql_connectivity_update_username}.
	SqlConnectivityUpdateUsername *string `field:"optional" json:"sqlConnectivityUpdateUsername" yaml:"sqlConnectivityUpdateUsername"`
	// sql_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#sql_instance MssqlVirtualMachine#sql_instance}
	SqlInstance *MssqlVirtualMachineSqlInstance `field:"optional" json:"sqlInstance" yaml:"sqlInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#sql_license_type MssqlVirtualMachine#sql_license_type}.
	SqlLicenseType *string `field:"optional" json:"sqlLicenseType" yaml:"sqlLicenseType"`
	// storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#storage_configuration MssqlVirtualMachine#storage_configuration}
	StorageConfiguration *MssqlVirtualMachineStorageConfiguration `field:"optional" json:"storageConfiguration" yaml:"storageConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#tags MssqlVirtualMachine#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/mssql_virtual_machine#timeouts MssqlVirtualMachine#timeouts}
	Timeouts *MssqlVirtualMachineTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

