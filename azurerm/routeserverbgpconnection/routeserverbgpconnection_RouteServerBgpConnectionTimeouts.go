package routeserverbgpconnection


type RouteServerBgpConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_server_bgp_connection#create RouteServerBgpConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_server_bgp_connection#delete RouteServerBgpConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/route_server_bgp_connection#read RouteServerBgpConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}
