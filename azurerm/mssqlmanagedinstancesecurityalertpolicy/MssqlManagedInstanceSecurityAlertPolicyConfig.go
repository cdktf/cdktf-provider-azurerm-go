package mssqlmanagedinstancesecurityalertpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlManagedInstanceSecurityAlertPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#managed_instance_name MssqlManagedInstanceSecurityAlertPolicy#managed_instance_name}.
	ManagedInstanceName *string `field:"required" json:"managedInstanceName" yaml:"managedInstanceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#resource_group_name MssqlManagedInstanceSecurityAlertPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#disabled_alerts MssqlManagedInstanceSecurityAlertPolicy#disabled_alerts}.
	DisabledAlerts *[]*string `field:"optional" json:"disabledAlerts" yaml:"disabledAlerts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#email_account_admins_enabled MssqlManagedInstanceSecurityAlertPolicy#email_account_admins_enabled}.
	EmailAccountAdminsEnabled interface{} `field:"optional" json:"emailAccountAdminsEnabled" yaml:"emailAccountAdminsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#email_addresses MssqlManagedInstanceSecurityAlertPolicy#email_addresses}.
	EmailAddresses *[]*string `field:"optional" json:"emailAddresses" yaml:"emailAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#enabled MssqlManagedInstanceSecurityAlertPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#id MssqlManagedInstanceSecurityAlertPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#retention_days MssqlManagedInstanceSecurityAlertPolicy#retention_days}.
	RetentionDays *float64 `field:"optional" json:"retentionDays" yaml:"retentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#storage_account_access_key MssqlManagedInstanceSecurityAlertPolicy#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#storage_endpoint MssqlManagedInstanceSecurityAlertPolicy#storage_endpoint}.
	StorageEndpoint *string `field:"optional" json:"storageEndpoint" yaml:"storageEndpoint"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/mssql_managed_instance_security_alert_policy#timeouts MssqlManagedInstanceSecurityAlertPolicy#timeouts}
	Timeouts *MssqlManagedInstanceSecurityAlertPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

