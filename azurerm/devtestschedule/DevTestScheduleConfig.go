// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevTestScheduleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#lab_name DevTestSchedule#lab_name}.
	LabName *string `field:"required" json:"labName" yaml:"labName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#location DevTestSchedule#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#name DevTestSchedule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#notification_settings DevTestSchedule#notification_settings}
	NotificationSettings *DevTestScheduleNotificationSettings `field:"required" json:"notificationSettings" yaml:"notificationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#resource_group_name DevTestSchedule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#task_type DevTestSchedule#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#time_zone_id DevTestSchedule#time_zone_id}.
	TimeZoneId *string `field:"required" json:"timeZoneId" yaml:"timeZoneId"`
	// daily_recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#daily_recurrence DevTestSchedule#daily_recurrence}
	DailyRecurrence *DevTestScheduleDailyRecurrence `field:"optional" json:"dailyRecurrence" yaml:"dailyRecurrence"`
	// hourly_recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#hourly_recurrence DevTestSchedule#hourly_recurrence}
	HourlyRecurrence *DevTestScheduleHourlyRecurrence `field:"optional" json:"hourlyRecurrence" yaml:"hourlyRecurrence"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#id DevTestSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#status DevTestSchedule#status}.
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#tags DevTestSchedule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#timeouts DevTestSchedule#timeouts}
	Timeouts *DevTestScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// weekly_recurrence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.84.0/docs/resources/dev_test_schedule#weekly_recurrence DevTestSchedule#weekly_recurrence}
	WeeklyRecurrence *DevTestScheduleWeeklyRecurrence `field:"optional" json:"weeklyRecurrence" yaml:"weeklyRecurrence"`
}

