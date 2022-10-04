package dnsmxrecord


type DnsMxRecordRecord struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_mx_record#exchange DnsMxRecord#exchange}.
	Exchange *string `field:"required" json:"exchange" yaml:"exchange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_mx_record#preference DnsMxRecord#preference}.
	Preference *string `field:"required" json:"preference" yaml:"preference"`
}

