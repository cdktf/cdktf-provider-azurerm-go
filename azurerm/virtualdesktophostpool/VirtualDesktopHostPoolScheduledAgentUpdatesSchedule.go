package virtualdesktophostpool


type VirtualDesktopHostPoolScheduledAgentUpdatesSchedule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#day_of_week VirtualDesktopHostPool#day_of_week}.
	DayOfWeek *string `field:"required" json:"dayOfWeek" yaml:"dayOfWeek"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_desktop_host_pool#hour_of_day VirtualDesktopHostPool#hour_of_day}.
	HourOfDay *float64 `field:"required" json:"hourOfDay" yaml:"hourOfDay"`
}
