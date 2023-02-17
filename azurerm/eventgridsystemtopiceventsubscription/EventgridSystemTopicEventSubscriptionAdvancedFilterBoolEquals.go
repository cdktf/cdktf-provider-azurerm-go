package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEquals struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#key EventgridSystemTopicEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#value EventgridSystemTopicEventSubscription#value}.
	Value interface{} `field:"required" json:"value" yaml:"value"`
}

