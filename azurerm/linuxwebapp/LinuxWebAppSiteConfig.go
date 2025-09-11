// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxwebapp


type LinuxWebAppSiteConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#always_on LinuxWebApp#always_on}.
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#api_definition_url LinuxWebApp#api_definition_url}.
	ApiDefinitionUrl *string `field:"optional" json:"apiDefinitionUrl" yaml:"apiDefinitionUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#api_management_api_id LinuxWebApp#api_management_api_id}.
	ApiManagementApiId *string `field:"optional" json:"apiManagementApiId" yaml:"apiManagementApiId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#app_command_line LinuxWebApp#app_command_line}.
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// application_stack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#application_stack LinuxWebApp#application_stack}
	ApplicationStack *LinuxWebAppSiteConfigApplicationStack `field:"optional" json:"applicationStack" yaml:"applicationStack"`
	// auto_heal_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#auto_heal_setting LinuxWebApp#auto_heal_setting}
	AutoHealSetting *LinuxWebAppSiteConfigAutoHealSetting `field:"optional" json:"autoHealSetting" yaml:"autoHealSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#container_registry_managed_identity_client_id LinuxWebApp#container_registry_managed_identity_client_id}.
	ContainerRegistryManagedIdentityClientId *string `field:"optional" json:"containerRegistryManagedIdentityClientId" yaml:"containerRegistryManagedIdentityClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#container_registry_use_managed_identity LinuxWebApp#container_registry_use_managed_identity}.
	ContainerRegistryUseManagedIdentity interface{} `field:"optional" json:"containerRegistryUseManagedIdentity" yaml:"containerRegistryUseManagedIdentity"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#cors LinuxWebApp#cors}
	Cors *LinuxWebAppSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#default_documents LinuxWebApp#default_documents}.
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#ftps_state LinuxWebApp#ftps_state}.
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// The amount of time in minutes that a node is unhealthy before being removed from the load balancer.
	//
	// Possible values are between `2` and `10`. Only valid in conjunction with `health_check_path`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#health_check_eviction_time_in_min LinuxWebApp#health_check_eviction_time_in_min}
	HealthCheckEvictionTimeInMin *float64 `field:"optional" json:"healthCheckEvictionTimeInMin" yaml:"healthCheckEvictionTimeInMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#health_check_path LinuxWebApp#health_check_path}.
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#http2_enabled LinuxWebApp#http2_enabled}.
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#ip_restriction LinuxWebApp#ip_restriction}
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#ip_restriction_default_action LinuxWebApp#ip_restriction_default_action}.
	IpRestrictionDefaultAction *string `field:"optional" json:"ipRestrictionDefaultAction" yaml:"ipRestrictionDefaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#load_balancing_mode LinuxWebApp#load_balancing_mode}.
	LoadBalancingMode *string `field:"optional" json:"loadBalancingMode" yaml:"loadBalancingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#local_mysql_enabled LinuxWebApp#local_mysql_enabled}.
	LocalMysqlEnabled interface{} `field:"optional" json:"localMysqlEnabled" yaml:"localMysqlEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#managed_pipeline_mode LinuxWebApp#managed_pipeline_mode}.
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#minimum_tls_version LinuxWebApp#minimum_tls_version}.
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#remote_debugging_enabled LinuxWebApp#remote_debugging_enabled}.
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#remote_debugging_version LinuxWebApp#remote_debugging_version}.
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// scm_ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#scm_ip_restriction LinuxWebApp#scm_ip_restriction}
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#scm_ip_restriction_default_action LinuxWebApp#scm_ip_restriction_default_action}.
	ScmIpRestrictionDefaultAction *string `field:"optional" json:"scmIpRestrictionDefaultAction" yaml:"scmIpRestrictionDefaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#scm_minimum_tls_version LinuxWebApp#scm_minimum_tls_version}.
	ScmMinimumTlsVersion *string `field:"optional" json:"scmMinimumTlsVersion" yaml:"scmMinimumTlsVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#scm_use_main_ip_restriction LinuxWebApp#scm_use_main_ip_restriction}.
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#use_32_bit_worker LinuxWebApp#use_32_bit_worker}.
	Use32BitWorker interface{} `field:"optional" json:"use32BitWorker" yaml:"use32BitWorker"`
	// Should all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#vnet_route_all_enabled LinuxWebApp#vnet_route_all_enabled}
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#websockets_enabled LinuxWebApp#websockets_enabled}.
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/linux_web_app#worker_count LinuxWebApp#worker_count}.
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

