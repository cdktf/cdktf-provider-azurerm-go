package eventgrideventsubscription


type EventgridEventSubscriptionAdvancedFilterNumberLessThan struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_event_subscription#key EventgridEventSubscription#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_event_subscription#value EventgridEventSubscription#value}.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

