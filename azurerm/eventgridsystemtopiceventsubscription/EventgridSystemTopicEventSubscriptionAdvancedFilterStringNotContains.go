package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#values EventgridSystemTopicEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

