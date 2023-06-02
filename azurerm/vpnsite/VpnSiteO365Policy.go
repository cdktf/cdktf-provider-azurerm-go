package vpnsite


type VpnSiteO365Policy struct {
	// traffic_category block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.59.0/docs/resources/vpn_site#traffic_category VpnSite#traffic_category}
	TrafficCategory *VpnSiteO365PolicyTrafficCategory `field:"optional" json:"trafficCategory" yaml:"trafficCategory"`
}

