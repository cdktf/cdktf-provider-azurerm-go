package firewall


type FirewallVirtualHub struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall#virtual_hub_id Firewall#virtual_hub_id}.
	VirtualHubId *string `field:"required" json:"virtualHubId" yaml:"virtualHubId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall#public_ip_count Firewall#public_ip_count}.
	PublicIpCount *float64 `field:"optional" json:"publicIpCount" yaml:"publicIpCount"`
}

