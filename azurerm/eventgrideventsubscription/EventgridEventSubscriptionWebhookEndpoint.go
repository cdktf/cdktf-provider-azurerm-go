package eventgrideventsubscription


type EventgridEventSubscriptionWebhookEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventgrid_event_subscription#url EventgridEventSubscription#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventgrid_event_subscription#active_directory_app_id_or_uri EventgridEventSubscription#active_directory_app_id_or_uri}.
	ActiveDirectoryAppIdOrUri *string `field:"optional" json:"activeDirectoryAppIdOrUri" yaml:"activeDirectoryAppIdOrUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventgrid_event_subscription#active_directory_tenant_id EventgridEventSubscription#active_directory_tenant_id}.
	ActiveDirectoryTenantId *string `field:"optional" json:"activeDirectoryTenantId" yaml:"activeDirectoryTenantId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventgrid_event_subscription#max_events_per_batch EventgridEventSubscription#max_events_per_batch}.
	MaxEventsPerBatch *float64 `field:"optional" json:"maxEventsPerBatch" yaml:"maxEventsPerBatch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/resources/eventgrid_event_subscription#preferred_batch_size_in_kilobytes EventgridEventSubscription#preferred_batch_size_in_kilobytes}.
	PreferredBatchSizeInKilobytes *float64 `field:"optional" json:"preferredBatchSizeInKilobytes" yaml:"preferredBatchSizeInKilobytes"`
}

