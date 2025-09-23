// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppFlexConsumptionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#location FunctionAppFlexConsumption#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Specifies the name of the Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#name FunctionAppFlexConsumption#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#resource_group_name FunctionAppFlexConsumption#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#runtime_name FunctionAppFlexConsumption#runtime_name}.
	RuntimeName *string `field:"required" json:"runtimeName" yaml:"runtimeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#runtime_version FunctionAppFlexConsumption#runtime_version}.
	RuntimeVersion *string `field:"required" json:"runtimeVersion" yaml:"runtimeVersion"`
	// The ID of the App Service Plan within which to create this Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#service_plan_id FunctionAppFlexConsumption#service_plan_id}
	ServicePlanId *string `field:"required" json:"servicePlanId" yaml:"servicePlanId"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#site_config FunctionAppFlexConsumption#site_config}
	SiteConfig *FunctionAppFlexConsumptionSiteConfig `field:"required" json:"siteConfig" yaml:"siteConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#storage_authentication_type FunctionAppFlexConsumption#storage_authentication_type}.
	StorageAuthenticationType *string `field:"required" json:"storageAuthenticationType" yaml:"storageAuthenticationType"`
	// The endpoint of the storage container where the function app's code is hosted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#storage_container_endpoint FunctionAppFlexConsumption#storage_container_endpoint}
	StorageContainerEndpoint *string `field:"required" json:"storageContainerEndpoint" yaml:"storageContainerEndpoint"`
	// The type of the storage container where the function app's code is hosted. Only `blobContainer` is supported currently.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#storage_container_type FunctionAppFlexConsumption#storage_container_type}
	StorageContainerType *string `field:"required" json:"storageContainerType" yaml:"storageContainerType"`
	// always_ready block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#always_ready FunctionAppFlexConsumption#always_ready}
	AlwaysReady interface{} `field:"optional" json:"alwaysReady" yaml:"alwaysReady"`
	// A map of key-value pairs for [App Settings](https://docs.microsoft.com/en-us/azure/azure-functions/functions-app-settings) and custom values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#app_settings FunctionAppFlexConsumption#app_settings}
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// auth_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#auth_settings FunctionAppFlexConsumption#auth_settings}
	AuthSettings *FunctionAppFlexConsumptionAuthSettings `field:"optional" json:"authSettings" yaml:"authSettings"`
	// auth_settings_v2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#auth_settings_v2 FunctionAppFlexConsumption#auth_settings_v2}
	AuthSettingsV2 *FunctionAppFlexConsumptionAuthSettingsV2 `field:"optional" json:"authSettingsV2" yaml:"authSettingsV2"`
	// Should the function app use Client Certificates.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#client_certificate_enabled FunctionAppFlexConsumption#client_certificate_enabled}
	ClientCertificateEnabled interface{} `field:"optional" json:"clientCertificateEnabled" yaml:"clientCertificateEnabled"`
	// Paths to exclude when using client certificates, separated by ;
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#client_certificate_exclusion_paths FunctionAppFlexConsumption#client_certificate_exclusion_paths}
	ClientCertificateExclusionPaths *string `field:"optional" json:"clientCertificateExclusionPaths" yaml:"clientCertificateExclusionPaths"`
	// The mode of the Function App's client certificates requirement for incoming requests.
	//
	// Possible values are `Required`, `Optional`, and `OptionalInteractiveUser`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#client_certificate_mode FunctionAppFlexConsumption#client_certificate_mode}
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#connection_string FunctionAppFlexConsumption#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Is the Function App enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#enabled FunctionAppFlexConsumption#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Can the Function App only be accessed via HTTPS?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#https_only FunctionAppFlexConsumption#https_only}
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#id FunctionAppFlexConsumption#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#identity FunctionAppFlexConsumption#identity}
	Identity *FunctionAppFlexConsumptionIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#instance_memory_in_mb FunctionAppFlexConsumption#instance_memory_in_mb}.
	InstanceMemoryInMb *float64 `field:"optional" json:"instanceMemoryInMb" yaml:"instanceMemoryInMb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#maximum_instance_count FunctionAppFlexConsumption#maximum_instance_count}.
	MaximumInstanceCount *float64 `field:"optional" json:"maximumInstanceCount" yaml:"maximumInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#public_network_access_enabled FunctionAppFlexConsumption#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// sticky_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#sticky_settings FunctionAppFlexConsumption#sticky_settings}
	StickySettings *FunctionAppFlexConsumptionStickySettings `field:"optional" json:"stickySettings" yaml:"stickySettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#storage_access_key FunctionAppFlexConsumption#storage_access_key}.
	StorageAccessKey *string `field:"optional" json:"storageAccessKey" yaml:"storageAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#storage_user_assigned_identity_id FunctionAppFlexConsumption#storage_user_assigned_identity_id}.
	StorageUserAssignedIdentityId *string `field:"optional" json:"storageUserAssignedIdentityId" yaml:"storageUserAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#tags FunctionAppFlexConsumption#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#timeouts FunctionAppFlexConsumption#timeouts}
	Timeouts *FunctionAppFlexConsumptionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#virtual_network_subnet_id FunctionAppFlexConsumption#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#webdeploy_publish_basic_authentication_enabled FunctionAppFlexConsumption#webdeploy_publish_basic_authentication_enabled}.
	WebdeployPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"webdeployPublishBasicAuthenticationEnabled" yaml:"webdeployPublishBasicAuthenticationEnabled"`
	// The local path and filename of the Zip packaged application to deploy to this Function App.
	//
	// **Note:** Using this value requires either `WEBSITE_RUN_FROM_PACKAGE=1` or `SCM_DO_BUILD_DURING_DEPLOYMENT=true` to be set on the App in `app_settings`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/function_app_flex_consumption#zip_deploy_file FunctionAppFlexConsumption#zip_deploy_file}
	ZipDeployFile *string `field:"optional" json:"zipDeployFile" yaml:"zipDeployFile"`
}

