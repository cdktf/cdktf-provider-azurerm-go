package kubernetesfleetmanager


type KubernetesFleetManagerTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_fleet_manager#create KubernetesFleetManager#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_fleet_manager#delete KubernetesFleetManager#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_fleet_manager#read KubernetesFleetManager#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/kubernetes_fleet_manager#update KubernetesFleetManager#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
