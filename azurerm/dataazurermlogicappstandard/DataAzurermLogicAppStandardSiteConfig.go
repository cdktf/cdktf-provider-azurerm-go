// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermlogicappstandard


type DataAzurermLogicAppStandardSiteConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#always_on DataAzurermLogicAppStandard#always_on}.
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#app_scale_limit DataAzurermLogicAppStandard#app_scale_limit}.
	AppScaleLimit *float64 `field:"optional" json:"appScaleLimit" yaml:"appScaleLimit"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#cors DataAzurermLogicAppStandard#cors}
	Cors *DataAzurermLogicAppStandardSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#dotnet_framework_version DataAzurermLogicAppStandard#dotnet_framework_version}.
	DotnetFrameworkVersion *string `field:"optional" json:"dotnetFrameworkVersion" yaml:"dotnetFrameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#elastic_instance_minimum DataAzurermLogicAppStandard#elastic_instance_minimum}.
	ElasticInstanceMinimum *float64 `field:"optional" json:"elasticInstanceMinimum" yaml:"elasticInstanceMinimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#ftps_state DataAzurermLogicAppStandard#ftps_state}.
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#health_check_path DataAzurermLogicAppStandard#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#http2_enabled DataAzurermLogicAppStandard#http2_enabled}.
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#ip_restriction DataAzurermLogicAppStandard#ip_restriction}.
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#linux_fx_version DataAzurermLogicAppStandard#linux_fx_version}.
	LinuxFxVersion *string `field:"optional" json:"linuxFxVersion" yaml:"linuxFxVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#min_tls_version DataAzurermLogicAppStandard#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#pre_warmed_instance_count DataAzurermLogicAppStandard#pre_warmed_instance_count}.
	PreWarmedInstanceCount *float64 `field:"optional" json:"preWarmedInstanceCount" yaml:"preWarmedInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#public_network_access_enabled DataAzurermLogicAppStandard#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#runtime_scale_monitoring_enabled DataAzurermLogicAppStandard#runtime_scale_monitoring_enabled}.
	RuntimeScaleMonitoringEnabled interface{} `field:"optional" json:"runtimeScaleMonitoringEnabled" yaml:"runtimeScaleMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#scm_ip_restriction DataAzurermLogicAppStandard#scm_ip_restriction}.
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#scm_min_tls_version DataAzurermLogicAppStandard#scm_min_tls_version}.
	ScmMinTlsVersion *string `field:"optional" json:"scmMinTlsVersion" yaml:"scmMinTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#scm_type DataAzurermLogicAppStandard#scm_type}.
	ScmType *string `field:"optional" json:"scmType" yaml:"scmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#scm_use_main_ip_restriction DataAzurermLogicAppStandard#scm_use_main_ip_restriction}.
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#use_32_bit_worker_process DataAzurermLogicAppStandard#use_32_bit_worker_process}.
	Use32BitWorkerProcess interface{} `field:"optional" json:"use32BitWorkerProcess" yaml:"use32BitWorkerProcess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#vnet_route_all_enabled DataAzurermLogicAppStandard#vnet_route_all_enabled}.
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/data-sources/logic_app_standard#websockets_enabled DataAzurermLogicAppStandard#websockets_enabled}.
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
}

