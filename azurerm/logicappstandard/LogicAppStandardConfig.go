// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappstandard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppStandardConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#app_service_plan_id LogicAppStandard#app_service_plan_id}.
	AppServicePlanId *string `field:"required" json:"appServicePlanId" yaml:"appServicePlanId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#location LogicAppStandard#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#name LogicAppStandard#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#resource_group_name LogicAppStandard#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#storage_account_access_key LogicAppStandard#storage_account_access_key}.
	StorageAccountAccessKey *string `field:"required" json:"storageAccountAccessKey" yaml:"storageAccountAccessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#storage_account_name LogicAppStandard#storage_account_name}.
	StorageAccountName *string `field:"required" json:"storageAccountName" yaml:"storageAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#app_settings LogicAppStandard#app_settings}.
	AppSettings *map[string]*string `field:"optional" json:"appSettings" yaml:"appSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#bundle_version LogicAppStandard#bundle_version}.
	BundleVersion *string `field:"optional" json:"bundleVersion" yaml:"bundleVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#client_affinity_enabled LogicAppStandard#client_affinity_enabled}.
	ClientAffinityEnabled interface{} `field:"optional" json:"clientAffinityEnabled" yaml:"clientAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#client_certificate_mode LogicAppStandard#client_certificate_mode}.
	ClientCertificateMode *string `field:"optional" json:"clientCertificateMode" yaml:"clientCertificateMode"`
	// connection_string block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#connection_string LogicAppStandard#connection_string}
	ConnectionString interface{} `field:"optional" json:"connectionString" yaml:"connectionString"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#enabled LogicAppStandard#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#ftp_publish_basic_authentication_enabled LogicAppStandard#ftp_publish_basic_authentication_enabled}.
	FtpPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"ftpPublishBasicAuthenticationEnabled" yaml:"ftpPublishBasicAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#https_only LogicAppStandard#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#id LogicAppStandard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#identity LogicAppStandard#identity}
	Identity *LogicAppStandardIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#public_network_access LogicAppStandard#public_network_access}.
	PublicNetworkAccess *string `field:"optional" json:"publicNetworkAccess" yaml:"publicNetworkAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#scm_publish_basic_authentication_enabled LogicAppStandard#scm_publish_basic_authentication_enabled}.
	ScmPublishBasicAuthenticationEnabled interface{} `field:"optional" json:"scmPublishBasicAuthenticationEnabled" yaml:"scmPublishBasicAuthenticationEnabled"`
	// site_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#site_config LogicAppStandard#site_config}
	SiteConfig *LogicAppStandardSiteConfig `field:"optional" json:"siteConfig" yaml:"siteConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#storage_account_share_name LogicAppStandard#storage_account_share_name}.
	StorageAccountShareName *string `field:"optional" json:"storageAccountShareName" yaml:"storageAccountShareName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#tags LogicAppStandard#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#timeouts LogicAppStandard#timeouts}
	Timeouts *LogicAppStandardTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#use_extension_bundle LogicAppStandard#use_extension_bundle}.
	UseExtensionBundle interface{} `field:"optional" json:"useExtensionBundle" yaml:"useExtensionBundle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#version LogicAppStandard#version}.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#virtual_network_subnet_id LogicAppStandard#virtual_network_subnet_id}.
	VirtualNetworkSubnetId *string `field:"optional" json:"virtualNetworkSubnetId" yaml:"virtualNetworkSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/logic_app_standard#vnet_content_share_enabled LogicAppStandard#vnet_content_share_enabled}.
	VnetContentShareEnabled interface{} `field:"optional" json:"vnetContentShareEnabled" yaml:"vnetContentShareEnabled"`
}

