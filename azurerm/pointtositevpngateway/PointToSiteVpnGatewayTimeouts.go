package pointtositevpngateway


type PointToSiteVpnGatewayTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#create PointToSiteVpnGateway#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#delete PointToSiteVpnGateway#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#read PointToSiteVpnGateway#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/point_to_site_vpn_gateway#update PointToSiteVpnGateway#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

