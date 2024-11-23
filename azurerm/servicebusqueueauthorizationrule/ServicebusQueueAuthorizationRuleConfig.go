// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusqueueauthorizationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusQueueAuthorizationRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#name ServicebusQueueAuthorizationRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#queue_id ServicebusQueueAuthorizationRule#queue_id}.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#id ServicebusQueueAuthorizationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#listen ServicebusQueueAuthorizationRule#listen}.
	Listen interface{} `field:"optional" json:"listen" yaml:"listen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#manage ServicebusQueueAuthorizationRule#manage}.
	Manage interface{} `field:"optional" json:"manage" yaml:"manage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#send ServicebusQueueAuthorizationRule#send}.
	Send interface{} `field:"optional" json:"send" yaml:"send"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/servicebus_queue_authorization_rule#timeouts ServicebusQueueAuthorizationRule#timeouts}
	Timeouts *ServicebusQueueAuthorizationRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

