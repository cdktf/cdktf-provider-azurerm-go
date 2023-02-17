package machinelearningcomputecluster


type MachineLearningComputeClusterIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_cluster#type MachineLearningComputeCluster#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_cluster#identity_ids MachineLearningComputeCluster#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}

