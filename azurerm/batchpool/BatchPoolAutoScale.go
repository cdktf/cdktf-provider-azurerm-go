package batchpool


type BatchPoolAutoScale struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#formula BatchPool#formula}.
	Formula *string `field:"required" json:"formula" yaml:"formula"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#evaluation_interval BatchPool#evaluation_interval}.
	EvaluationInterval *string `field:"optional" json:"evaluationInterval" yaml:"evaluationInterval"`
}

