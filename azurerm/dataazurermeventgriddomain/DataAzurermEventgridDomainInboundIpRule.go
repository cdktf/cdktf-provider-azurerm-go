package dataazurermeventgriddomain


type DataAzurermEventgridDomainInboundIpRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/data-sources/eventgrid_domain#action DataAzurermEventgridDomain#action}.
	Action *string `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs/data-sources/eventgrid_domain#ip_mask DataAzurermEventgridDomain#ip_mask}.
	IpMask *string `field:"optional" json:"ipMask" yaml:"ipMask"`
}

