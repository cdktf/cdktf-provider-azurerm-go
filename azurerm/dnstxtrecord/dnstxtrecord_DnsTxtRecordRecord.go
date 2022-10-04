package dnstxtrecord


type DnsTxtRecordRecord struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dns_txt_record#value DnsTxtRecord#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

