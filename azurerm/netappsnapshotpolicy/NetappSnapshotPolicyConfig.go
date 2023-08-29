// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappsnapshotpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappSnapshotPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#account_name NetappSnapshotPolicy#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#enabled NetappSnapshotPolicy#enabled}.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#location NetappSnapshotPolicy#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#name NetappSnapshotPolicy#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#resource_group_name NetappSnapshotPolicy#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// daily_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#daily_schedule NetappSnapshotPolicy#daily_schedule}
	DailySchedule *NetappSnapshotPolicyDailySchedule `field:"optional" json:"dailySchedule" yaml:"dailySchedule"`
	// hourly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#hourly_schedule NetappSnapshotPolicy#hourly_schedule}
	HourlySchedule *NetappSnapshotPolicyHourlySchedule `field:"optional" json:"hourlySchedule" yaml:"hourlySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#id NetappSnapshotPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// monthly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#monthly_schedule NetappSnapshotPolicy#monthly_schedule}
	MonthlySchedule *NetappSnapshotPolicyMonthlySchedule `field:"optional" json:"monthlySchedule" yaml:"monthlySchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#tags NetappSnapshotPolicy#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#timeouts NetappSnapshotPolicy#timeouts}
	Timeouts *NetappSnapshotPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// weekly_schedule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/netapp_snapshot_policy#weekly_schedule NetappSnapshotPolicy#weekly_schedule}
	WeeklySchedule *NetappSnapshotPolicyWeeklySchedule `field:"optional" json:"weeklySchedule" yaml:"weeklySchedule"`
}

