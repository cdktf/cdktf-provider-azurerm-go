package dataazurermeventgriddomain


type DataAzurermEventgridDomainInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/eventgrid_domain#action DataAzurermEventgridDomain#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/eventgrid_domain#ip_mask DataAzurermEventgridDomain#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}

