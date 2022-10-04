package devtestvirtualnetwork


type DevTestVirtualNetworkSubnet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_virtual_network#use_in_virtual_machine_creation DevTestVirtualNetwork#use_in_virtual_machine_creation}.
	UseInVirtualMachineCreation *string `field:"optional" json:"useInVirtualMachineCreation" yaml:"useInVirtualMachineCreation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/dev_test_virtual_network#use_public_ip_address DevTestVirtualNetwork#use_public_ip_address}.
	UsePublicIpAddress *string `field:"optional" json:"usePublicIpAddress" yaml:"usePublicIpAddress"`
}

