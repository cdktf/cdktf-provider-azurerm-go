package linuxfunctionappslot


type LinuxFunctionAppSlotSiteConfig struct {
	// If this Linux Web App is Always On enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#always_on LinuxFunctionAppSlot#always_on}
	AlwaysOn interface{} `field:"optional" json:"alwaysOn" yaml:"alwaysOn"`
	// The URL of the API definition that describes this Linux Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#api_definition_url LinuxFunctionAppSlot#api_definition_url}
	ApiDefinitionUrl *string `field:"optional" json:"apiDefinitionUrl" yaml:"apiDefinitionUrl"`
	// The ID of the API Management API for this Linux Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#api_management_api_id LinuxFunctionAppSlot#api_management_api_id}
	ApiManagementApiId *string `field:"optional" json:"apiManagementApiId" yaml:"apiManagementApiId"`
	// The program and any arguments used to launch this app via the command line. (Example `node myapp.js`).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#app_command_line LinuxFunctionAppSlot#app_command_line}
	AppCommandLine *string `field:"optional" json:"appCommandLine" yaml:"appCommandLine"`
	// The Connection String for linking the Linux Function App to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#application_insights_connection_string LinuxFunctionAppSlot#application_insights_connection_string}
	ApplicationInsightsConnectionString *string `field:"optional" json:"applicationInsightsConnectionString" yaml:"applicationInsightsConnectionString"`
	// The Instrumentation Key for connecting the Linux Function App to Application Insights.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#application_insights_key LinuxFunctionAppSlot#application_insights_key}
	ApplicationInsightsKey *string `field:"optional" json:"applicationInsightsKey" yaml:"applicationInsightsKey"`
	// application_stack block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#application_stack LinuxFunctionAppSlot#application_stack}
	ApplicationStack *LinuxFunctionAppSlotSiteConfigApplicationStack `field:"optional" json:"applicationStack" yaml:"applicationStack"`
	// The number of workers this function app can scale out to.
	//
	// Only applicable to apps on the Consumption and Premium plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#app_scale_limit LinuxFunctionAppSlot#app_scale_limit}
	AppScaleLimit *float64 `field:"optional" json:"appScaleLimit" yaml:"appScaleLimit"`
	// app_service_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#app_service_logs LinuxFunctionAppSlot#app_service_logs}
	AppServiceLogs *LinuxFunctionAppSlotSiteConfigAppServiceLogs `field:"optional" json:"appServiceLogs" yaml:"appServiceLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#auto_swap_slot_name LinuxFunctionAppSlot#auto_swap_slot_name}.
	AutoSwapSlotName *string `field:"optional" json:"autoSwapSlotName" yaml:"autoSwapSlotName"`
	// The Client ID of the Managed Service Identity to use for connections to the Azure Container Registry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#container_registry_managed_identity_client_id LinuxFunctionAppSlot#container_registry_managed_identity_client_id}
	ContainerRegistryManagedIdentityClientId *string `field:"optional" json:"containerRegistryManagedIdentityClientId" yaml:"containerRegistryManagedIdentityClientId"`
	// Should connections for Azure Container Registry use Managed Identity.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#container_registry_use_managed_identity LinuxFunctionAppSlot#container_registry_use_managed_identity}
	ContainerRegistryUseManagedIdentity interface{} `field:"optional" json:"containerRegistryUseManagedIdentity" yaml:"containerRegistryUseManagedIdentity"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#cors LinuxFunctionAppSlot#cors}
	Cors *LinuxFunctionAppSlotSiteConfigCors `field:"optional" json:"cors" yaml:"cors"`
	// Specifies a list of Default Documents for the Linux Web App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#default_documents LinuxFunctionAppSlot#default_documents}
	DefaultDocuments *[]*string `field:"optional" json:"defaultDocuments" yaml:"defaultDocuments"`
	// The number of minimum instances for this Linux Function App. Only affects apps on Elastic Premium plans.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#elastic_instance_minimum LinuxFunctionAppSlot#elastic_instance_minimum}
	ElasticInstanceMinimum *float64 `field:"optional" json:"elasticInstanceMinimum" yaml:"elasticInstanceMinimum"`
	// State of FTP / FTPS service for this function app.
	//
	// Possible values include: `AllAllowed`, `FtpsOnly` and `Disabled`. Defaults to `Disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#ftps_state LinuxFunctionAppSlot#ftps_state}
	FtpsState *string `field:"optional" json:"ftpsState" yaml:"ftpsState"`
	// The amount of time in minutes that a node is unhealthy before being removed from the load balancer.
	//
	// Possible values are between `2` and `10`. Defaults to `10`. Only valid in conjunction with `health_check_path`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#health_check_eviction_time_in_min LinuxFunctionAppSlot#health_check_eviction_time_in_min}
	HealthCheckEvictionTimeInMin *float64 `field:"optional" json:"healthCheckEvictionTimeInMin" yaml:"healthCheckEvictionTimeInMin"`
	// The path to be checked for this function app health.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#health_check_path LinuxFunctionAppSlot#health_check_path}
	HealthCheckPath *string `field:"optional" json:"healthCheckPath" yaml:"healthCheckPath"`
	// Specifies if the http2 protocol should be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#http2_enabled LinuxFunctionAppSlot#http2_enabled}
	Http2Enabled interface{} `field:"optional" json:"http2Enabled" yaml:"http2Enabled"`
	// ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#ip_restriction LinuxFunctionAppSlot#ip_restriction}
	IpRestriction interface{} `field:"optional" json:"ipRestriction" yaml:"ipRestriction"`
	// The Site load balancing mode. Possible values include: `WeightedRoundRobin`, `LeastRequests`, `LeastResponseTime`, `WeightedTotalTraffic`, `RequestHash`, `PerSiteRoundRobin`. Defaults to `LeastRequests` if omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#load_balancing_mode LinuxFunctionAppSlot#load_balancing_mode}
	LoadBalancingMode *string `field:"optional" json:"loadBalancingMode" yaml:"loadBalancingMode"`
	// The Managed Pipeline mode. Possible values include: `Integrated`, `Classic`. Defaults to `Integrated`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#managed_pipeline_mode LinuxFunctionAppSlot#managed_pipeline_mode}
	ManagedPipelineMode *string `field:"optional" json:"managedPipelineMode" yaml:"managedPipelineMode"`
	// The configures the minimum version of TLS required for SSL requests.
	//
	// Possible values include: `1.0`, `1.1`, and  `1.2`. Defaults to `1.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#minimum_tls_version LinuxFunctionAppSlot#minimum_tls_version}
	MinimumTlsVersion *string `field:"optional" json:"minimumTlsVersion" yaml:"minimumTlsVersion"`
	// The number of pre-warmed instances for this function app. Only affects apps on an Elastic Premium plan.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#pre_warmed_instance_count LinuxFunctionAppSlot#pre_warmed_instance_count}
	PreWarmedInstanceCount *float64 `field:"optional" json:"preWarmedInstanceCount" yaml:"preWarmedInstanceCount"`
	// Should Remote Debugging be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#remote_debugging_enabled LinuxFunctionAppSlot#remote_debugging_enabled}
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// The Remote Debugging Version. Possible values include `VS2017`, `VS2019`, and `VS2022`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#remote_debugging_version LinuxFunctionAppSlot#remote_debugging_version}
	RemoteDebuggingVersion *string `field:"optional" json:"remoteDebuggingVersion" yaml:"remoteDebuggingVersion"`
	// Should Functions Runtime Scale Monitoring be enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#runtime_scale_monitoring_enabled LinuxFunctionAppSlot#runtime_scale_monitoring_enabled}
	RuntimeScaleMonitoringEnabled interface{} `field:"optional" json:"runtimeScaleMonitoringEnabled" yaml:"runtimeScaleMonitoringEnabled"`
	// scm_ip_restriction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#scm_ip_restriction LinuxFunctionAppSlot#scm_ip_restriction}
	ScmIpRestriction interface{} `field:"optional" json:"scmIpRestriction" yaml:"scmIpRestriction"`
	// Configures the minimum version of TLS required for SSL requests to the SCM site Possible values include: `1.0`, `1.1`, and  `1.2`. Defaults to `1.2`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#scm_minimum_tls_version LinuxFunctionAppSlot#scm_minimum_tls_version}
	ScmMinimumTlsVersion *string `field:"optional" json:"scmMinimumTlsVersion" yaml:"scmMinimumTlsVersion"`
	// Should the Linux Function App `ip_restriction` configuration be used for the SCM also.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#scm_use_main_ip_restriction LinuxFunctionAppSlot#scm_use_main_ip_restriction}
	ScmUseMainIpRestriction interface{} `field:"optional" json:"scmUseMainIpRestriction" yaml:"scmUseMainIpRestriction"`
	// Should the Linux Web App use a 32-bit worker.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#use_32_bit_worker LinuxFunctionAppSlot#use_32_bit_worker}
	Use32BitWorker interface{} `field:"optional" json:"use32BitWorker" yaml:"use32BitWorker"`
	// Should all outbound traffic to have Virtual Network Security Groups and User Defined Routes applied? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#vnet_route_all_enabled LinuxFunctionAppSlot#vnet_route_all_enabled}
	VnetRouteAllEnabled interface{} `field:"optional" json:"vnetRouteAllEnabled" yaml:"vnetRouteAllEnabled"`
	// Should Web Sockets be enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#websockets_enabled LinuxFunctionAppSlot#websockets_enabled}
	WebsocketsEnabled interface{} `field:"optional" json:"websocketsEnabled" yaml:"websocketsEnabled"`
	// The number of Workers for this Linux Function App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_function_app_slot#worker_count LinuxFunctionAppSlot#worker_count}
	WorkerCount *float64 `field:"optional" json:"workerCount" yaml:"workerCount"`
}

