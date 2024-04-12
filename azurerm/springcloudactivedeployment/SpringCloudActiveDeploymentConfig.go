// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudactivedeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudActiveDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/spring_cloud_active_deployment#deployment_name SpringCloudActiveDeployment#deployment_name}.
	DeploymentName *string `field:"required" json:"deploymentName" yaml:"deploymentName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/spring_cloud_active_deployment#spring_cloud_app_id SpringCloudActiveDeployment#spring_cloud_app_id}.
	SpringCloudAppId *string `field:"required" json:"springCloudAppId" yaml:"springCloudAppId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/spring_cloud_active_deployment#id SpringCloudActiveDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/spring_cloud_active_deployment#timeouts SpringCloudActiveDeployment#timeouts}
	Timeouts *SpringCloudActiveDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

