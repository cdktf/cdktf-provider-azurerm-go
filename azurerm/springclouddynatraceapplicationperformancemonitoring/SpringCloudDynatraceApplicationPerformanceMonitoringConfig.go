// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springclouddynatraceapplicationperformancemonitoring

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudDynatraceApplicationPerformanceMonitoringConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#connection_point SpringCloudDynatraceApplicationPerformanceMonitoring#connection_point}.
	ConnectionPoint *string `field:"required" json:"connectionPoint" yaml:"connectionPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#name SpringCloudDynatraceApplicationPerformanceMonitoring#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#spring_cloud_service_id SpringCloudDynatraceApplicationPerformanceMonitoring#spring_cloud_service_id}.
	SpringCloudServiceId *string `field:"required" json:"springCloudServiceId" yaml:"springCloudServiceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#tenant SpringCloudDynatraceApplicationPerformanceMonitoring#tenant}.
	Tenant *string `field:"required" json:"tenant" yaml:"tenant"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#tenant_token SpringCloudDynatraceApplicationPerformanceMonitoring#tenant_token}.
	TenantToken *string `field:"required" json:"tenantToken" yaml:"tenantToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#api_token SpringCloudDynatraceApplicationPerformanceMonitoring#api_token}.
	ApiToken *string `field:"optional" json:"apiToken" yaml:"apiToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#api_url SpringCloudDynatraceApplicationPerformanceMonitoring#api_url}.
	ApiUrl *string `field:"optional" json:"apiUrl" yaml:"apiUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#environment_id SpringCloudDynatraceApplicationPerformanceMonitoring#environment_id}.
	EnvironmentId *string `field:"optional" json:"environmentId" yaml:"environmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#globally_enabled SpringCloudDynatraceApplicationPerformanceMonitoring#globally_enabled}.
	GloballyEnabled interface{} `field:"optional" json:"globallyEnabled" yaml:"globallyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#id SpringCloudDynatraceApplicationPerformanceMonitoring#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/spring_cloud_dynatrace_application_performance_monitoring#timeouts SpringCloudDynatraceApplicationPerformanceMonitoring#timeouts}
	Timeouts *SpringCloudDynatraceApplicationPerformanceMonitoringTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

