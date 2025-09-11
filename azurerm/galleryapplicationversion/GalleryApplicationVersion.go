// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package galleryapplicationversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/galleryapplicationversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/gallery_application_version azurerm_gallery_application_version}.
type GalleryApplicationVersion interface {
	cdktf.TerraformResource
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ConfigFile() *string
	SetConfigFile(val *string)
	ConfigFileInput() *string
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
	EnableHealthCheck() interface{}
	SetEnableHealthCheck(val interface{})
	EnableHealthCheckInput() interface{}
	EndOfLifeDate() *string
	SetEndOfLifeDate(val *string)
	EndOfLifeDateInput() *string
	ExcludeFromLatest() interface{}
	SetExcludeFromLatest(val interface{})
	ExcludeFromLatestInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryApplicationId() *string
	SetGalleryApplicationId(val *string)
	GalleryApplicationIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManageAction() GalleryApplicationVersionManageActionOutputReference
	ManageActionInput() *GalleryApplicationVersionManageAction
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PackageFile() *string
	SetPackageFile(val *string)
	PackageFileInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Source() GalleryApplicationVersionSourceOutputReference
	SourceInput() *GalleryApplicationVersionSource
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetRegion() GalleryApplicationVersionTargetRegionList
	TargetRegionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() GalleryApplicationVersionTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutManageAction(value *GalleryApplicationVersionManageAction)
	PutSource(value *GalleryApplicationVersionSource)
	PutTargetRegion(value interface{})
	PutTimeouts(value *GalleryApplicationVersionTimeouts)
	ResetConfigFile()
	ResetEnableHealthCheck()
	ResetEndOfLifeDate()
	ResetExcludeFromLatest()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPackageFile()
	ResetTags()
	ResetTimeouts()
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

// The jsii proxy struct for GalleryApplicationVersion
type jsiiProxy_GalleryApplicationVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_GalleryApplicationVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ConfigFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ConfigFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) EnableHealthCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHealthCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) EnableHealthCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableHealthCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) EndOfLifeDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfLifeDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) EndOfLifeDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfLifeDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ExcludeFromLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeFromLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ExcludeFromLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeFromLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) GalleryApplicationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryApplicationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) GalleryApplicationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryApplicationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ManageAction() GalleryApplicationVersionManageActionOutputReference {
	var returns GalleryApplicationVersionManageActionOutputReference
	_jsii_.Get(
		j,
		"manageAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) ManageActionInput() *GalleryApplicationVersionManageAction {
	var returns *GalleryApplicationVersionManageAction
	_jsii_.Get(
		j,
		"manageActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) PackageFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) PackageFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"packageFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Source() GalleryApplicationVersionSourceOutputReference {
	var returns GalleryApplicationVersionSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) SourceInput() *GalleryApplicationVersionSource {
	var returns *GalleryApplicationVersionSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TargetRegion() GalleryApplicationVersionTargetRegionList {
	var returns GalleryApplicationVersionTargetRegionList
	_jsii_.Get(
		j,
		"targetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TargetRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) Timeouts() GalleryApplicationVersionTimeoutsOutputReference {
	var returns GalleryApplicationVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GalleryApplicationVersion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/gallery_application_version azurerm_gallery_application_version} Resource.
func NewGalleryApplicationVersion(scope constructs.Construct, id *string, config *GalleryApplicationVersionConfig) GalleryApplicationVersion {
	_init_.Initialize()

	if err := validateNewGalleryApplicationVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GalleryApplicationVersion{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/gallery_application_version azurerm_gallery_application_version} Resource.
func NewGalleryApplicationVersion_Override(g GalleryApplicationVersion, scope constructs.Construct, id *string, config *GalleryApplicationVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetConfigFile(val *string) {
	if err := j.validateSetConfigFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configFile",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetEnableHealthCheck(val interface{}) {
	if err := j.validateSetEnableHealthCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableHealthCheck",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetEndOfLifeDate(val *string) {
	if err := j.validateSetEndOfLifeDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endOfLifeDate",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetExcludeFromLatest(val interface{}) {
	if err := j.validateSetExcludeFromLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeFromLatest",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetGalleryApplicationId(val *string) {
	if err := j.validateSetGalleryApplicationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"galleryApplicationId",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetPackageFile(val *string) {
	if err := j.validateSetPackageFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"packageFile",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_GalleryApplicationVersion)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a GalleryApplicationVersion resource upon running "cdktf plan <stack-name>".
func GalleryApplicationVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateGalleryApplicationVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
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
func GalleryApplicationVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGalleryApplicationVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GalleryApplicationVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGalleryApplicationVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GalleryApplicationVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGalleryApplicationVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GalleryApplicationVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.galleryApplicationVersion.GalleryApplicationVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) AddMoveTarget(moveTarget *string) {
	if err := g.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := g.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) MoveFromId(id *string) {
	if err := g.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveFromId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := g.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) MoveToId(id *string) {
	if err := g.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"moveToId",
		[]interface{}{id},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) PutManageAction(value *GalleryApplicationVersionManageAction) {
	if err := g.validatePutManageActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putManageAction",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) PutSource(value *GalleryApplicationVersionSource) {
	if err := g.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) PutTargetRegion(value interface{}) {
	if err := g.validatePutTargetRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTargetRegion",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) PutTimeouts(value *GalleryApplicationVersionTimeouts) {
	if err := g.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetConfigFile() {
	_jsii_.InvokeVoid(
		g,
		"resetConfigFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetEnableHealthCheck() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableHealthCheck",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetEndOfLifeDate() {
	_jsii_.InvokeVoid(
		g,
		"resetEndOfLifeDate",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetExcludeFromLatest() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludeFromLatest",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetId() {
	_jsii_.InvokeVoid(
		g,
		"resetId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetPackageFile() {
	_jsii_.InvokeVoid(
		g,
		"resetPackageFile",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetTags() {
	_jsii_.InvokeVoid(
		g,
		"resetTags",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		g,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GalleryApplicationVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GalleryApplicationVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

