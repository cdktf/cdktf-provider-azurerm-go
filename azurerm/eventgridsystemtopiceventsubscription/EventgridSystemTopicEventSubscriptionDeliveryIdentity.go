package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionDeliveryIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#type EventgridSystemTopicEventSubscription#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_system_topic_event_subscription#user_assigned_identity EventgridSystemTopicEventSubscription#user_assigned_identity}.
	UserAssignedIdentity *string `field:"optional" json:"userAssignedIdentity" yaml:"userAssignedIdentity"`
}

