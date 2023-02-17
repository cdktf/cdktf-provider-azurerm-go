package synapsesparkpool


type SynapseSparkPoolAutoPause struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_spark_pool#delay_in_minutes SynapseSparkPool#delay_in_minutes}.
	DelayInMinutes *float64 `field:"required" json:"delayInMinutes" yaml:"delayInMinutes"`
}

