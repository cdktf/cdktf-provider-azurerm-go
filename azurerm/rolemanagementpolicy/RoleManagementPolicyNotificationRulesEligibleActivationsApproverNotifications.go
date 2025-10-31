// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyNotificationRulesEligibleActivationsApproverNotifications struct {
	// Whether the default recipients are notified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/role_management_policy#default_recipients RoleManagementPolicy#default_recipients}
	DefaultRecipients interface{} `field:"required" json:"defaultRecipients" yaml:"defaultRecipients"`
	// What level of notifications are sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/role_management_policy#notification_level RoleManagementPolicy#notification_level}
	NotificationLevel *string `field:"required" json:"notificationLevel" yaml:"notificationLevel"`
	// The additional recipients to notify.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.51.0/docs/resources/role_management_policy#additional_recipients RoleManagementPolicy#additional_recipients}
	AdditionalRecipients *[]*string `field:"optional" json:"additionalRecipients" yaml:"additionalRecipients"`
}

