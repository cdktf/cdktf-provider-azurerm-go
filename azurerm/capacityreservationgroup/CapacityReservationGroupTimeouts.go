package capacityreservationgroup


type CapacityReservationGroupTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/capacity_reservation_group#create CapacityReservationGroup#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/capacity_reservation_group#delete CapacityReservationGroup#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/capacity_reservation_group#read CapacityReservationGroup#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/capacity_reservation_group#update CapacityReservationGroup#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}
