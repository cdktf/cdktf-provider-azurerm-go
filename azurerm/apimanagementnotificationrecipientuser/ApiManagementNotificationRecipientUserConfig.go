// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementnotificationrecipientuser

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementNotificationRecipientUserConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/api_management_notification_recipient_user#api_management_id ApiManagementNotificationRecipientUser#api_management_id}.
	ApiManagementId *string `field:"required" json:"apiManagementId" yaml:"apiManagementId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/api_management_notification_recipient_user#notification_type ApiManagementNotificationRecipientUser#notification_type}.
	NotificationType *string `field:"required" json:"notificationType" yaml:"notificationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/api_management_notification_recipient_user#user_id ApiManagementNotificationRecipientUser#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/api_management_notification_recipient_user#id ApiManagementNotificationRecipientUser#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/api_management_notification_recipient_user#timeouts ApiManagementNotificationRecipientUser#timeouts}
	Timeouts *ApiManagementNotificationRecipientUserTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

