package springcloudgatewayrouteconfig

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudGatewayRouteConfigConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#name SpringCloudGatewayRouteConfig#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#spring_cloud_gateway_id SpringCloudGatewayRouteConfig#spring_cloud_gateway_id}.
	SpringCloudGatewayId *string `field:"required" json:"springCloudGatewayId" yaml:"springCloudGatewayId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#id SpringCloudGatewayRouteConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// open_api block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#open_api SpringCloudGatewayRouteConfig#open_api}
	OpenApi *SpringCloudGatewayRouteConfigOpenApi `field:"optional" json:"openApi" yaml:"openApi"`
	// route block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#route SpringCloudGatewayRouteConfig#route}
	Route interface{} `field:"optional" json:"route" yaml:"route"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#spring_cloud_app_id SpringCloudGatewayRouteConfig#spring_cloud_app_id}.
	SpringCloudAppId *string `field:"optional" json:"springCloudAppId" yaml:"springCloudAppId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_gateway_route_config#timeouts SpringCloudGatewayRouteConfig#timeouts}
	Timeouts *SpringCloudGatewayRouteConfigTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}
