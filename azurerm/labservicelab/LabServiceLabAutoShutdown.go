package labservicelab


type LabServiceLabAutoShutdown struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#disconnect_delay LabServiceLab#disconnect_delay}.
	DisconnectDelay *string `field:"optional" json:"disconnectDelay" yaml:"disconnectDelay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#idle_delay LabServiceLab#idle_delay}.
	IdleDelay *string `field:"optional" json:"idleDelay" yaml:"idleDelay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#no_connect_delay LabServiceLab#no_connect_delay}.
	NoConnectDelay *string `field:"optional" json:"noConnectDelay" yaml:"noConnectDelay"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lab_service_lab#shutdown_on_idle LabServiceLab#shutdown_on_idle}.
	ShutdownOnIdle *string `field:"optional" json:"shutdownOnIdle" yaml:"shutdownOnIdle"`
}
