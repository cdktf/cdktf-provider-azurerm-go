// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationhub


type NotificationHubBrowserCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/notification_hub#subject NotificationHub#subject}.
	Subject *string `field:"required" json:"subject" yaml:"subject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/notification_hub#vapid_private_key NotificationHub#vapid_private_key}.
	VapidPrivateKey *string `field:"required" json:"vapidPrivateKey" yaml:"vapidPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.37.0/docs/resources/notification_hub#vapid_public_key NotificationHub#vapid_public_key}.
	VapidPublicKey *string `field:"required" json:"vapidPublicKey" yaml:"vapidPublicKey"`
}

