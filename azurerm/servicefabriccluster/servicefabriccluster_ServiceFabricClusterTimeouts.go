package servicefabriccluster


type ServiceFabricClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#create ServiceFabricCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#delete ServiceFabricCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#read ServiceFabricCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/service_fabric_cluster#update ServiceFabricCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
