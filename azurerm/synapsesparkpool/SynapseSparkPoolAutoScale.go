package synapsesparkpool


type SynapseSparkPoolAutoScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_spark_pool#max_node_count SynapseSparkPool#max_node_count}.
	MaxNodeCount *float64 `field:"required" json:"maxNodeCount" yaml:"maxNodeCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/synapse_spark_pool#min_node_count SynapseSparkPool#min_node_count}.
	MinNodeCount *float64 `field:"required" json:"minNodeCount" yaml:"minNodeCount"`
}
