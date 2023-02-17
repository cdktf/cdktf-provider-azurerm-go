package devtestlinuxvirtualmachine


type DevTestLinuxVirtualMachineInboundNatRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#backend_port DevTestLinuxVirtualMachine#backend_port}.
	BackendPort *float64 `field:"required" json:"backendPort" yaml:"backendPort"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_linux_virtual_machine#protocol DevTestLinuxVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
}

