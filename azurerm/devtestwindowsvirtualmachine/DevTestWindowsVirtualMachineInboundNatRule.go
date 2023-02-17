package devtestwindowsvirtualmachine


type DevTestWindowsVirtualMachineInboundNatRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_windows_virtual_machine#backend_port DevTestWindowsVirtualMachine#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_windows_virtual_machine#protocol DevTestWindowsVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

