// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package streamanalyticsjobschedule


type StreamAnalyticsJobScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/stream_analytics_job_schedule#create StreamAnalyticsJobSchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/stream_analytics_job_schedule#delete StreamAnalyticsJobSchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/stream_analytics_job_schedule#read StreamAnalyticsJobSchedule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/stream_analytics_job_schedule#update StreamAnalyticsJobSchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

