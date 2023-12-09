// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebussubscriptionrule


type ServicebusSubscriptionRuleCorrelationFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#content_type ServicebusSubscriptionRule#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#correlation_id ServicebusSubscriptionRule#correlation_id}.
	CorrelationId *string `field:"optional" json:"correlationId" yaml:"correlationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#label ServicebusSubscriptionRule#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#message_id ServicebusSubscriptionRule#message_id}.
	MessageId *string `field:"optional" json:"messageId" yaml:"messageId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#properties ServicebusSubscriptionRule#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#reply_to ServicebusSubscriptionRule#reply_to}.
	ReplyTo *string `field:"optional" json:"replyTo" yaml:"replyTo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#reply_to_session_id ServicebusSubscriptionRule#reply_to_session_id}.
	ReplyToSessionId *string `field:"optional" json:"replyToSessionId" yaml:"replyToSessionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#session_id ServicebusSubscriptionRule#session_id}.
	SessionId *string `field:"optional" json:"sessionId" yaml:"sessionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/servicebus_subscription_rule#to ServicebusSubscriptionRule#to}.
	To *string `field:"optional" json:"to" yaml:"to"`
}

