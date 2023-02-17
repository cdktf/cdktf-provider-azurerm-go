package capacityreservation


type CapacityReservationSku struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/capacity_reservation#capacity CapacityReservation#capacity}.
	Capacity *float64 `field:"required" json:"capacity" yaml:"capacity"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/capacity_reservation#name CapacityReservation#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

