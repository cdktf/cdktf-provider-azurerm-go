package machinelearningcomputeinstance


type MachineLearningComputeInstanceSsh struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_instance#public_key MachineLearningComputeInstance#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}

