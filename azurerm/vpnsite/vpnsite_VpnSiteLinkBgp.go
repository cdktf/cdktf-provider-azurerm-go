package vpnsite


type VpnSiteLinkBgp struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_site#asn VpnSite#asn}.
	Asn *float64 `field:"required" json:"asn" yaml:"asn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_site#peering_address VpnSite#peering_address}.
	PeeringAddress *string `field:"required" json:"peeringAddress" yaml:"peeringAddress"`
}
