package webpubsubhub


type WebPubsubHubEventListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/web_pubsub_hub#eventhub_name WebPubsubHub#eventhub_name}.
	EventhubName *string `field:"required" json:"eventhubName" yaml:"eventhubName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/web_pubsub_hub#eventhub_namespace_name WebPubsubHub#eventhub_namespace_name}.
	EventhubNamespaceName *string `field:"required" json:"eventhubNamespaceName" yaml:"eventhubNamespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/web_pubsub_hub#system_event_name_filter WebPubsubHub#system_event_name_filter}.
	SystemEventNameFilter *[]*string `field:"optional" json:"systemEventNameFilter" yaml:"systemEventNameFilter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/web_pubsub_hub#user_event_name_filter WebPubsubHub#user_event_name_filter}.
	UserEventNameFilter *[]*string `field:"optional" json:"userEventNameFilter" yaml:"userEventNameFilter"`
}

