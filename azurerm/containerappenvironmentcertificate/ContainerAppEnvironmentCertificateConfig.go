package containerappenvironmentcertificate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppEnvironmentCertificateConfig struct {
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
	// The Certificate Private Key as a base64 encoded PFX or PEM.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#certificate_blob_base64 ContainerAppEnvironmentCertificate#certificate_blob_base64}
	CertificateBlobBase64 *string `field:"required" json:"certificateBlobBase64" yaml:"certificateBlobBase64"`
	// The password for the Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#certificate_password ContainerAppEnvironmentCertificate#certificate_password}
	CertificatePassword *string `field:"required" json:"certificatePassword" yaml:"certificatePassword"`
	// The Container App Managed Environment ID to configure this Certificate on.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#container_app_environment_id ContainerAppEnvironmentCertificate#container_app_environment_id}
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The name of the Container Apps Environment Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#name ContainerAppEnvironmentCertificate#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#id ContainerAppEnvironmentCertificate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#tags ContainerAppEnvironmentCertificate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#timeouts ContainerAppEnvironmentCertificate#timeouts}
	Timeouts *ContainerAppEnvironmentCertificateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

