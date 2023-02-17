package vpngatewayconnection


type VpnGatewayConnectionVpnLink struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#name VpnGatewayConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#vpn_site_link_id VpnGatewayConnection#vpn_site_link_id}.
	VpnSiteLinkId *string `field:"required" json:"vpnSiteLinkId" yaml:"vpnSiteLinkId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#bandwidth_mbps VpnGatewayConnection#bandwidth_mbps}.
	BandwidthMbps *float64 `field:"optional" json:"bandwidthMbps" yaml:"bandwidthMbps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#bgp_enabled VpnGatewayConnection#bgp_enabled}.
	BgpEnabled interface{} `field:"optional" json:"bgpEnabled" yaml:"bgpEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#connection_mode VpnGatewayConnection#connection_mode}.
	ConnectionMode *string `field:"optional" json:"connectionMode" yaml:"connectionMode"`
	// custom_bgp_address block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#custom_bgp_address VpnGatewayConnection#custom_bgp_address}
	CustomBgpAddress interface{} `field:"optional" json:"customBgpAddress" yaml:"customBgpAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#egress_nat_rule_ids VpnGatewayConnection#egress_nat_rule_ids}.
	EgressNatRuleIds *[]*string `field:"optional" json:"egressNatRuleIds" yaml:"egressNatRuleIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ingress_nat_rule_ids VpnGatewayConnection#ingress_nat_rule_ids}.
	IngressNatRuleIds *[]*string `field:"optional" json:"ingressNatRuleIds" yaml:"ingressNatRuleIds"`
	// ipsec_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ipsec_policy VpnGatewayConnection#ipsec_policy}
	IpsecPolicy interface{} `field:"optional" json:"ipsecPolicy" yaml:"ipsecPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#local_azure_ip_address_enabled VpnGatewayConnection#local_azure_ip_address_enabled}.
	LocalAzureIpAddressEnabled interface{} `field:"optional" json:"localAzureIpAddressEnabled" yaml:"localAzureIpAddressEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#policy_based_traffic_selector_enabled VpnGatewayConnection#policy_based_traffic_selector_enabled}.
	PolicyBasedTrafficSelectorEnabled interface{} `field:"optional" json:"policyBasedTrafficSelectorEnabled" yaml:"policyBasedTrafficSelectorEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#protocol VpnGatewayConnection#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#ratelimit_enabled VpnGatewayConnection#ratelimit_enabled}.
	RatelimitEnabled interface{} `field:"optional" json:"ratelimitEnabled" yaml:"ratelimitEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#route_weight VpnGatewayConnection#route_weight}.
	RouteWeight *float64 `field:"optional" json:"routeWeight" yaml:"routeWeight"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_gateway_connection#shared_key VpnGatewayConnection#shared_key}.
	SharedKey *string `field:"optional" json:"sharedKey" yaml:"sharedKey"`
}

