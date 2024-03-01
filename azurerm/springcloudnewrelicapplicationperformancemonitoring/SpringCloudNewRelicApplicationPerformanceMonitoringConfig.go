// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudnewrelicapplicationperformancemonitoring

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudNewRelicApplicationPerformanceMonitoringConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#app_name SpringCloudNewRelicApplicationPerformanceMonitoring#app_name}.
	AppName *string `field:"required" json:"appName" yaml:"appName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#license_key SpringCloudNewRelicApplicationPerformanceMonitoring#license_key}.
	LicenseKey *string `field:"required" json:"licenseKey" yaml:"licenseKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#name SpringCloudNewRelicApplicationPerformanceMonitoring#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#spring_cloud_service_id SpringCloudNewRelicApplicationPerformanceMonitoring#spring_cloud_service_id}.
	SpringCloudServiceId *string `field:"required" json:"springCloudServiceId" yaml:"springCloudServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#agent_enabled SpringCloudNewRelicApplicationPerformanceMonitoring#agent_enabled}.
	AgentEnabled interface{} `field:"optional" json:"agentEnabled" yaml:"agentEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#app_server_port SpringCloudNewRelicApplicationPerformanceMonitoring#app_server_port}.
	AppServerPort *float64 `field:"optional" json:"appServerPort" yaml:"appServerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#audit_mode_enabled SpringCloudNewRelicApplicationPerformanceMonitoring#audit_mode_enabled}.
	AuditModeEnabled interface{} `field:"optional" json:"auditModeEnabled" yaml:"auditModeEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#auto_app_naming_enabled SpringCloudNewRelicApplicationPerformanceMonitoring#auto_app_naming_enabled}.
	AutoAppNamingEnabled interface{} `field:"optional" json:"autoAppNamingEnabled" yaml:"autoAppNamingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#auto_transaction_naming_enabled SpringCloudNewRelicApplicationPerformanceMonitoring#auto_transaction_naming_enabled}.
	AutoTransactionNamingEnabled interface{} `field:"optional" json:"autoTransactionNamingEnabled" yaml:"autoTransactionNamingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#custom_tracing_enabled SpringCloudNewRelicApplicationPerformanceMonitoring#custom_tracing_enabled}.
	CustomTracingEnabled interface{} `field:"optional" json:"customTracingEnabled" yaml:"customTracingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#globally_enabled SpringCloudNewRelicApplicationPerformanceMonitoring#globally_enabled}.
	GloballyEnabled interface{} `field:"optional" json:"globallyEnabled" yaml:"globallyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#id SpringCloudNewRelicApplicationPerformanceMonitoring#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#labels SpringCloudNewRelicApplicationPerformanceMonitoring#labels}.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.94.0/docs/resources/spring_cloud_new_relic_application_performance_monitoring#timeouts SpringCloudNewRelicApplicationPerformanceMonitoring#timeouts}
	Timeouts *SpringCloudNewRelicApplicationPerformanceMonitoringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

