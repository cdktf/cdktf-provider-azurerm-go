package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilterStringNotContains struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#values EventgridEventSubscription#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

