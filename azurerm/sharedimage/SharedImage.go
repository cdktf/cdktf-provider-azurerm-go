// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/sharedimage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/shared_image azurerm_shared_image}.
type SharedImage interface {
	cdktf.TerraformResource
	AcceleratedNetworkSupportEnabled() interface{}
	SetAcceleratedNetworkSupportEnabled(val interface{})
	AcceleratedNetworkSupportEnabledInput() interface{}
	Architecture() *string
	SetArchitecture(val *string)
	ArchitectureInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfidentialVmEnabled() interface{}
	SetConfidentialVmEnabled(val interface{})
	ConfidentialVmEnabledInput() interface{}
	ConfidentialVmSupported() interface{}
	SetConfidentialVmSupported(val interface{})
	ConfidentialVmSupportedInput() interface{}
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DiskTypesNotAllowed() *[]*string
	SetDiskTypesNotAllowed(val *[]*string)
	DiskTypesNotAllowedInput() *[]*string
	EndOfLifeDate() *string
	SetEndOfLifeDate(val *string)
	EndOfLifeDateInput() *string
	Eula() *string
	SetEula(val *string)
	EulaInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryName() *string
	SetGalleryName(val *string)
	GalleryNameInput() *string
	HyperVGeneration() *string
	SetHyperVGeneration(val *string)
	HyperVGenerationInput() *string
	Id() *string
	SetId(val *string)
	Identifier() SharedImageIdentifierOutputReference
	IdentifierInput() *SharedImageIdentifier
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxRecommendedMemoryInGb() *float64
	SetMaxRecommendedMemoryInGb(val *float64)
	MaxRecommendedMemoryInGbInput() *float64
	MaxRecommendedVcpuCount() *float64
	SetMaxRecommendedVcpuCount(val *float64)
	MaxRecommendedVcpuCountInput() *float64
	MinRecommendedMemoryInGb() *float64
	SetMinRecommendedMemoryInGb(val *float64)
	MinRecommendedMemoryInGbInput() *float64
	MinRecommendedVcpuCount() *float64
	SetMinRecommendedVcpuCount(val *float64)
	MinRecommendedVcpuCountInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OsType() *string
	SetOsType(val *string)
	OsTypeInput() *string
	PrivacyStatementUri() *string
	SetPrivacyStatementUri(val *string)
	PrivacyStatementUriInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PurchasePlan() SharedImagePurchasePlanOutputReference
	PurchasePlanInput() *SharedImagePurchasePlan
	// Experimental.
	RawOverrides() interface{}
	ReleaseNoteUri() *string
	SetReleaseNoteUri(val *string)
	ReleaseNoteUriInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Specialized() interface{}
	SetSpecialized(val interface{})
	SpecializedInput() interface{}
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SharedImageTimeoutsOutputReference
	TimeoutsInput() interface{}
	TrustedLaunchEnabled() interface{}
	SetTrustedLaunchEnabled(val interface{})
	TrustedLaunchEnabledInput() interface{}
	TrustedLaunchSupported() interface{}
	SetTrustedLaunchSupported(val interface{})
	TrustedLaunchSupportedInput() interface{}
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutIdentifier(value *SharedImageIdentifier)
	PutPurchasePlan(value *SharedImagePurchasePlan)
	PutTimeouts(value *SharedImageTimeouts)
	ResetAcceleratedNetworkSupportEnabled()
	ResetArchitecture()
	ResetConfidentialVmEnabled()
	ResetConfidentialVmSupported()
	ResetDescription()
	ResetDiskTypesNotAllowed()
	ResetEndOfLifeDate()
	ResetEula()
	ResetHyperVGeneration()
	ResetId()
	ResetMaxRecommendedMemoryInGb()
	ResetMaxRecommendedVcpuCount()
	ResetMinRecommendedMemoryInGb()
	ResetMinRecommendedVcpuCount()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrivacyStatementUri()
	ResetPurchasePlan()
	ResetReleaseNoteUri()
	ResetSpecialized()
	ResetTags()
	ResetTimeouts()
	ResetTrustedLaunchEnabled()
	ResetTrustedLaunchSupported()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for SharedImage
type jsiiProxy_SharedImage struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SharedImage) AcceleratedNetworkSupportEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkSupportEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) AcceleratedNetworkSupportEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkSupportEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Architecture() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ArchitectureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"architectureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ConfidentialVmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialVmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ConfidentialVmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialVmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ConfidentialVmSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialVmSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ConfidentialVmSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"confidentialVmSupportedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) DiskTypesNotAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"diskTypesNotAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) DiskTypesNotAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"diskTypesNotAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) EndOfLifeDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfLifeDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) EndOfLifeDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfLifeDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Eula() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eula",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) EulaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eulaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) GalleryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) GalleryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) HyperVGeneration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hyperVGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) HyperVGenerationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hyperVGenerationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Identifier() SharedImageIdentifierOutputReference {
	var returns SharedImageIdentifierOutputReference
	_jsii_.Get(
		j,
		"identifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) IdentifierInput() *SharedImageIdentifier {
	var returns *SharedImageIdentifier
	_jsii_.Get(
		j,
		"identifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MaxRecommendedMemoryInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRecommendedMemoryInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MaxRecommendedMemoryInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRecommendedMemoryInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MaxRecommendedVcpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRecommendedVcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MaxRecommendedVcpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRecommendedVcpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MinRecommendedMemoryInGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRecommendedMemoryInGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MinRecommendedMemoryInGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRecommendedMemoryInGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MinRecommendedVcpuCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRecommendedVcpuCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) MinRecommendedVcpuCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minRecommendedVcpuCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) OsType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) OsTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) PrivacyStatementUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyStatementUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) PrivacyStatementUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privacyStatementUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) PurchasePlan() SharedImagePurchasePlanOutputReference {
	var returns SharedImagePurchasePlanOutputReference
	_jsii_.Get(
		j,
		"purchasePlan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) PurchasePlanInput() *SharedImagePurchasePlan {
	var returns *SharedImagePurchasePlan
	_jsii_.Get(
		j,
		"purchasePlanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ReleaseNoteUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNoteUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ReleaseNoteUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"releaseNoteUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Specialized() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specialized",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) SpecializedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"specializedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) Timeouts() SharedImageTimeoutsOutputReference {
	var returns SharedImageTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TrustedLaunchEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedLaunchEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TrustedLaunchEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedLaunchEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TrustedLaunchSupported() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedLaunchSupported",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImage) TrustedLaunchSupportedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"trustedLaunchSupportedInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/shared_image azurerm_shared_image} Resource.
func NewSharedImage(scope constructs.Construct, id *string, config *SharedImageConfig) SharedImage {
	_init_.Initialize()

	if err := validateNewSharedImageParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SharedImage{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/shared_image azurerm_shared_image} Resource.
func NewSharedImage_Override(s SharedImage, scope constructs.Construct, id *string, config *SharedImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SharedImage)SetAcceleratedNetworkSupportEnabled(val interface{}) {
	if err := j.validateSetAcceleratedNetworkSupportEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratedNetworkSupportEnabled",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetArchitecture(val *string) {
	if err := j.validateSetArchitectureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetConfidentialVmEnabled(val interface{}) {
	if err := j.validateSetConfidentialVmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialVmEnabled",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetConfidentialVmSupported(val interface{}) {
	if err := j.validateSetConfidentialVmSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"confidentialVmSupported",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetDiskTypesNotAllowed(val *[]*string) {
	if err := j.validateSetDiskTypesNotAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskTypesNotAllowed",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetEndOfLifeDate(val *string) {
	if err := j.validateSetEndOfLifeDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endOfLifeDate",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetEula(val *string) {
	if err := j.validateSetEulaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"eula",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetGalleryName(val *string) {
	if err := j.validateSetGalleryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"galleryName",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetHyperVGeneration(val *string) {
	if err := j.validateSetHyperVGenerationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hyperVGeneration",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetMaxRecommendedMemoryInGb(val *float64) {
	if err := j.validateSetMaxRecommendedMemoryInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRecommendedMemoryInGb",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetMaxRecommendedVcpuCount(val *float64) {
	if err := j.validateSetMaxRecommendedVcpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRecommendedVcpuCount",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetMinRecommendedMemoryInGb(val *float64) {
	if err := j.validateSetMinRecommendedMemoryInGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRecommendedMemoryInGb",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetMinRecommendedVcpuCount(val *float64) {
	if err := j.validateSetMinRecommendedVcpuCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minRecommendedVcpuCount",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetOsType(val *string) {
	if err := j.validateSetOsTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetPrivacyStatementUri(val *string) {
	if err := j.validateSetPrivacyStatementUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privacyStatementUri",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetReleaseNoteUri(val *string) {
	if err := j.validateSetReleaseNoteUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"releaseNoteUri",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetSpecialized(val interface{}) {
	if err := j.validateSetSpecializedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"specialized",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetTrustedLaunchEnabled(val interface{}) {
	if err := j.validateSetTrustedLaunchEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedLaunchEnabled",
		val,
	)
}

func (j *jsiiProxy_SharedImage)SetTrustedLaunchSupported(val interface{}) {
	if err := j.validateSetTrustedLaunchSupportedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedLaunchSupported",
		val,
	)
}

// Generates CDKTF code for importing a SharedImage resource upon running "cdktf plan <stack-name>".
func SharedImage_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSharedImage_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func SharedImage_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSharedImage_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SharedImage_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSharedImage_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SharedImage_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSharedImage_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SharedImage_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SharedImage) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SharedImage) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SharedImage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SharedImage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SharedImage) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SharedImage) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SharedImage) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SharedImage) PutIdentifier(value *SharedImageIdentifier) {
	if err := s.validatePutIdentifierParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putIdentifier",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImage) PutPurchasePlan(value *SharedImagePurchasePlan) {
	if err := s.validatePutPurchasePlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putPurchasePlan",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImage) PutTimeouts(value *SharedImageTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImage) ResetAcceleratedNetworkSupportEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetAcceleratedNetworkSupportEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetArchitecture() {
	_jsii_.InvokeVoid(
		s,
		"resetArchitecture",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetConfidentialVmEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidentialVmEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetConfidentialVmSupported() {
	_jsii_.InvokeVoid(
		s,
		"resetConfidentialVmSupported",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetDescription() {
	_jsii_.InvokeVoid(
		s,
		"resetDescription",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetDiskTypesNotAllowed() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskTypesNotAllowed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetEndOfLifeDate() {
	_jsii_.InvokeVoid(
		s,
		"resetEndOfLifeDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetEula() {
	_jsii_.InvokeVoid(
		s,
		"resetEula",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetHyperVGeneration() {
	_jsii_.InvokeVoid(
		s,
		"resetHyperVGeneration",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetMaxRecommendedMemoryInGb() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxRecommendedMemoryInGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetMaxRecommendedVcpuCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxRecommendedVcpuCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetMinRecommendedMemoryInGb() {
	_jsii_.InvokeVoid(
		s,
		"resetMinRecommendedMemoryInGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetMinRecommendedVcpuCount() {
	_jsii_.InvokeVoid(
		s,
		"resetMinRecommendedVcpuCount",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetPrivacyStatementUri() {
	_jsii_.InvokeVoid(
		s,
		"resetPrivacyStatementUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetPurchasePlan() {
	_jsii_.InvokeVoid(
		s,
		"resetPurchasePlan",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetReleaseNoteUri() {
	_jsii_.InvokeVoid(
		s,
		"resetReleaseNoteUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetSpecialized() {
	_jsii_.InvokeVoid(
		s,
		"resetSpecialized",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetTrustedLaunchEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetTrustedLaunchEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) ResetTrustedLaunchSupported() {
	_jsii_.InvokeVoid(
		s,
		"resetTrustedLaunchSupported",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImage) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

