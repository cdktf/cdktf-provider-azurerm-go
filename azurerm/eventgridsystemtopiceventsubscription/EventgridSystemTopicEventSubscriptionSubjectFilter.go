// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionSubjectFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/eventgrid_system_topic_event_subscription#case_sensitive EventgridSystemTopicEventSubscription#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/eventgrid_system_topic_event_subscription#subject_begins_with EventgridSystemTopicEventSubscription#subject_begins_with}.
	SubjectBeginsWith *string `field:"optional" json:"subjectBeginsWith" yaml:"subjectBeginsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.76.0/docs/resources/eventgrid_system_topic_event_subscription#subject_ends_with EventgridSystemTopicEventSubscription#subject_ends_with}.
	SubjectEndsWith *string `field:"optional" json:"subjectEndsWith" yaml:"subjectEndsWith"`
}

