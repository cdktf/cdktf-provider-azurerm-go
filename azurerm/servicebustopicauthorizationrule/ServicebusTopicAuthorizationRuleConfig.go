// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebustopicauthorizationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusTopicAuthorizationRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#name ServicebusTopicAuthorizationRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#topic_id ServicebusTopicAuthorizationRule#topic_id}.
	TopicId *string `field:"required" json:"topicId" yaml:"topicId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#id ServicebusTopicAuthorizationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#listen ServicebusTopicAuthorizationRule#listen}.
	Listen interface{} `field:"optional" json:"listen" yaml:"listen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#manage ServicebusTopicAuthorizationRule#manage}.
	Manage interface{} `field:"optional" json:"manage" yaml:"manage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#send ServicebusTopicAuthorizationRule#send}.
	Send interface{} `field:"optional" json:"send" yaml:"send"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/servicebus_topic_authorization_rule#timeouts ServicebusTopicAuthorizationRule#timeouts}
	Timeouts *ServicebusTopicAuthorizationRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

