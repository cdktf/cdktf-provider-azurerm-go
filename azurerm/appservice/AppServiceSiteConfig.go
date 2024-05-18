// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservice


type AppServiceSiteConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#acr_use_managed_identity_credentials AppService#acr_use_managed_identity_credentials}.
	AcrUseManagedIdentityCredentials interface{} `field:"optional" json:"acrUseManagedIdentityCredentials" yaml:"acrUseManagedIdentityCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#acr_user_managed_identity_client_id AppService#acr_user_managed_identity_client_id}.
	AcrUserManagedIdentityClientId *string `field:"optional" json:"acrUserManagedIdentityClientId" yaml:"acrUserManagedIdentityClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#always_on AppService#always_on}.
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#app_command_line AppService#app_command_line}.
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#auto_swap_slot_name AppService#auto_swap_slot_name}.
	AutoSwapSlotName *string `field:"optional" json:"autoSwapSlotName" yaml:"autoSwapSlotName"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#cors AppService#cors}
	Cors *AppServiceSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#default_documents AppService#default_documents}.
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#dotnet_framework_version AppService#dotnet_framework_version}.
	DotnetFrameworkVersion *string `field:"optional" json:"dotnetFrameworkVersion" yaml:"dotnetFrameworkVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#ftps_state AppService#ftps_state}.
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#health_check_path AppService#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#http2_enabled AppService#http2_enabled}.
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#ip_restriction AppService#ip_restriction}.
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#java_container AppService#java_container}.
	JavaContainer *string `field:"optional" json:"javaContainer" yaml:"javaContainer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#java_container_version AppService#java_container_version}.
	JavaContainerVersion *string `field:"optional" json:"javaContainerVersion" yaml:"javaContainerVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#java_version AppService#java_version}.
	JavaVersion *string `field:"optional" json:"javaVersion" yaml:"javaVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#linux_fx_version AppService#linux_fx_version}.
	LinuxFxVersion *string `field:"optional" json:"linuxFxVersion" yaml:"linuxFxVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#local_mysql_enabled AppService#local_mysql_enabled}.
	LocalMysqlEnabled interface{} `field:"optional" json:"localMysqlEnabled" yaml:"localMysqlEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#managed_pipeline_mode AppService#managed_pipeline_mode}.
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#min_tls_version AppService#min_tls_version}.
	MinTlsVersion *string `field:"optional" json:"minTlsVersion" yaml:"minTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#number_of_workers AppService#number_of_workers}.
	NumberOfWorkers *float64 `field:"optional" json:"numberOfWorkers" yaml:"numberOfWorkers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#php_version AppService#php_version}.
	PhpVersion *string `field:"optional" json:"phpVersion" yaml:"phpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#python_version AppService#python_version}.
	PythonVersion *string `field:"optional" json:"pythonVersion" yaml:"pythonVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#remote_debugging_enabled AppService#remote_debugging_enabled}.
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#remote_debugging_version AppService#remote_debugging_version}.
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#scm_ip_restriction AppService#scm_ip_restriction}.
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#scm_type AppService#scm_type}.
	ScmType *string `field:"optional" json:"scmType" yaml:"scmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#scm_use_main_ip_restriction AppService#scm_use_main_ip_restriction}.
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#use_32_bit_worker_process AppService#use_32_bit_worker_process}.
	Use32BitWorkerProcess interface{} `field:"optional" json:"use32BitWorkerProcess" yaml:"use32BitWorkerProcess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#vnet_route_all_enabled AppService#vnet_route_all_enabled}.
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#websockets_enabled AppService#websockets_enabled}.
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/app_service#windows_fx_version AppService#windows_fx_version}.
	WindowsFxVersion *string `field:"optional" json:"windowsFxVersion" yaml:"windowsFxVersion"`
}

