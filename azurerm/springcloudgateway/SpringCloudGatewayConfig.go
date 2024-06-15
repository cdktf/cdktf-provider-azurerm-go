// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudgateway

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudGatewayConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#name SpringCloudGateway#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#spring_cloud_service_id SpringCloudGateway#spring_cloud_service_id}.
	SpringCloudServiceId *string `field:"required" json:"springCloudServiceId" yaml:"springCloudServiceId"`
	// api_metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#api_metadata SpringCloudGateway#api_metadata}
	ApiMetadata *SpringCloudGatewayApiMetadata `field:"optional" json:"apiMetadata" yaml:"apiMetadata"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#application_performance_monitoring_ids SpringCloudGateway#application_performance_monitoring_ids}.
	ApplicationPerformanceMonitoringIds *[]*string `field:"optional" json:"applicationPerformanceMonitoringIds" yaml:"applicationPerformanceMonitoringIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#application_performance_monitoring_types SpringCloudGateway#application_performance_monitoring_types}.
	ApplicationPerformanceMonitoringTypes *[]*string `field:"optional" json:"applicationPerformanceMonitoringTypes" yaml:"applicationPerformanceMonitoringTypes"`
	// client_authorization block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#client_authorization SpringCloudGateway#client_authorization}
	ClientAuthorization *SpringCloudGatewayClientAuthorization `field:"optional" json:"clientAuthorization" yaml:"clientAuthorization"`
	// cors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#cors SpringCloudGateway#cors}
	Cors *SpringCloudGatewayCors `field:"optional" json:"cors" yaml:"cors"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#environment_variables SpringCloudGateway#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#https_only SpringCloudGateway#https_only}.
	HttpsOnly interface{} `field:"optional" json:"httpsOnly" yaml:"httpsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#id SpringCloudGateway#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#instance_count SpringCloudGateway#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// local_response_cache_per_instance block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#local_response_cache_per_instance SpringCloudGateway#local_response_cache_per_instance}
	LocalResponseCachePerInstance *SpringCloudGatewayLocalResponseCachePerInstance `field:"optional" json:"localResponseCachePerInstance" yaml:"localResponseCachePerInstance"`
	// local_response_cache_per_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#local_response_cache_per_route SpringCloudGateway#local_response_cache_per_route}
	LocalResponseCachePerRoute *SpringCloudGatewayLocalResponseCachePerRoute `field:"optional" json:"localResponseCachePerRoute" yaml:"localResponseCachePerRoute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#public_network_access_enabled SpringCloudGateway#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#quota SpringCloudGateway#quota}
	Quota *SpringCloudGatewayQuota `field:"optional" json:"quota" yaml:"quota"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#sensitive_environment_variables SpringCloudGateway#sensitive_environment_variables}.
	SensitiveEnvironmentVariables *map[string]*string `field:"optional" json:"sensitiveEnvironmentVariables" yaml:"sensitiveEnvironmentVariables"`
	// sso block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#sso SpringCloudGateway#sso}
	Sso *SpringCloudGatewaySso `field:"optional" json:"sso" yaml:"sso"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/spring_cloud_gateway#timeouts SpringCloudGateway#timeouts}
	Timeouts *SpringCloudGatewayTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

