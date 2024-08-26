// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devtestglobalvmshutdownschedule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DevTestGlobalVmShutdownScheduleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#daily_recurrence_time DevTestGlobalVmShutdownSchedule#daily_recurrence_time}.
	DailyRecurrenceTime *string `field:"required" json:"dailyRecurrenceTime" yaml:"dailyRecurrenceTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#location DevTestGlobalVmShutdownSchedule#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// notification_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#notification_settings DevTestGlobalVmShutdownSchedule#notification_settings}
	NotificationSettings *DevTestGlobalVmShutdownScheduleNotificationSettings `field:"required" json:"notificationSettings" yaml:"notificationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#timezone DevTestGlobalVmShutdownSchedule#timezone}.
	Timezone *string `field:"required" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#virtual_machine_id DevTestGlobalVmShutdownSchedule#virtual_machine_id}.
	VirtualMachineId *string `field:"required" json:"virtualMachineId" yaml:"virtualMachineId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#enabled DevTestGlobalVmShutdownSchedule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#id DevTestGlobalVmShutdownSchedule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#tags DevTestGlobalVmShutdownSchedule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_test_global_vm_shutdown_schedule#timeouts DevTestGlobalVmShutdownSchedule#timeouts}
	Timeouts *DevTestGlobalVmShutdownScheduleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

