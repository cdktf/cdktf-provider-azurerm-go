package eventgriddomain


type EventgridDomainInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_domain#action EventgridDomain#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/eventgrid_domain#ip_mask EventgridDomain#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}

