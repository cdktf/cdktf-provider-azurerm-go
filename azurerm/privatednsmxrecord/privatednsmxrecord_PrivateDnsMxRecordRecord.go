package privatednsmxrecord


type PrivateDnsMxRecordRecord struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_mx_record#exchange PrivateDnsMxRecord#exchange}.
	Exchange *string `field:"required" json:"exchange" yaml:"exchange"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/private_dns_mx_record#preference PrivateDnsMxRecord#preference}.
	Preference *float64 `field:"required" json:"preference" yaml:"preference"`
}

