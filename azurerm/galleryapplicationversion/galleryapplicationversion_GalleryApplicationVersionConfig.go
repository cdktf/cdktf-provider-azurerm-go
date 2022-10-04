package galleryapplicationversion

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GalleryApplicationVersionConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#gallery_application_id GalleryApplicationVersion#gallery_application_id}.
	GalleryApplicationId *string `field:"required" json:"galleryApplicationId" yaml:"galleryApplicationId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#location GalleryApplicationVersion#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// manage_action block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#manage_action GalleryApplicationVersion#manage_action}
	ManageAction *GalleryApplicationVersionManageAction `field:"required" json:"manageAction" yaml:"manageAction"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#name GalleryApplicationVersion#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#source GalleryApplicationVersion#source}
	Source *GalleryApplicationVersionSource `field:"required" json:"source" yaml:"source"`
	// target_region block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#target_region GalleryApplicationVersion#target_region}
	TargetRegion interface{} `field:"required" json:"targetRegion" yaml:"targetRegion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#enable_health_check GalleryApplicationVersion#enable_health_check}.
	EnableHealthCheck interface{} `field:"optional" json:"enableHealthCheck" yaml:"enableHealthCheck"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#end_of_life_date GalleryApplicationVersion#end_of_life_date}.
	EndOfLifeDate *string `field:"optional" json:"endOfLifeDate" yaml:"endOfLifeDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#exclude_from_latest GalleryApplicationVersion#exclude_from_latest}.
	ExcludeFromLatest interface{} `field:"optional" json:"excludeFromLatest" yaml:"excludeFromLatest"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#id GalleryApplicationVersion#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#tags GalleryApplicationVersion#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/gallery_application_version#timeouts GalleryApplicationVersion#timeouts}
	Timeouts *GalleryApplicationVersionTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

