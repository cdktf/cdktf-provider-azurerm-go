package dnszone


type DnsZoneSoaRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#email DnsZone#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#host_name DnsZone#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#expire_time DnsZone#expire_time}.
	ExpireTime *float64 `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#minimum_ttl DnsZone#minimum_ttl}.
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#refresh_time DnsZone#refresh_time}.
	RefreshTime *float64 `field:"optional" json:"refreshTime" yaml:"refreshTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#retry_time DnsZone#retry_time}.
	RetryTime *float64 `field:"optional" json:"retryTime" yaml:"retryTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#serial_number DnsZone#serial_number}.
	SerialNumber *float64 `field:"optional" json:"serialNumber" yaml:"serialNumber"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#tags DnsZone#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/dns_zone#ttl DnsZone#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}

