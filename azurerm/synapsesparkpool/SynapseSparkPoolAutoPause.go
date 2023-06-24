package synapsesparkpool


type SynapseSparkPoolAutoPause struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/synapse_spark_pool#delay_in_minutes SynapseSparkPool#delay_in_minutes}.
	DelayInMinutes *float64 `field:"required" json:"delayInMinutes" yaml:"delayInMinutes"`
}

