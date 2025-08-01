// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicebussubscription


type ServicebusSubscriptionClientScopedSubscription struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/servicebus_subscription#client_id ServicebusSubscription#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/servicebus_subscription#is_client_scoped_subscription_shareable ServicebusSubscription#is_client_scoped_subscription_shareable}.
	IsClientScopedSubscriptionShareable interface{} `field:"optional" json:"isClientScopedSubscriptionShareable" yaml:"isClientScopedSubscriptionShareable"`
}

