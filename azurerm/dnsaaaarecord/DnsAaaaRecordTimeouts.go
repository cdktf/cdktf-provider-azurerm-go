package dnsaaaarecord


type DnsAaaaRecordTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_aaaa_record#create DnsAaaaRecord#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_aaaa_record#delete DnsAaaaRecord#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_aaaa_record#read DnsAaaaRecord#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_aaaa_record#update DnsAaaaRecord#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

