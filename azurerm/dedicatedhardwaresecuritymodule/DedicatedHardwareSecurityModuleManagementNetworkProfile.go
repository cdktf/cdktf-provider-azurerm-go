package dedicatedhardwaresecuritymodule


type DedicatedHardwareSecurityModuleManagementNetworkProfile struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dedicated_hardware_security_module#network_interface_private_ip_addresses DedicatedHardwareSecurityModule#network_interface_private_ip_addresses}.
	NetworkInterfacePrivateIpAddresses *[]*string `field:"required" json:"networkInterfacePrivateIpAddresses" yaml:"networkInterfacePrivateIpAddresses"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dedicated_hardware_security_module#subnet_id DedicatedHardwareSecurityModule#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}
