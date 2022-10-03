package sharedimage

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/sharedimage/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image azurerm_shared_image}.
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
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutIdentifier(value *SharedImageIdentifier)
	PutPurchasePlan(value *SharedImagePurchasePlan)
	PutTimeouts(value *SharedImageTimeouts)
	ResetAcceleratedNetworkSupportEnabled()
	ResetArchitecture()
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
	SynthesizeAttributes() *map[string]interface{}
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

func (j *jsiiProxy_SharedImage) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image azurerm_shared_image} Resource.
func NewSharedImage(scope constructs.Construct, id *string, config *SharedImageConfig) SharedImage {
	_init_.Initialize()

	j := jsiiProxy_SharedImage{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image azurerm_shared_image} Resource.
func NewSharedImage_Override(s SharedImage, scope constructs.Construct, id *string, config *SharedImageConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SharedImage) SetAcceleratedNetworkSupportEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"acceleratedNetworkSupportEnabled",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetArchitecture(val *string) {
	_jsii_.Set(
		j,
		"architecture",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetDiskTypesNotAllowed(val *[]*string) {
	_jsii_.Set(
		j,
		"diskTypesNotAllowed",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetEndOfLifeDate(val *string) {
	_jsii_.Set(
		j,
		"endOfLifeDate",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetEula(val *string) {
	_jsii_.Set(
		j,
		"eula",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetGalleryName(val *string) {
	_jsii_.Set(
		j,
		"galleryName",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetHyperVGeneration(val *string) {
	_jsii_.Set(
		j,
		"hyperVGeneration",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetLocation(val *string) {
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetMaxRecommendedMemoryInGb(val *float64) {
	_jsii_.Set(
		j,
		"maxRecommendedMemoryInGb",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetMaxRecommendedVcpuCount(val *float64) {
	_jsii_.Set(
		j,
		"maxRecommendedVcpuCount",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetMinRecommendedMemoryInGb(val *float64) {
	_jsii_.Set(
		j,
		"minRecommendedMemoryInGb",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetMinRecommendedVcpuCount(val *float64) {
	_jsii_.Set(
		j,
		"minRecommendedVcpuCount",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetOsType(val *string) {
	_jsii_.Set(
		j,
		"osType",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetPrivacyStatementUri(val *string) {
	_jsii_.Set(
		j,
		"privacyStatementUri",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetReleaseNoteUri(val *string) {
	_jsii_.Set(
		j,
		"releaseNoteUri",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetResourceGroupName(val *string) {
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetSpecialized(val interface{}) {
	_jsii_.Set(
		j,
		"specialized",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetTags(val *map[string]*string) {
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_SharedImage) SetTrustedLaunchEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"trustedLaunchEnabled",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImage.SharedImage",
		"isConstruct",
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

func (s *jsiiProxy_SharedImage) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SharedImage) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImage) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SharedImage) PutIdentifier(value *SharedImageIdentifier) {
	_jsii_.InvokeVoid(
		s,
		"putIdentifier",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImage) PutPurchasePlan(value *SharedImagePurchasePlan) {
	_jsii_.InvokeVoid(
		s,
		"putPurchasePlan",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImage) PutTimeouts(value *SharedImageTimeouts) {
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

type SharedImageConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#gallery_name SharedImage#gallery_name}.
	GalleryName *string `field:"required" json:"galleryName" yaml:"galleryName"`
	// identifier block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#identifier SharedImage#identifier}
	Identifier *SharedImageIdentifier `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#location SharedImage#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#name SharedImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#os_type SharedImage#os_type}.
	OsType *string `field:"required" json:"osType" yaml:"osType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#resource_group_name SharedImage#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#accelerated_network_support_enabled SharedImage#accelerated_network_support_enabled}.
	AcceleratedNetworkSupportEnabled interface{} `field:"optional" json:"acceleratedNetworkSupportEnabled" yaml:"acceleratedNetworkSupportEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#architecture SharedImage#architecture}.
	Architecture *string `field:"optional" json:"architecture" yaml:"architecture"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#description SharedImage#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#disk_types_not_allowed SharedImage#disk_types_not_allowed}.
	DiskTypesNotAllowed *[]*string `field:"optional" json:"diskTypesNotAllowed" yaml:"diskTypesNotAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#end_of_life_date SharedImage#end_of_life_date}.
	EndOfLifeDate *string `field:"optional" json:"endOfLifeDate" yaml:"endOfLifeDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#eula SharedImage#eula}.
	Eula *string `field:"optional" json:"eula" yaml:"eula"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#hyper_v_generation SharedImage#hyper_v_generation}.
	HyperVGeneration *string `field:"optional" json:"hyperVGeneration" yaml:"hyperVGeneration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#id SharedImage#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#max_recommended_memory_in_gb SharedImage#max_recommended_memory_in_gb}.
	MaxRecommendedMemoryInGb *float64 `field:"optional" json:"maxRecommendedMemoryInGb" yaml:"maxRecommendedMemoryInGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#max_recommended_vcpu_count SharedImage#max_recommended_vcpu_count}.
	MaxRecommendedVcpuCount *float64 `field:"optional" json:"maxRecommendedVcpuCount" yaml:"maxRecommendedVcpuCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#min_recommended_memory_in_gb SharedImage#min_recommended_memory_in_gb}.
	MinRecommendedMemoryInGb *float64 `field:"optional" json:"minRecommendedMemoryInGb" yaml:"minRecommendedMemoryInGb"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#min_recommended_vcpu_count SharedImage#min_recommended_vcpu_count}.
	MinRecommendedVcpuCount *float64 `field:"optional" json:"minRecommendedVcpuCount" yaml:"minRecommendedVcpuCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#privacy_statement_uri SharedImage#privacy_statement_uri}.
	PrivacyStatementUri *string `field:"optional" json:"privacyStatementUri" yaml:"privacyStatementUri"`
	// purchase_plan block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#purchase_plan SharedImage#purchase_plan}
	PurchasePlan *SharedImagePurchasePlan `field:"optional" json:"purchasePlan" yaml:"purchasePlan"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#release_note_uri SharedImage#release_note_uri}.
	ReleaseNoteUri *string `field:"optional" json:"releaseNoteUri" yaml:"releaseNoteUri"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#specialized SharedImage#specialized}.
	Specialized interface{} `field:"optional" json:"specialized" yaml:"specialized"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#tags SharedImage#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#timeouts SharedImage#timeouts}
	Timeouts *SharedImageTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#trusted_launch_enabled SharedImage#trusted_launch_enabled}.
	TrustedLaunchEnabled interface{} `field:"optional" json:"trustedLaunchEnabled" yaml:"trustedLaunchEnabled"`
}

type SharedImageIdentifier struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#offer SharedImage#offer}.
	Offer *string `field:"required" json:"offer" yaml:"offer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#publisher SharedImage#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#sku SharedImage#sku}.
	Sku *string `field:"required" json:"sku" yaml:"sku"`
}

type SharedImageIdentifierOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SharedImageIdentifier
	SetInternalValue(val *SharedImageIdentifier)
	Offer() *string
	SetOffer(val *string)
	OfferInput() *string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SharedImageIdentifierOutputReference
type jsiiProxy_SharedImageIdentifierOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) InternalValue() *SharedImageIdentifier {
	var returns *SharedImageIdentifier
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) Offer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) OfferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSharedImageIdentifierOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SharedImageIdentifierOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SharedImageIdentifierOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImageIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSharedImageIdentifierOutputReference_Override(s SharedImageIdentifierOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImageIdentifierOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetInternalValue(val *SharedImageIdentifier) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetOffer(val *string) {
	_jsii_.Set(
		j,
		"offer",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetSku(val *string) {
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SharedImageIdentifierOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageIdentifierOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SharedImagePurchasePlan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#name SharedImage#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#product SharedImage#product}.
	Product *string `field:"optional" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#publisher SharedImage#publisher}.
	Publisher *string `field:"optional" json:"publisher" yaml:"publisher"`
}

type SharedImagePurchasePlanOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SharedImagePurchasePlan
	SetInternalValue(val *SharedImagePurchasePlan)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Product() *string
	SetProduct(val *string)
	ProductInput() *string
	Publisher() *string
	SetPublisher(val *string)
	PublisherInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetProduct()
	ResetPublisher()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SharedImagePurchasePlanOutputReference
type jsiiProxy_SharedImagePurchasePlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) InternalValue() *SharedImagePurchasePlan {
	var returns *SharedImagePurchasePlan
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) Product() *string {
	var returns *string
	_jsii_.Get(
		j,
		"product",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) ProductInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"productInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) Publisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) PublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSharedImagePurchasePlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SharedImagePurchasePlanOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SharedImagePurchasePlanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImagePurchasePlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSharedImagePurchasePlanOutputReference_Override(s SharedImagePurchasePlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImagePurchasePlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetInternalValue(val *SharedImagePurchasePlan) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetProduct(val *string) {
	_jsii_.Set(
		j,
		"product",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetPublisher(val *string) {
	_jsii_.Set(
		j,
		"publisher",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SharedImagePurchasePlanOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) ResetProduct() {
	_jsii_.InvokeVoid(
		s,
		"resetProduct",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) ResetPublisher() {
	_jsii_.InvokeVoid(
		s,
		"resetPublisher",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImagePurchasePlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type SharedImageTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#create SharedImage#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#delete SharedImage#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#read SharedImage#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/shared_image#update SharedImage#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type SharedImageTimeoutsOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	Create() *string
	SetCreate(val *string)
	CreateInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Delete() *string
	SetDelete(val *string)
	DeleteInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Read() *string
	SetRead(val *string)
	ReadInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() *string
	SetUpdate(val *string)
	UpdateInput() *string
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetCreate()
	ResetDelete()
	ResetRead()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SharedImageTimeoutsOutputReference
type jsiiProxy_SharedImageTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewSharedImageTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SharedImageTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_SharedImageTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImageTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSharedImageTimeoutsOutputReference_Override(s SharedImageTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImage.SharedImageTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SharedImageTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		s,
		"resetCreate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		s,
		"resetDelete",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		s,
		"resetRead",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		s,
		"resetUpdate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

