package eventgriddomain


type EventgridDomainInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/eventgrid_domain#action EventgridDomain#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/eventgrid_domain#ip_mask EventgridDomain#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}

