// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudcontainerdeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudContainerDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#image SpringCloudContainerDeployment#image}.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#name SpringCloudContainerDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#server SpringCloudContainerDeployment#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#spring_cloud_app_id SpringCloudContainerDeployment#spring_cloud_app_id}.
	SpringCloudAppId *string `field:"required" json:"springCloudAppId" yaml:"springCloudAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#addon_json SpringCloudContainerDeployment#addon_json}.
	AddonJson *string `field:"optional" json:"addonJson" yaml:"addonJson"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#application_performance_monitoring_ids SpringCloudContainerDeployment#application_performance_monitoring_ids}.
	ApplicationPerformanceMonitoringIds *[]*string `field:"optional" json:"applicationPerformanceMonitoringIds" yaml:"applicationPerformanceMonitoringIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#arguments SpringCloudContainerDeployment#arguments}.
	Arguments *[]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#commands SpringCloudContainerDeployment#commands}.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#environment_variables SpringCloudContainerDeployment#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#id SpringCloudContainerDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#instance_count SpringCloudContainerDeployment#instance_count}.
	InstanceCount *float64 `field:"optional" json:"instanceCount" yaml:"instanceCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#language_framework SpringCloudContainerDeployment#language_framework}.
	LanguageFramework *string `field:"optional" json:"languageFramework" yaml:"languageFramework"`
	// quota block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#quota SpringCloudContainerDeployment#quota}
	Quota *SpringCloudContainerDeploymentQuota `field:"optional" json:"quota" yaml:"quota"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/spring_cloud_container_deployment#timeouts SpringCloudContainerDeployment#timeouts}
	Timeouts *SpringCloudContainerDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

