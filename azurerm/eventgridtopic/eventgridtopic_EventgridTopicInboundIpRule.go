package eventgridtopic


type EventgridTopicInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_topic#action EventgridTopic#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_topic#ip_mask EventgridTopic#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}
