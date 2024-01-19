// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sharedimageversion

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/sharedimageversion/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/shared_image_version azurerm_shared_image_version}.
type SharedImageVersion interface {
	cdktf.TerraformResource
	BlobUri() *string
	SetBlobUri(val *string)
	BlobUriInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	DeletionOfReplicatedLocationsEnabled() interface{}
	SetDeletionOfReplicatedLocationsEnabled(val interface{})
	DeletionOfReplicatedLocationsEnabledInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	GalleryName() *string
	SetGalleryName(val *string)
	GalleryNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	ImageName() *string
	SetImageName(val *string)
	ImageNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	ManagedImageId() *string
	SetManagedImageId(val *string)
	ManagedImageIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OsDiskSnapshotId() *string
	SetOsDiskSnapshotId(val *string)
	OsDiskSnapshotIdInput() *string
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
	ReplicationMode() *string
	SetReplicationMode(val *string)
	ReplicationModeInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TargetRegion() SharedImageVersionTargetRegionList
	TargetRegionInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SharedImageVersionTimeoutsOutputReference
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
	PutTargetRegion(value interface{})
	PutTimeouts(value *SharedImageVersionTimeouts)
	ResetBlobUri()
	ResetDeletionOfReplicatedLocationsEnabled()
	ResetEndOfLifeDate()
	ResetExcludeFromLatest()
	ResetId()
	ResetManagedImageId()
	ResetOsDiskSnapshotId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetReplicationMode()
	ResetStorageAccountId()
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

// The jsii proxy struct for SharedImageVersion
type jsiiProxy_SharedImageVersion struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SharedImageVersion) BlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) BlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) DeletionOfReplicatedLocationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionOfReplicatedLocationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) DeletionOfReplicatedLocationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionOfReplicatedLocationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) EndOfLifeDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfLifeDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) EndOfLifeDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endOfLifeDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ExcludeFromLatest() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeFromLatest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ExcludeFromLatestInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeFromLatestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) GalleryName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) GalleryNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"galleryNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ImageName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ImageNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ManagedImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ManagedImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) OsDiskSnapshotId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskSnapshotId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) OsDiskSnapshotIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"osDiskSnapshotIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ReplicationMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ReplicationModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TargetRegion() SharedImageVersionTargetRegionList {
	var returns SharedImageVersionTargetRegionList
	_jsii_.Get(
		j,
		"targetRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TargetRegionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"targetRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) Timeouts() SharedImageVersionTimeoutsOutputReference {
	var returns SharedImageVersionTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SharedImageVersion) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/shared_image_version azurerm_shared_image_version} Resource.
func NewSharedImageVersion(scope constructs.Construct, id *string, config *SharedImageVersionConfig) SharedImageVersion {
	_init_.Initialize()

	if err := validateNewSharedImageVersionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SharedImageVersion{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/shared_image_version azurerm_shared_image_version} Resource.
func NewSharedImageVersion_Override(s SharedImageVersion, scope constructs.Construct, id *string, config *SharedImageVersionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetBlobUri(val *string) {
	if err := j.validateSetBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobUri",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetDeletionOfReplicatedLocationsEnabled(val interface{}) {
	if err := j.validateSetDeletionOfReplicatedLocationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletionOfReplicatedLocationsEnabled",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetEndOfLifeDate(val *string) {
	if err := j.validateSetEndOfLifeDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endOfLifeDate",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetExcludeFromLatest(val interface{}) {
	if err := j.validateSetExcludeFromLatestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeFromLatest",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetGalleryName(val *string) {
	if err := j.validateSetGalleryNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"galleryName",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetImageName(val *string) {
	if err := j.validateSetImageNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageName",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetManagedImageId(val *string) {
	if err := j.validateSetManagedImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedImageId",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetOsDiskSnapshotId(val *string) {
	if err := j.validateSetOsDiskSnapshotIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osDiskSnapshotId",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetReplicationMode(val *string) {
	if err := j.validateSetReplicationModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationMode",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

func (j *jsiiProxy_SharedImageVersion)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

// Generates CDKTF code for importing a SharedImageVersion resource upon running "cdktf plan <stack-name>".
func SharedImageVersion_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSharedImageVersion_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
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
func SharedImageVersion_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSharedImageVersion_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SharedImageVersion_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSharedImageVersion_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SharedImageVersion_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSharedImageVersion_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SharedImageVersion_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sharedImageVersion.SharedImageVersion",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SharedImageVersion) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SharedImageVersion) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SharedImageVersion) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SharedImageVersion) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SharedImageVersion) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SharedImageVersion) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SharedImageVersion) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SharedImageVersion) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SharedImageVersion) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SharedImageVersion) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SharedImageVersion) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SharedImageVersion) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersion) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SharedImageVersion) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SharedImageVersion) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SharedImageVersion) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SharedImageVersion) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SharedImageVersion) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SharedImageVersion) PutTargetRegion(value interface{}) {
	if err := s.validatePutTargetRegionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTargetRegion",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImageVersion) PutTimeouts(value *SharedImageVersionTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetBlobUri() {
	_jsii_.InvokeVoid(
		s,
		"resetBlobUri",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetDeletionOfReplicatedLocationsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetDeletionOfReplicatedLocationsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetEndOfLifeDate() {
	_jsii_.InvokeVoid(
		s,
		"resetEndOfLifeDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetExcludeFromLatest() {
	_jsii_.InvokeVoid(
		s,
		"resetExcludeFromLatest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetManagedImageId() {
	_jsii_.InvokeVoid(
		s,
		"resetManagedImageId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetOsDiskSnapshotId() {
	_jsii_.InvokeVoid(
		s,
		"resetOsDiskSnapshotId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetReplicationMode() {
	_jsii_.InvokeVoid(
		s,
		"resetReplicationMode",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetStorageAccountId() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageAccountId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetTags() {
	_jsii_.InvokeVoid(
		s,
		"resetTags",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SharedImageVersion) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersion) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersion) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersion) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersion) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SharedImageVersion) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

