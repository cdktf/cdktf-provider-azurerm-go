package containergroup


type ContainerGroupContainerLivenessProbeHttpGet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#http_headers ContainerGroup#http_headers}.
	HttpHeaders *map[string]*string `field:"optional" json:"httpHeaders" yaml:"httpHeaders"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#path ContainerGroup#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#port ContainerGroup#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_group#scheme ContainerGroup#scheme}.
	Scheme *string `field:"optional" json:"scheme" yaml:"scheme"`
}

