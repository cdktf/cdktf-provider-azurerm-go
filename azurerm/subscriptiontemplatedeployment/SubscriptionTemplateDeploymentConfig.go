// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package subscriptiontemplatedeployment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SubscriptionTemplateDeploymentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#location SubscriptionTemplateDeployment#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#name SubscriptionTemplateDeployment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#debug_level SubscriptionTemplateDeployment#debug_level}.
	DebugLevel *string `field:"optional" json:"debugLevel" yaml:"debugLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#id SubscriptionTemplateDeployment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#parameters_content SubscriptionTemplateDeployment#parameters_content}.
	ParametersContent *string `field:"optional" json:"parametersContent" yaml:"parametersContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#tags SubscriptionTemplateDeployment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#template_content SubscriptionTemplateDeployment#template_content}.
	TemplateContent *string `field:"optional" json:"templateContent" yaml:"templateContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#template_spec_version_id SubscriptionTemplateDeployment#template_spec_version_id}.
	TemplateSpecVersionId *string `field:"optional" json:"templateSpecVersionId" yaml:"templateSpecVersionId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/subscription_template_deployment#timeouts SubscriptionTemplateDeployment#timeouts}
	Timeouts *SubscriptionTemplateDeploymentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

