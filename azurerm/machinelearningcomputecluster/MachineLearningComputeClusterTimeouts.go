package machinelearningcomputecluster


type MachineLearningComputeClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/machine_learning_compute_cluster#create MachineLearningComputeCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/machine_learning_compute_cluster#delete MachineLearningComputeCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.66.0/docs/resources/machine_learning_compute_cluster#read MachineLearningComputeCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

