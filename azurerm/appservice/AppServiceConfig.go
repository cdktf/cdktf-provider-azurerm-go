package appservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#app_service_plan_id AppService#app_service_plan_id}.
	AppServicePlanId *string `field:"required" json:"appServicePlanId" yaml:"appServicePlanId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#location AppService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#name AppService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#resource_group_name AppService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#app_settings AppService#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#auth_settings AppService#auth_settings}
	AuthSettings *AppServiceAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#backup AppService#backup}
	Backup *AppServiceBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#client_affinity_enabled AppService#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#client_cert_enabled AppService#client_cert_enabled}.
	ClientCertEnabled interface{} `field:"optional" json:"clientCertEnabled" yaml:"clientCertEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#client_cert_mode AppService#client_cert_mode}.
	ClientCertMode *string `field:"optional" json:"clientCertMode" yaml:"clientCertMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#connection_string AppService#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#enabled AppService#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#https_only AppService#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#id AppService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#identity AppService#identity}
	Identity *AppServiceIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#key_vault_reference_identity_id AppService#key_vault_reference_identity_id}.
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#logs AppService#logs}
	Logs *AppServiceLogs `field:"optional" json:"logs" yaml:"logs"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#site_config AppService#site_config}
	SiteConfig *AppServiceSiteConfig `field:"optional" json:"siteConfig" yaml:"siteConfig"`
	// source_control block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#source_control AppService#source_control}
	SourceControl *AppServiceSourceControl `field:"optional" json:"sourceControl" yaml:"sourceControl"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#storage_account AppService#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#tags AppService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#timeouts AppService#timeouts}
	Timeouts *AppServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

