package iotsecuritydevicegroup


type IotSecurityDeviceGroupRangeRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_device_group#duration IotSecurityDeviceGroup#duration}.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_device_group#max IotSecurityDeviceGroup#max}.
	Max *float64 `field:"required" json:"max" yaml:"max"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_device_group#min IotSecurityDeviceGroup#min}.
	Min *float64 `field:"required" json:"min" yaml:"min"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iot_security_device_group#type IotSecurityDeviceGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}
