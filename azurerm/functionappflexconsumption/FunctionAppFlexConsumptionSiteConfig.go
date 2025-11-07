// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappflexconsumption


type FunctionAppFlexConsumptionSiteConfig struct {
	// The URL of the API definition that describes this Linux Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#api_definition_url FunctionAppFlexConsumption#api_definition_url}
	ApiDefinitionUrl *string `field:"optional" json:"apiDefinitionUrl" yaml:"apiDefinitionUrl"`
	// The ID of the API Management API for this Linux Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#api_management_api_id FunctionAppFlexConsumption#api_management_api_id}
	ApiManagementApiId *string `field:"optional" json:"apiManagementApiId" yaml:"apiManagementApiId"`
	// The program and any arguments used to launch this app via the command line. (Example `node myapp.js`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#app_command_line FunctionAppFlexConsumption#app_command_line}
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// The Connection String for linking the Linux Function App to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#application_insights_connection_string FunctionAppFlexConsumption#application_insights_connection_string}
	ApplicationInsightsConnectionString *string `field:"optional" json:"applicationInsightsConnectionString" yaml:"applicationInsightsConnectionString"`
	// The Instrumentation Key for connecting the Linux Function App to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#application_insights_key FunctionAppFlexConsumption#application_insights_key}
	ApplicationInsightsKey *string `field:"optional" json:"applicationInsightsKey" yaml:"applicationInsightsKey"`
	// app_service_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#app_service_logs FunctionAppFlexConsumption#app_service_logs}
	AppServiceLogs *FunctionAppFlexConsumptionSiteConfigAppServiceLogs `field:"optional" json:"appServiceLogs" yaml:"appServiceLogs"`
	// The Client ID of the Managed Service Identity to use for connections to the Azure Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#container_registry_managed_identity_client_id FunctionAppFlexConsumption#container_registry_managed_identity_client_id}
	ContainerRegistryManagedIdentityClientId *string `field:"optional" json:"containerRegistryManagedIdentityClientId" yaml:"containerRegistryManagedIdentityClientId"`
	// Should connections for Azure Container Registry use Managed Identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#container_registry_use_managed_identity FunctionAppFlexConsumption#container_registry_use_managed_identity}
	ContainerRegistryUseManagedIdentity interface{} `field:"optional" json:"containerRegistryUseManagedIdentity" yaml:"containerRegistryUseManagedIdentity"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#cors FunctionAppFlexConsumption#cors}
	Cors *FunctionAppFlexConsumptionSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Specifies a list of Default Documents for the Linux Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#default_documents FunctionAppFlexConsumption#default_documents}
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// The number of minimum instances for this Linux Function App. Only affects apps on Elastic Premium plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#elastic_instance_minimum FunctionAppFlexConsumption#elastic_instance_minimum}
	ElasticInstanceMinimum *float64 `field:"optional" json:"elasticInstanceMinimum" yaml:"elasticInstanceMinimum"`
	// The amount of time in minutes that a node is unhealthy before being removed from the load balancer.
	//
	// Possible values are between `2` and `10`. Only valid in conjunction with `health_check_path`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#health_check_eviction_time_in_min FunctionAppFlexConsumption#health_check_eviction_time_in_min}
	HealthCheckEvictionTimeInMin *float64 `field:"optional" json:"healthCheckEvictionTimeInMin" yaml:"healthCheckEvictionTimeInMin"`
	// The path to be checked for this function app health.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#health_check_path FunctionAppFlexConsumption#health_check_path}
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Specifies if the http2 protocol should be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#http2_enabled FunctionAppFlexConsumption#http2_enabled}
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#ip_restriction FunctionAppFlexConsumption#ip_restriction}
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#ip_restriction_default_action FunctionAppFlexConsumption#ip_restriction_default_action}.
	IpRestrictionDefaultAction *string `field:"optional" json:"ipRestrictionDefaultAction" yaml:"ipRestrictionDefaultAction"`
	// The Site load balancing mode. Possible values include: `WeightedRoundRobin`, `LeastRequests`, `LeastResponseTime`, `WeightedTotalTraffic`, `RequestHash`, `PerSiteRoundRobin`. Defaults to `LeastRequests` if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#load_balancing_mode FunctionAppFlexConsumption#load_balancing_mode}
	LoadBalancingMode *string `field:"optional" json:"loadBalancingMode" yaml:"loadBalancingMode"`
	// The Managed Pipeline mode. Possible values include: `Integrated`, `Classic`. Defaults to `Integrated`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#managed_pipeline_mode FunctionAppFlexConsumption#managed_pipeline_mode}
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// The configures the minimum version of TLS required for SSL requests.
	//
	// Possible values include: `1.0`, `1.1`, and  `1.2`. Defaults to `1.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#minimum_tls_version FunctionAppFlexConsumption#minimum_tls_version}
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// Should Remote Debugging be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#remote_debugging_enabled FunctionAppFlexConsumption#remote_debugging_enabled}
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// The Remote Debugging Version. Possible values include `VS2017`, `VS2019`, and `VS2022``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#remote_debugging_version FunctionAppFlexConsumption#remote_debugging_version}
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// Should Functions Runtime Scale Monitoring be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#runtime_scale_monitoring_enabled FunctionAppFlexConsumption#runtime_scale_monitoring_enabled}
	RuntimeScaleMonitoringEnabled interface{} `field:"optional" json:"runtimeScaleMonitoringEnabled" yaml:"runtimeScaleMonitoringEnabled"`
	// scm_ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#scm_ip_restriction FunctionAppFlexConsumption#scm_ip_restriction}
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#scm_ip_restriction_default_action FunctionAppFlexConsumption#scm_ip_restriction_default_action}.
	ScmIpRestrictionDefaultAction *string `field:"optional" json:"scmIpRestrictionDefaultAction" yaml:"scmIpRestrictionDefaultAction"`
	// Configures the minimum version of TLS required for SSL requests to the SCM site Possible values include: `1.0`, `1.1`, and  `1.2`. Defaults to `1.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#scm_minimum_tls_version FunctionAppFlexConsumption#scm_minimum_tls_version}
	ScmMinimumTlsVersion *string `field:"optional" json:"scmMinimumTlsVersion" yaml:"scmMinimumTlsVersion"`
	// Should the Linux Function App `ip_restriction` configuration be used for the SCM also.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#scm_use_main_ip_restriction FunctionAppFlexConsumption#scm_use_main_ip_restriction}
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Should the Linux Function App use a 32-bit worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#use_32_bit_worker FunctionAppFlexConsumption#use_32_bit_worker}
	Use32BitWorker interface{} `field:"optional" json:"use32BitWorker" yaml:"use32BitWorker"`
	// Should the Linux Function App route all traffic through the virtual network.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#vnet_route_all_enabled FunctionAppFlexConsumption#vnet_route_all_enabled}
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Should Web Sockets be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#websockets_enabled FunctionAppFlexConsumption#websockets_enabled}
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// The number of Workers for this Linux Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/function_app_flex_consumption#worker_count FunctionAppFlexConsumption#worker_count}
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

