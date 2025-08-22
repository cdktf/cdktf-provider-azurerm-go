// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionapp


type WindowsFunctionAppSiteConfig struct {
	// If this Windows Web App is Always On enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#always_on WindowsFunctionApp#always_on}
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// The URL of the API definition that describes this Windows Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#api_definition_url WindowsFunctionApp#api_definition_url}
	ApiDefinitionUrl *string `field:"optional" json:"apiDefinitionUrl" yaml:"apiDefinitionUrl"`
	// The ID of the API Management API for this Windows Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#api_management_api_id WindowsFunctionApp#api_management_api_id}
	ApiManagementApiId *string `field:"optional" json:"apiManagementApiId" yaml:"apiManagementApiId"`
	// The program and any arguments used to launch this app via the command line. (Example `node myapp.js`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#app_command_line WindowsFunctionApp#app_command_line}
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// The Connection String for linking the Windows Function App to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#application_insights_connection_string WindowsFunctionApp#application_insights_connection_string}
	ApplicationInsightsConnectionString *string `field:"optional" json:"applicationInsightsConnectionString" yaml:"applicationInsightsConnectionString"`
	// The Instrumentation Key for connecting the Windows Function App to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#application_insights_key WindowsFunctionApp#application_insights_key}
	ApplicationInsightsKey *string `field:"optional" json:"applicationInsightsKey" yaml:"applicationInsightsKey"`
	// application_stack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#application_stack WindowsFunctionApp#application_stack}
	ApplicationStack *WindowsFunctionAppSiteConfigApplicationStack `field:"optional" json:"applicationStack" yaml:"applicationStack"`
	// The number of workers this function app can scale out to.
	//
	// Only applicable to apps on the Consumption and Premium plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#app_scale_limit WindowsFunctionApp#app_scale_limit}
	AppScaleLimit *float64 `field:"optional" json:"appScaleLimit" yaml:"appScaleLimit"`
	// app_service_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#app_service_logs WindowsFunctionApp#app_service_logs}
	AppServiceLogs *WindowsFunctionAppSiteConfigAppServiceLogs `field:"optional" json:"appServiceLogs" yaml:"appServiceLogs"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#cors WindowsFunctionApp#cors}
	Cors *WindowsFunctionAppSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Specifies a list of Default Documents for the Windows Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#default_documents WindowsFunctionApp#default_documents}
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// The number of minimum instances for this Windows Function App. Only affects apps on Elastic Premium plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#elastic_instance_minimum WindowsFunctionApp#elastic_instance_minimum}
	ElasticInstanceMinimum *float64 `field:"optional" json:"elasticInstanceMinimum" yaml:"elasticInstanceMinimum"`
	// State of FTP / FTPS service for this function app.
	//
	// Possible values include: `AllAllowed`, `FtpsOnly` and `Disabled`. Defaults to `Disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#ftps_state WindowsFunctionApp#ftps_state}
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// The amount of time in minutes that a node is unhealthy before being removed from the load balancer.
	//
	// Possible values are between `2` and `10`. Only valid in conjunction with `health_check_path`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#health_check_eviction_time_in_min WindowsFunctionApp#health_check_eviction_time_in_min}
	HealthCheckEvictionTimeInMin *float64 `field:"optional" json:"healthCheckEvictionTimeInMin" yaml:"healthCheckEvictionTimeInMin"`
	// The path to be checked for this function app health.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#health_check_path WindowsFunctionApp#health_check_path}
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Specifies if the http2 protocol should be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#http2_enabled WindowsFunctionApp#http2_enabled}
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#ip_restriction WindowsFunctionApp#ip_restriction}
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#ip_restriction_default_action WindowsFunctionApp#ip_restriction_default_action}.
	IpRestrictionDefaultAction *string `field:"optional" json:"ipRestrictionDefaultAction" yaml:"ipRestrictionDefaultAction"`
	// The Site load balancing mode. Possible values include: `WeightedRoundRobin`, `LeastRequests`, `LeastResponseTime`, `WeightedTotalTraffic`, `RequestHash`, `PerSiteRoundRobin`. Defaults to `LeastRequests` if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#load_balancing_mode WindowsFunctionApp#load_balancing_mode}
	LoadBalancingMode *string `field:"optional" json:"loadBalancingMode" yaml:"loadBalancingMode"`
	// The Managed Pipeline mode. Possible values include: `Integrated`, `Classic`. Defaults to `Integrated`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#managed_pipeline_mode WindowsFunctionApp#managed_pipeline_mode}
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// The configures the minimum version of TLS required for SSL requests.
	//
	// Possible values include: `1.0`, `1.1`, and  `1.2`. Defaults to `1.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#minimum_tls_version WindowsFunctionApp#minimum_tls_version}
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// The number of pre-warmed instances for this function app. Only affects apps on an Elastic Premium plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#pre_warmed_instance_count WindowsFunctionApp#pre_warmed_instance_count}
	PreWarmedInstanceCount *float64 `field:"optional" json:"preWarmedInstanceCount" yaml:"preWarmedInstanceCount"`
	// Should Remote Debugging be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#remote_debugging_enabled WindowsFunctionApp#remote_debugging_enabled}
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// The Remote Debugging Version. Currently only `VS2022` is supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#remote_debugging_version WindowsFunctionApp#remote_debugging_version}
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// Should Functions Runtime Scale Monitoring be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#runtime_scale_monitoring_enabled WindowsFunctionApp#runtime_scale_monitoring_enabled}
	RuntimeScaleMonitoringEnabled interface{} `field:"optional" json:"runtimeScaleMonitoringEnabled" yaml:"runtimeScaleMonitoringEnabled"`
	// scm_ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#scm_ip_restriction WindowsFunctionApp#scm_ip_restriction}
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#scm_ip_restriction_default_action WindowsFunctionApp#scm_ip_restriction_default_action}.
	ScmIpRestrictionDefaultAction *string `field:"optional" json:"scmIpRestrictionDefaultAction" yaml:"scmIpRestrictionDefaultAction"`
	// Configures the minimum version of TLS required for SSL requests to the SCM site Possible values include: `1.0`, `1.1`, `1.2` and  `1.3`. Defaults to `1.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#scm_minimum_tls_version WindowsFunctionApp#scm_minimum_tls_version}
	ScmMinimumTlsVersion *string `field:"optional" json:"scmMinimumTlsVersion" yaml:"scmMinimumTlsVersion"`
	// Should the Windows Function App `ip_restriction` configuration be used for the SCM also.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#scm_use_main_ip_restriction WindowsFunctionApp#scm_use_main_ip_restriction}
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Should the Windows Web App use a 32-bit worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#use_32_bit_worker WindowsFunctionApp#use_32_bit_worker}
	Use32BitWorker interface{} `field:"optional" json:"use32BitWorker" yaml:"use32BitWorker"`
	// Should all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#vnet_route_all_enabled WindowsFunctionApp#vnet_route_all_enabled}
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Should Web Sockets be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#websockets_enabled WindowsFunctionApp#websockets_enabled}
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// The number of Workers for this Windows Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/windows_function_app#worker_count WindowsFunctionApp#worker_count}
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

