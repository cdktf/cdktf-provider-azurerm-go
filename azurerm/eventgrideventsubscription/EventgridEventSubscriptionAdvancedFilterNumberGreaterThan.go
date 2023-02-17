package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilterNumberGreaterThan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

