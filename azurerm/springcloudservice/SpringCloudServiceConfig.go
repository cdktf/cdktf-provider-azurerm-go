// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#location SpringCloudService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#name SpringCloudService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#resource_group_name SpringCloudService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#build_agent_pool_size SpringCloudService#build_agent_pool_size}.
	BuildAgentPoolSize *string `field:"optional" json:"buildAgentPoolSize" yaml:"buildAgentPoolSize"`
	// config_server_git_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#config_server_git_setting SpringCloudService#config_server_git_setting}
	ConfigServerGitSetting *SpringCloudServiceConfigServerGitSetting `field:"optional" json:"configServerGitSetting" yaml:"configServerGitSetting"`
	// container_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#container_registry SpringCloudService#container_registry}
	ContainerRegistry interface{} `field:"optional" json:"containerRegistry" yaml:"containerRegistry"`
	// default_build_service block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#default_build_service SpringCloudService#default_build_service}
	DefaultBuildService *SpringCloudServiceDefaultBuildService `field:"optional" json:"defaultBuildService" yaml:"defaultBuildService"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#id SpringCloudService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#log_stream_public_endpoint_enabled SpringCloudService#log_stream_public_endpoint_enabled}.
	LogStreamPublicEndpointEnabled interface{} `field:"optional" json:"logStreamPublicEndpointEnabled" yaml:"logStreamPublicEndpointEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#managed_environment_id SpringCloudService#managed_environment_id}.
	ManagedEnvironmentId *string `field:"optional" json:"managedEnvironmentId" yaml:"managedEnvironmentId"`
	// marketplace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#marketplace SpringCloudService#marketplace}
	Marketplace *SpringCloudServiceMarketplace `field:"optional" json:"marketplace" yaml:"marketplace"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#network SpringCloudService#network}
	Network *SpringCloudServiceNetwork `field:"optional" json:"network" yaml:"network"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#service_registry_enabled SpringCloudService#service_registry_enabled}.
	ServiceRegistryEnabled interface{} `field:"optional" json:"serviceRegistryEnabled" yaml:"serviceRegistryEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#sku_name SpringCloudService#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#sku_tier SpringCloudService#sku_tier}.
	SkuTier *string `field:"optional" json:"skuTier" yaml:"skuTier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#tags SpringCloudService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#timeouts SpringCloudService#timeouts}
	Timeouts *SpringCloudServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// trace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#trace SpringCloudService#trace}
	Trace *SpringCloudServiceTrace `field:"optional" json:"trace" yaml:"trace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/spring_cloud_service#zone_redundant SpringCloudService#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

