package webpubsubnetworkacl


type WebPubsubNetworkAclTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_network_acl#create WebPubsubNetworkAcl#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_network_acl#delete WebPubsubNetworkAcl#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_network_acl#read WebPubsubNetworkAcl#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_pubsub_network_acl#update WebPubsubNetworkAcl#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
