// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appserviceenvironmentv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceEnvironmentV3Config struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#name AppServiceEnvironmentV3#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#resource_group_name AppServiceEnvironmentV3#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#subnet_id AppServiceEnvironmentV3#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#allow_new_private_endpoint_connections AppServiceEnvironmentV3#allow_new_private_endpoint_connections}.
	AllowNewPrivateEndpointConnections interface{} `field:"optional" json:"allowNewPrivateEndpointConnections" yaml:"allowNewPrivateEndpointConnections"`
	// cluster_setting block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#cluster_setting AppServiceEnvironmentV3#cluster_setting}
	ClusterSetting interface{} `field:"optional" json:"clusterSetting" yaml:"clusterSetting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#dedicated_host_count AppServiceEnvironmentV3#dedicated_host_count}.
	DedicatedHostCount *float64 `field:"optional" json:"dedicatedHostCount" yaml:"dedicatedHostCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#id AppServiceEnvironmentV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#internal_load_balancing_mode AppServiceEnvironmentV3#internal_load_balancing_mode}.
	InternalLoadBalancingMode *string `field:"optional" json:"internalLoadBalancingMode" yaml:"internalLoadBalancingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#remote_debugging_enabled AppServiceEnvironmentV3#remote_debugging_enabled}.
	RemoteDebuggingEnabled interface{} `field:"optional" json:"remoteDebuggingEnabled" yaml:"remoteDebuggingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#tags AppServiceEnvironmentV3#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#timeouts AppServiceEnvironmentV3#timeouts}
	Timeouts *AppServiceEnvironmentV3Timeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/app_service_environment_v3#zone_redundant AppServiceEnvironmentV3#zone_redundant}.
	ZoneRedundant interface{} `field:"optional" json:"zoneRedundant" yaml:"zoneRedundant"`
}

