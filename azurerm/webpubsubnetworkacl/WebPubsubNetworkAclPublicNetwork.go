package webpubsubnetworkacl


type WebPubsubNetworkAclPublicNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/web_pubsub_network_acl#allowed_request_types WebPubsubNetworkAcl#allowed_request_types}.
	AllowedRequestTypes *[]*string `field:"optional" json:"allowedRequestTypes" yaml:"allowedRequestTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/web_pubsub_network_acl#denied_request_types WebPubsubNetworkAcl#denied_request_types}.
	DeniedRequestTypes *[]*string `field:"optional" json:"deniedRequestTypes" yaml:"deniedRequestTypes"`
}

