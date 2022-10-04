package webapplicationfirewallpolicy


type WebApplicationFirewallPolicyPolicySettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#enabled WebApplicationFirewallPolicy#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#file_upload_limit_in_mb WebApplicationFirewallPolicy#file_upload_limit_in_mb}.
	FileUploadLimitInMb *float64 `field:"optional" json:"fileUploadLimitInMb" yaml:"fileUploadLimitInMb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#max_request_body_size_in_kb WebApplicationFirewallPolicy#max_request_body_size_in_kb}.
	MaxRequestBodySizeInKb *float64 `field:"optional" json:"maxRequestBodySizeInKb" yaml:"maxRequestBodySizeInKb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#mode WebApplicationFirewallPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/web_application_firewall_policy#request_body_check WebApplicationFirewallPolicy#request_body_check}.
	RequestBodyCheck interface{} `field:"optional" json:"requestBodyCheck" yaml:"requestBodyCheck"`
}

