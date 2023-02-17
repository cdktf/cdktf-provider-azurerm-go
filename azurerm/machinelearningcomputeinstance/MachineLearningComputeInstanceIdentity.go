package machinelearningcomputeinstance


type MachineLearningComputeInstanceIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_instance#type MachineLearningComputeInstance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_instance#identity_ids MachineLearningComputeInstance#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

