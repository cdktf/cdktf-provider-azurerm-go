// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package notificationhub


type NotificationHubGcmCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/notification_hub#api_key NotificationHub#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
}

