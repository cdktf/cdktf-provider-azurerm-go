package eventgridsystemtopiceventsubscription


type EventgridSystemTopicEventSubscriptionWebhookEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_system_topic_event_subscription#url EventgridSystemTopicEventSubscription#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_system_topic_event_subscription#active_directory_app_id_or_uri EventgridSystemTopicEventSubscription#active_directory_app_id_or_uri}.
	ActiveDirectoryAppIdOrUri *string `field:"optional" json:"activeDirectoryAppIdOrUri" yaml:"activeDirectoryAppIdOrUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_system_topic_event_subscription#active_directory_tenant_id EventgridSystemTopicEventSubscription#active_directory_tenant_id}.
	ActiveDirectoryTenantId *string `field:"optional" json:"activeDirectoryTenantId" yaml:"activeDirectoryTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_system_topic_event_subscription#max_events_per_batch EventgridSystemTopicEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/eventgrid_system_topic_event_subscription#preferred_batch_size_in_kilobytes EventgridSystemTopicEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

