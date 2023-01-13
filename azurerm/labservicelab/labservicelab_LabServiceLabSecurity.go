package labservicelab


type LabServiceLabSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#open_access_enabled LabServiceLab#open_access_enabled}.
	OpenAccessEnabled interface{} `field:"required" json:"openAccessEnabled" yaml:"openAccessEnabled"`
}

