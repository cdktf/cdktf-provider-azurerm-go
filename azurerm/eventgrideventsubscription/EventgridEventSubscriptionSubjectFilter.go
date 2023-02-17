package eventgrideventsubscription


type EventgridEventSubscriptionSubjectFilter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#case_sensitive EventgridEventSubscription#case_sensitive}.
	CaseSensitive interface{} `field:"optional" json:"caseSensitive" yaml:"caseSensitive"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#subject_begins_with EventgridEventSubscription#subject_begins_with}.
	SubjectBeginsWith *string `field:"optional" json:"subjectBeginsWith" yaml:"subjectBeginsWith"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#subject_ends_with EventgridEventSubscription#subject_ends_with}.
	SubjectEndsWith *string `field:"optional" json:"subjectEndsWith" yaml:"subjectEndsWith"`
}

