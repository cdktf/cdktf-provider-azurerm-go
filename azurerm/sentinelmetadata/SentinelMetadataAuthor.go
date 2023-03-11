package sentinelmetadata


type SentinelMetadataAuthor struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_metadata#email SentinelMetadata#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_metadata#link SentinelMetadata#link}.
	Link *string `field:"optional" json:"link" yaml:"link"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_metadata#name SentinelMetadata#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

