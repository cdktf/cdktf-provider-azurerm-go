package machinelearningcomputecluster


type MachineLearningComputeClusterSsh struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_cluster#admin_username MachineLearningComputeCluster#admin_username}.
	AdminUsername *string `field:"required" json:"adminUsername" yaml:"adminUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_cluster#admin_password MachineLearningComputeCluster#admin_password}.
	AdminPassword *string `field:"optional" json:"adminPassword" yaml:"adminPassword"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/machine_learning_compute_cluster#key_value MachineLearningComputeCluster#key_value}.
	KeyValue *string `field:"optional" json:"keyValue" yaml:"keyValue"`
}
