package webpubsubhub


type WebPubsubHubEventHandlerAuth struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_hub#managed_identity_id WebPubsubHub#managed_identity_id}.
	ManagedIdentityId *string `field:"required" json:"managedIdentityId" yaml:"managedIdentityId"`
}

