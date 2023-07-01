package webpubsubhub


type WebPubsubHubEventHandler struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/web_pubsub_hub#url_template WebPubsubHub#url_template}.
	UrlTemplate *string `field:"required" json:"urlTemplate" yaml:"urlTemplate"`
	// auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/web_pubsub_hub#auth WebPubsubHub#auth}
	Auth *WebPubsubHubEventHandlerAuth `field:"optional" json:"auth" yaml:"auth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/web_pubsub_hub#system_events WebPubsubHub#system_events}.
	SystemEvents *[]*string `field:"optional" json:"systemEvents" yaml:"systemEvents"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/web_pubsub_hub#user_event_pattern WebPubsubHub#user_event_pattern}.
	UserEventPattern *string `field:"optional" json:"userEventPattern" yaml:"userEventPattern"`
}

