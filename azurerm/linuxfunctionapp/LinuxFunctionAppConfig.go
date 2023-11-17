// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxfunctionapp

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxFunctionAppConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#location LinuxFunctionApp#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Specifies the name of the Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#name LinuxFunctionApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#resource_group_name LinuxFunctionApp#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// The ID of the App Service Plan within which to create this Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#service_plan_id LinuxFunctionApp#service_plan_id}
	ServicePlanId *string `field:"required" json:"servicePlanId" yaml:"servicePlanId"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#site_config LinuxFunctionApp#site_config}
	SiteConfig *LinuxFunctionAppSiteConfig `field:"required" json:"siteConfig" yaml:"siteConfig"`
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#app_settings LinuxFunctionApp#app_settings}
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#auth_settings LinuxFunctionApp#auth_settings}
	AuthSettings *LinuxFunctionAppAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// auth_settings_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#auth_settings_v2 LinuxFunctionApp#auth_settings_v2}
	AuthSettingsV2 *LinuxFunctionAppAuthSettingsV2 `field:"optional" json:"authSettingsV2" yaml:"authSettingsV2"`
	// backup block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#backup LinuxFunctionApp#backup}
	Backup *LinuxFunctionAppBackup `field:"optional" json:"backup" yaml:"backup"`
	// Should built in logging be enabled. Configures `AzureWebJobsDashboard` app setting based on the configured storage setting.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#builtin_logging_enabled LinuxFunctionApp#builtin_logging_enabled}
	BuiltinLoggingEnabled interface{} `field:"optional" json:"builtinLoggingEnabled" yaml:"builtinLoggingEnabled"`
	// Should the function app use Client Certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#client_certificate_enabled LinuxFunctionApp#client_certificate_enabled}
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Paths to exclude when using client certificates, separated by ;
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#client_certificate_exclusion_paths LinuxFunctionApp#client_certificate_exclusion_paths}
	ClientCertificateExclusionPaths *string `field:"optional" json:"clientCertificateExclusionPaths" yaml:"clientCertificateExclusionPaths"`
	// The mode of the Function App's client certificates requirement for incoming requests.
	//
	// Possible values are `Required`, `Optional`, and `OptionalInteractiveUser`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#client_certificate_mode LinuxFunctionApp#client_certificate_mode}
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#connection_string LinuxFunctionApp#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Force disable the content share settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#content_share_force_disabled LinuxFunctionApp#content_share_force_disabled}
	ContentShareForceDisabled interface{} `field:"optional" json:"contentShareForceDisabled" yaml:"contentShareForceDisabled"`
	// The amount of memory in gigabyte-seconds that your application is allowed to consume per day.
	//
	// Setting this value only affects function apps in Consumption Plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#daily_memory_time_quota LinuxFunctionApp#daily_memory_time_quota}
	DailyMemoryTimeQuota *float64 `field:"optional" json:"dailyMemoryTimeQuota" yaml:"dailyMemoryTimeQuota"`
	// Is the Linux Function App enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#enabled LinuxFunctionApp#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#ftp_publish_basic_authentication_enabled LinuxFunctionApp#ftp_publish_basic_authentication_enabled}.
	FtpPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"ftpPublishBasicAuthenticationEnabled" yaml:"ftpPublishBasicAuthenticationEnabled"`
	// The runtime version associated with the Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#functions_extension_version LinuxFunctionApp#functions_extension_version}
	FunctionsExtensionVersion *string `field:"optional" json:"functionsExtensionVersion" yaml:"functionsExtensionVersion"`
	// Can the Function App only be accessed via HTTPS?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#https_only LinuxFunctionApp#https_only}
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#id LinuxFunctionApp#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#identity LinuxFunctionApp#identity}
	Identity *LinuxFunctionAppIdentity `field:"optional" json:"identity" yaml:"identity"`
	// The User Assigned Identity to use for Key Vault access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#key_vault_reference_identity_id LinuxFunctionApp#key_vault_reference_identity_id}
	KeyVaultReferenceIdentityId *string `field:"optional" json:"keyVaultReferenceIdentityId" yaml:"keyVaultReferenceIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#public_network_access_enabled LinuxFunctionApp#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// sticky_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#sticky_settings LinuxFunctionApp#sticky_settings}
	StickySettings *LinuxFunctionAppStickySettings `field:"optional" json:"stickySettings" yaml:"stickySettings"`
	// storage_account block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#storage_account LinuxFunctionApp#storage_account}
	StorageAccount interface{} `field:"optional" json:"storageAccount" yaml:"storageAccount"`
	// The access key which will be used to access the storage account for the Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#storage_account_access_key LinuxFunctionApp#storage_account_access_key}
	StorageAccountAccessKey *string `field:"optional" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// The backend storage account name which will be used by this Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#storage_account_name LinuxFunctionApp#storage_account_name}
	StorageAccountName *string `field:"optional" json:"storageAccountName" yaml:"storageAccountName"`
	// The Key Vault Secret ID, including version, that contains the Connection String to connect to the storage account for this Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#storage_key_vault_secret_id LinuxFunctionApp#storage_key_vault_secret_id}
	StorageKeyVaultSecretId *string `field:"optional" json:"storageKeyVaultSecretId" yaml:"storageKeyVaultSecretId"`
	// Should the Function App use its Managed Identity to access storage?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#storage_uses_managed_identity LinuxFunctionApp#storage_uses_managed_identity}
	StorageUsesManagedIdentity interface{} `field:"optional" json:"storageUsesManagedIdentity" yaml:"storageUsesManagedIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#tags LinuxFunctionApp#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#timeouts LinuxFunctionApp#timeouts}
	Timeouts *LinuxFunctionAppTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#virtual_network_subnet_id LinuxFunctionApp#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#webdeploy_publish_basic_authentication_enabled LinuxFunctionApp#webdeploy_publish_basic_authentication_enabled}.
	WebdeployPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"webdeployPublishBasicAuthenticationEnabled" yaml:"webdeployPublishBasicAuthenticationEnabled"`
	// The local path and filename of the Zip packaged application to deploy to this Linux Function App.
	//
	// **Note:** Using this value requires either `WEBSITE_RUN_FROM_PACKAGE=1` or `SCM_DO_BUILD_DURING_DEPLOYMENT=true` to be set on the App in `app_settings`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.81.0/docs/resources/linux_function_app#zip_deploy_file LinuxFunctionApp#zip_deploy_file}
	ZipDeployFile *string `field:"optional" json:"zipDeployFile" yaml:"zipDeployFile"`
}

