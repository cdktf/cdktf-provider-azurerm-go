// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#location WindowsWebApp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#name WindowsWebApp#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#resource_group_name WindowsWebApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#service_plan_id WindowsWebApp#service_plan_id}.
	ServicePlanId *string `field:"required" json:"servicePlanId" yaml:"servicePlanId"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#site_config WindowsWebApp#site_config}
	SiteConfig *WindowsWebAppSiteConfig `field:"required" json:"siteConfig" yaml:"siteConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#app_settings WindowsWebApp#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#auth_settings WindowsWebApp#auth_settings}
	AuthSettings *WindowsWebAppAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// auth_settings_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#auth_settings_v2 WindowsWebApp#auth_settings_v2}
	AuthSettingsV2 *WindowsWebAppAuthSettingsV2 `field:"optional" json:"authSettingsV2" yaml:"authSettingsV2"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#backup WindowsWebApp#backup}
	Backup *WindowsWebAppBackup `field:"optional" json:"backup" yaml:"backup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#client_affinity_enabled WindowsWebApp#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#client_certificate_enabled WindowsWebApp#client_certificate_enabled}.
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Paths to exclude when using client certificates, separated by ;
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#client_certificate_exclusion_paths WindowsWebApp#client_certificate_exclusion_paths}
	ClientCertificateExclusionPaths *string `field:"optional" json:"clientCertificateExclusionPaths" yaml:"clientCertificateExclusionPaths"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#client_certificate_mode WindowsWebApp#client_certificate_mode}.
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#connection_string WindowsWebApp#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#enabled WindowsWebApp#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#ftp_publish_basic_authentication_enabled WindowsWebApp#ftp_publish_basic_authentication_enabled}.
	FtpPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"ftpPublishBasicAuthenticationEnabled" yaml:"ftpPublishBasicAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#https_only WindowsWebApp#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#id WindowsWebApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#identity WindowsWebApp#identity}
	Identity *WindowsWebAppIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#key_vault_reference_identity_id WindowsWebApp#key_vault_reference_identity_id}.
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#logs WindowsWebApp#logs}
	Logs *WindowsWebAppLogs `field:"optional" json:"logs" yaml:"logs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#public_network_access_enabled WindowsWebApp#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// sticky_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#sticky_settings WindowsWebApp#sticky_settings}
	StickySettings *WindowsWebAppStickySettings `field:"optional" json:"stickySettings" yaml:"stickySettings"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#storage_account WindowsWebApp#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#tags WindowsWebApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#timeouts WindowsWebApp#timeouts}
	Timeouts *WindowsWebAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#virtual_network_backup_restore_enabled WindowsWebApp#virtual_network_backup_restore_enabled}.
	VirtualNetworkBackupRestoreEnabled interface{} `field:"optional" json:"virtualNetworkBackupRestoreEnabled" yaml:"virtualNetworkBackupRestoreEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#virtual_network_subnet_id WindowsWebApp#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#webdeploy_publish_basic_authentication_enabled WindowsWebApp#webdeploy_publish_basic_authentication_enabled}.
	WebdeployPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"webdeployPublishBasicAuthenticationEnabled" yaml:"webdeployPublishBasicAuthenticationEnabled"`
	// The local path and filename of the Zip packaged application to deploy to this Windows Web App.
	//
	// **Note:** Using this value requires either `WEBSITE_RUN_FROM_PACKAGE=1` or `SCM_DO_BUILD_DURING_DEPLOYMENT=true` to be set on the App in `app_settings`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/windows_web_app#zip_deploy_file WindowsWebApp#zip_deploy_file}
	ZipDeployFile *string `field:"optional" json:"zipDeployFile" yaml:"zipDeployFile"`
}

