package expressroutecircuitconnection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRouteCircuitConnectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#address_prefix_ipv4 ExpressRouteCircuitConnection#address_prefix_ipv4}.
	AddressPrefixIpv4 *string `field:"required" json:"addressPrefixIpv4" yaml:"addressPrefixIpv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#name ExpressRouteCircuitConnection#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#peering_id ExpressRouteCircuitConnection#peering_id}.
	PeeringId *string `field:"required" json:"peeringId" yaml:"peeringId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#peer_peering_id ExpressRouteCircuitConnection#peer_peering_id}.
	PeerPeeringId *string `field:"required" json:"peerPeeringId" yaml:"peerPeeringId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#address_prefix_ipv6 ExpressRouteCircuitConnection#address_prefix_ipv6}.
	AddressPrefixIpv6 *string `field:"optional" json:"addressPrefixIpv6" yaml:"addressPrefixIpv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#authorization_key ExpressRouteCircuitConnection#authorization_key}.
	AuthorizationKey *string `field:"optional" json:"authorizationKey" yaml:"authorizationKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#id ExpressRouteCircuitConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/express_route_circuit_connection#timeouts ExpressRouteCircuitConnection#timeouts}
	Timeouts *ExpressRouteCircuitConnectionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

