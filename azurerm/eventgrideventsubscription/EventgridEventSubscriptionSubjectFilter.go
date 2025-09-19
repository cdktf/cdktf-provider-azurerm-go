// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription


type EventgridEventSubscriptionSubjectFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/eventgrid_event_subscription#case_sensitive EventgridEventSubscription#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/eventgrid_event_subscription#subject_begins_with EventgridEventSubscription#subject_begins_with}.
	SubjectBeginsWith *string `field:"optional" json:"subjectBeginsWith" yaml:"subjectBeginsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/eventgrid_event_subscription#subject_ends_with EventgridEventSubscription#subject_ends_with}.
	SubjectEndsWith *string `field:"optional" json:"subjectEndsWith" yaml:"subjectEndsWith"`
}

