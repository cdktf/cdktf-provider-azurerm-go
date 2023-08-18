package appconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppConfigurationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#location AppConfiguration#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#name AppConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#resource_group_name AppConfiguration#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// encryption block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#encryption AppConfiguration#encryption}
	Encryption *AppConfigurationEncryption `field:"optional" json:"encryption" yaml:"encryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#id AppConfiguration#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#identity AppConfiguration#identity}
	Identity *AppConfigurationIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#local_auth_enabled AppConfiguration#local_auth_enabled}.
	LocalAuthEnabled interface{} `field:"optional" json:"localAuthEnabled" yaml:"localAuthEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#public_network_access AppConfiguration#public_network_access}.
	PublicNetworkAccess *string `field:"optional" json:"publicNetworkAccess" yaml:"publicNetworkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#purge_protection_enabled AppConfiguration#purge_protection_enabled}.
	PurgeProtectionEnabled interface{} `field:"optional" json:"purgeProtectionEnabled" yaml:"purgeProtectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#sku AppConfiguration#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#soft_delete_retention_days AppConfiguration#soft_delete_retention_days}.
	SoftDeleteRetentionDays *float64 `field:"optional" json:"softDeleteRetentionDays" yaml:"softDeleteRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#tags AppConfiguration#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_configuration#timeouts AppConfiguration#timeouts}
	Timeouts *AppConfigurationTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

