package vpnsite


type VpnSiteO365Policy struct {
	// traffic_category block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_site#traffic_category VpnSite#traffic_category}
	TrafficCategory *VpnSiteO365PolicyTrafficCategory `field:"optional" json:"trafficCategory" yaml:"trafficCategory"`
}
