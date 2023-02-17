package frontdoor


type FrontdoorBackendPoolSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enforce_backend_pools_certificate_name_check Frontdoor#enforce_backend_pools_certificate_name_check}.
	EnforceBackendPoolsCertificateNameCheck interface{} `field:"required" json:"enforceBackendPoolsCertificateNameCheck" yaml:"enforceBackendPoolsCertificateNameCheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#backend_pools_send_receive_timeout_seconds Frontdoor#backend_pools_send_receive_timeout_seconds}.
	BackendPoolsSendReceiveTimeoutSeconds *float64 `field:"optional" json:"backendPoolsSendReceiveTimeoutSeconds" yaml:"backendPoolsSendReceiveTimeoutSeconds"`
}

