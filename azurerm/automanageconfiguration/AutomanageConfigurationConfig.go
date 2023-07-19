package automanageconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomanageConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#location AutomanageConfiguration#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#name AutomanageConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#resource_group_name AutomanageConfiguration#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// antimalware block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#antimalware AutomanageConfiguration#antimalware}
	Antimalware *AutomanageConfigurationAntimalware `field:"optional" json:"antimalware" yaml:"antimalware"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#automation_account_enabled AutomanageConfiguration#automation_account_enabled}.
	AutomationAccountEnabled interface{} `field:"optional" json:"automationAccountEnabled" yaml:"automationAccountEnabled"`
	// azure_security_baseline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#azure_security_baseline AutomanageConfiguration#azure_security_baseline}
	AzureSecurityBaseline *AutomanageConfigurationAzureSecurityBaseline `field:"optional" json:"azureSecurityBaseline" yaml:"azureSecurityBaseline"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#backup AutomanageConfiguration#backup}
	Backup *AutomanageConfigurationBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#boot_diagnostics_enabled AutomanageConfiguration#boot_diagnostics_enabled}.
	BootDiagnosticsEnabled interface{} `field:"optional" json:"bootDiagnosticsEnabled" yaml:"bootDiagnosticsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#defender_for_cloud_enabled AutomanageConfiguration#defender_for_cloud_enabled}.
	DefenderForCloudEnabled interface{} `field:"optional" json:"defenderForCloudEnabled" yaml:"defenderForCloudEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#guest_configuration_enabled AutomanageConfiguration#guest_configuration_enabled}.
	GuestConfigurationEnabled interface{} `field:"optional" json:"guestConfigurationEnabled" yaml:"guestConfigurationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#id AutomanageConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#log_analytics_enabled AutomanageConfiguration#log_analytics_enabled}.
	LogAnalyticsEnabled interface{} `field:"optional" json:"logAnalyticsEnabled" yaml:"logAnalyticsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#status_change_alert_enabled AutomanageConfiguration#status_change_alert_enabled}.
	StatusChangeAlertEnabled interface{} `field:"optional" json:"statusChangeAlertEnabled" yaml:"statusChangeAlertEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#tags AutomanageConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/automanage_configuration#timeouts AutomanageConfiguration#timeouts}
	Timeouts *AutomanageConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

