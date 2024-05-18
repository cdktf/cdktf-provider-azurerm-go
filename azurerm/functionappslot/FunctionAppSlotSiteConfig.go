// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot


type FunctionAppSlotSiteConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#always_on FunctionAppSlot#always_on}.
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#app_scale_limit FunctionAppSlot#app_scale_limit}.
	AppScaleLimit *float64 `field:"optional" json:"appScaleLimit" yaml:"appScaleLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#auto_swap_slot_name FunctionAppSlot#auto_swap_slot_name}.
	AutoSwapSlotName *string `field:"optional" json:"autoSwapSlotName" yaml:"autoSwapSlotName"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#cors FunctionAppSlot#cors}
	Cors *FunctionAppSlotSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#dotnet_framework_version FunctionAppSlot#dotnet_framework_version}.
	DotnetFrameworkVersion *string `field:"optional" json:"dotnetFrameworkVersion" yaml:"dotnetFrameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#elastic_instance_minimum FunctionAppSlot#elastic_instance_minimum}.
	ElasticInstanceMinimum *float64 `field:"optional" json:"elasticInstanceMinimum" yaml:"elasticInstanceMinimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#ftps_state FunctionAppSlot#ftps_state}.
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#health_check_path FunctionAppSlot#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#http2_enabled FunctionAppSlot#http2_enabled}.
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#ip_restriction FunctionAppSlot#ip_restriction}.
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#java_version FunctionAppSlot#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#linux_fx_version FunctionAppSlot#linux_fx_version}.
	LinuxFxVersion *string `field:"optional" json:"linuxFxVersion" yaml:"linuxFxVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#min_tls_version FunctionAppSlot#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#pre_warmed_instance_count FunctionAppSlot#pre_warmed_instance_count}.
	PreWarmedInstanceCount *float64 `field:"optional" json:"preWarmedInstanceCount" yaml:"preWarmedInstanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#runtime_scale_monitoring_enabled FunctionAppSlot#runtime_scale_monitoring_enabled}.
	RuntimeScaleMonitoringEnabled interface{} `field:"optional" json:"runtimeScaleMonitoringEnabled" yaml:"runtimeScaleMonitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#scm_ip_restriction FunctionAppSlot#scm_ip_restriction}.
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#scm_type FunctionAppSlot#scm_type}.
	ScmType *string `field:"optional" json:"scmType" yaml:"scmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#scm_use_main_ip_restriction FunctionAppSlot#scm_use_main_ip_restriction}.
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#use_32_bit_worker_process FunctionAppSlot#use_32_bit_worker_process}.
	Use32BitWorkerProcess interface{} `field:"optional" json:"use32BitWorkerProcess" yaml:"use32BitWorkerProcess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#vnet_route_all_enabled FunctionAppSlot#vnet_route_all_enabled}.
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/function_app_slot#websockets_enabled FunctionAppSlot#websockets_enabled}.
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
}

