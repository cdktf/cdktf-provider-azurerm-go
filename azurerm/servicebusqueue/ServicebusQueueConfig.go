// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebusqueue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServicebusQueueConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#name ServicebusQueue#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#namespace_id ServicebusQueue#namespace_id}.
	NamespaceId *string `field:"required" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#auto_delete_on_idle ServicebusQueue#auto_delete_on_idle}.
	AutoDeleteOnIdle *string `field:"optional" json:"autoDeleteOnIdle" yaml:"autoDeleteOnIdle"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#dead_lettering_on_message_expiration ServicebusQueue#dead_lettering_on_message_expiration}.
	DeadLetteringOnMessageExpiration interface{} `field:"optional" json:"deadLetteringOnMessageExpiration" yaml:"deadLetteringOnMessageExpiration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#default_message_ttl ServicebusQueue#default_message_ttl}.
	DefaultMessageTtl *string `field:"optional" json:"defaultMessageTtl" yaml:"defaultMessageTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#duplicate_detection_history_time_window ServicebusQueue#duplicate_detection_history_time_window}.
	DuplicateDetectionHistoryTimeWindow *string `field:"optional" json:"duplicateDetectionHistoryTimeWindow" yaml:"duplicateDetectionHistoryTimeWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#enable_batched_operations ServicebusQueue#enable_batched_operations}.
	EnableBatchedOperations interface{} `field:"optional" json:"enableBatchedOperations" yaml:"enableBatchedOperations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#enable_express ServicebusQueue#enable_express}.
	EnableExpress interface{} `field:"optional" json:"enableExpress" yaml:"enableExpress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#enable_partitioning ServicebusQueue#enable_partitioning}.
	EnablePartitioning interface{} `field:"optional" json:"enablePartitioning" yaml:"enablePartitioning"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#forward_dead_lettered_messages_to ServicebusQueue#forward_dead_lettered_messages_to}.
	ForwardDeadLetteredMessagesTo *string `field:"optional" json:"forwardDeadLetteredMessagesTo" yaml:"forwardDeadLetteredMessagesTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#forward_to ServicebusQueue#forward_to}.
	ForwardTo *string `field:"optional" json:"forwardTo" yaml:"forwardTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#id ServicebusQueue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#lock_duration ServicebusQueue#lock_duration}.
	LockDuration *string `field:"optional" json:"lockDuration" yaml:"lockDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#max_delivery_count ServicebusQueue#max_delivery_count}.
	MaxDeliveryCount *float64 `field:"optional" json:"maxDeliveryCount" yaml:"maxDeliveryCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#max_message_size_in_kilobytes ServicebusQueue#max_message_size_in_kilobytes}.
	MaxMessageSizeInKilobytes *float64 `field:"optional" json:"maxMessageSizeInKilobytes" yaml:"maxMessageSizeInKilobytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#max_size_in_megabytes ServicebusQueue#max_size_in_megabytes}.
	MaxSizeInMegabytes *float64 `field:"optional" json:"maxSizeInMegabytes" yaml:"maxSizeInMegabytes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#requires_duplicate_detection ServicebusQueue#requires_duplicate_detection}.
	RequiresDuplicateDetection interface{} `field:"optional" json:"requiresDuplicateDetection" yaml:"requiresDuplicateDetection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#requires_session ServicebusQueue#requires_session}.
	RequiresSession interface{} `field:"optional" json:"requiresSession" yaml:"requiresSession"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#status ServicebusQueue#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/servicebus_queue#timeouts ServicebusQueue#timeouts}
	Timeouts *ServicebusQueueTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

