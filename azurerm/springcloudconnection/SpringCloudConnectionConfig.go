// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudConnectionConfig struct {
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
	// authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#authentication SpringCloudConnection#authentication}
	Authentication *SpringCloudConnectionAuthentication `field:"required" json:"authentication" yaml:"authentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#name SpringCloudConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#spring_cloud_id SpringCloudConnection#spring_cloud_id}.
	SpringCloudId *string `field:"required" json:"springCloudId" yaml:"springCloudId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#target_resource_id SpringCloudConnection#target_resource_id}.
	TargetResourceId *string `field:"required" json:"targetResourceId" yaml:"targetResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#client_type SpringCloudConnection#client_type}.
	ClientType *string `field:"optional" json:"clientType" yaml:"clientType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#id SpringCloudConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// secret_store block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#secret_store SpringCloudConnection#secret_store}
	SecretStore *SpringCloudConnectionSecretStore `field:"optional" json:"secretStore" yaml:"secretStore"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#timeouts SpringCloudConnection#timeouts}
	Timeouts *SpringCloudConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/spring_cloud_connection#vnet_solution SpringCloudConnection#vnet_solution}.
	VnetSolution *string `field:"optional" json:"vnetSolution" yaml:"vnetSolution"`
}

