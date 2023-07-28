package synapsesparkpool


type SynapseSparkPoolSparkConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/synapse_spark_pool#content SynapseSparkPool#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/synapse_spark_pool#filename SynapseSparkPool#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}

