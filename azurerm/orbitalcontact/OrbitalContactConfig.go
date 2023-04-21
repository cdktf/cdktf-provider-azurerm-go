package orbitalcontact

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrbitalContactConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#contact_profile_id OrbitalContact#contact_profile_id}.
	ContactProfileId *string `field:"required" json:"contactProfileId" yaml:"contactProfileId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#ground_station_name OrbitalContact#ground_station_name}.
	GroundStationName *string `field:"required" json:"groundStationName" yaml:"groundStationName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#name OrbitalContact#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#reservation_end_time OrbitalContact#reservation_end_time}.
	ReservationEndTime *string `field:"required" json:"reservationEndTime" yaml:"reservationEndTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#reservation_start_time OrbitalContact#reservation_start_time}.
	ReservationStartTime *string `field:"required" json:"reservationStartTime" yaml:"reservationStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#spacecraft_id OrbitalContact#spacecraft_id}.
	SpacecraftId *string `field:"required" json:"spacecraftId" yaml:"spacecraftId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#id OrbitalContact#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/orbital_contact#timeouts OrbitalContact#timeouts}
	Timeouts *OrbitalContactTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

