package proximityplacementgroup


type ProximityPlacementGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/proximity_placement_group#create ProximityPlacementGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/proximity_placement_group#delete ProximityPlacementGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/proximity_placement_group#read ProximityPlacementGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/proximity_placement_group#update ProximityPlacementGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
