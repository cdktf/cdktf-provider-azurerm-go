package orbitalcontactprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrbitalContactProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#auto_tracking OrbitalContactProfile#auto_tracking}.
	AutoTracking *string `field:"required" json:"autoTracking" yaml:"autoTracking"`
	// links block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#links OrbitalContactProfile#links}
	Links interface{} `field:"required" json:"links" yaml:"links"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#location OrbitalContactProfile#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#minimum_variable_contact_duration OrbitalContactProfile#minimum_variable_contact_duration}.
	MinimumVariableContactDuration *string `field:"required" json:"minimumVariableContactDuration" yaml:"minimumVariableContactDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#name OrbitalContactProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#network_configuration_subnet_id OrbitalContactProfile#network_configuration_subnet_id}.
	NetworkConfigurationSubnetId *string `field:"required" json:"networkConfigurationSubnetId" yaml:"networkConfigurationSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#resource_group_name OrbitalContactProfile#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#event_hub_uri OrbitalContactProfile#event_hub_uri}.
	EventHubUri *string `field:"optional" json:"eventHubUri" yaml:"eventHubUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#id OrbitalContactProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#minimum_elevation_degrees OrbitalContactProfile#minimum_elevation_degrees}.
	MinimumElevationDegrees *float64 `field:"optional" json:"minimumElevationDegrees" yaml:"minimumElevationDegrees"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#tags OrbitalContactProfile#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.57.0/docs/resources/orbital_contact_profile#timeouts OrbitalContactProfile#timeouts}
	Timeouts *OrbitalContactProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

