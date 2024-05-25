// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelmetadata

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/sentinelmetadata/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_metadata azurerm_sentinel_metadata}.
type SentinelMetadata interface {
	cdktf.TerraformResource
	Author() SentinelMetadataAuthorOutputReference
	AuthorInput() *SentinelMetadataAuthor
	Category() SentinelMetadataCategoryOutputReference
	CategoryInput() *SentinelMetadataCategory
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	ContentId() *string
	SetContentId(val *string)
	ContentIdInput() *string
	ContentSchemaVersion() *string
	SetContentSchemaVersion(val *string)
	ContentSchemaVersionInput() *string
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomVersion() *string
	SetCustomVersion(val *string)
	CustomVersionInput() *string
	Dependency() *string
	SetDependency(val *string)
	DependencyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	FirstPublishDate() *string
	SetFirstPublishDate(val *string)
	FirstPublishDateInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	IconId() *string
	SetIconId(val *string)
	IconIdInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Kind() *string
	SetKind(val *string)
	KindInput() *string
	LastPublishDate() *string
	SetLastPublishDate(val *string)
	LastPublishDateInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	ParentId() *string
	SetParentId(val *string)
	ParentIdInput() *string
	PreviewImages() *[]*string
	SetPreviewImages(val *[]*string)
	PreviewImagesDark() *[]*string
	SetPreviewImagesDark(val *[]*string)
	PreviewImagesDarkInput() *[]*string
	PreviewImagesInput() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	Providers() *[]*string
	SetProviders(val *[]*string)
	ProvidersInput() *[]*string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Source() SentinelMetadataSourceOutputReference
	SourceInput() *SentinelMetadataSource
	Support() SentinelMetadataSupportOutputReference
	SupportInput() *SentinelMetadataSupport
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ThreatAnalysisTactics() *[]*string
	SetThreatAnalysisTactics(val *[]*string)
	ThreatAnalysisTacticsInput() *[]*string
	ThreatAnalysisTechniques() *[]*string
	SetThreatAnalysisTechniques(val *[]*string)
	ThreatAnalysisTechniquesInput() *[]*string
	Timeouts() SentinelMetadataTimeoutsOutputReference
	TimeoutsInput() interface{}
	Version() *string
	SetVersion(val *string)
	VersionInput() *string
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
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
	PutAuthor(value *SentinelMetadataAuthor)
	PutCategory(value *SentinelMetadataCategory)
	PutSource(value *SentinelMetadataSource)
	PutSupport(value *SentinelMetadataSupport)
	PutTimeouts(value *SentinelMetadataTimeouts)
	ResetAuthor()
	ResetCategory()
	ResetContentSchemaVersion()
	ResetCustomVersion()
	ResetDependency()
	ResetFirstPublishDate()
	ResetIconId()
	ResetId()
	ResetLastPublishDate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPreviewImages()
	ResetPreviewImagesDark()
	ResetProviders()
	ResetSource()
	ResetSupport()
	ResetThreatAnalysisTactics()
	ResetThreatAnalysisTechniques()
	ResetTimeouts()
	ResetVersion()
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

// The jsii proxy struct for SentinelMetadata
type jsiiProxy_SentinelMetadata struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SentinelMetadata) Author() SentinelMetadataAuthorOutputReference {
	var returns SentinelMetadataAuthorOutputReference
	_jsii_.Get(
		j,
		"author",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) AuthorInput() *SentinelMetadataAuthor {
	var returns *SentinelMetadataAuthor
	_jsii_.Get(
		j,
		"authorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Category() SentinelMetadataCategoryOutputReference {
	var returns SentinelMetadataCategoryOutputReference
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) CategoryInput() *SentinelMetadataCategory {
	var returns *SentinelMetadataCategory
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ContentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ContentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ContentSchemaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentSchemaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ContentSchemaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentSchemaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) CustomVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) CustomVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Dependency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dependency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) DependencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dependencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) FirstPublishDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPublishDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) FirstPublishDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"firstPublishDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) IconId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) IconIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iconIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) KindInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) LastPublishDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastPublishDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) LastPublishDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastPublishDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ParentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ParentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parentIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) PreviewImages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewImages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) PreviewImagesDark() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewImagesDark",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) PreviewImagesDarkInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewImagesDarkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) PreviewImagesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"previewImagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Providers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"providersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Source() SentinelMetadataSourceOutputReference {
	var returns SentinelMetadataSourceOutputReference
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) SourceInput() *SentinelMetadataSource {
	var returns *SentinelMetadataSource
	_jsii_.Get(
		j,
		"sourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Support() SentinelMetadataSupportOutputReference {
	var returns SentinelMetadataSupportOutputReference
	_jsii_.Get(
		j,
		"support",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) SupportInput() *SentinelMetadataSupport {
	var returns *SentinelMetadataSupport
	_jsii_.Get(
		j,
		"supportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ThreatAnalysisTactics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"threatAnalysisTactics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ThreatAnalysisTacticsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"threatAnalysisTacticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ThreatAnalysisTechniques() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"threatAnalysisTechniques",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) ThreatAnalysisTechniquesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"threatAnalysisTechniquesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Timeouts() SentinelMetadataTimeoutsOutputReference {
	var returns SentinelMetadataTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) VersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SentinelMetadata) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_metadata azurerm_sentinel_metadata} Resource.
func NewSentinelMetadata(scope constructs.Construct, id *string, config *SentinelMetadataConfig) SentinelMetadata {
	_init_.Initialize()

	if err := validateNewSentinelMetadataParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SentinelMetadata{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/sentinel_metadata azurerm_sentinel_metadata} Resource.
func NewSentinelMetadata_Override(s SentinelMetadata, scope constructs.Construct, id *string, config *SentinelMetadataConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetContentId(val *string) {
	if err := j.validateSetContentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentId",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetContentSchemaVersion(val *string) {
	if err := j.validateSetContentSchemaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentSchemaVersion",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetCustomVersion(val *string) {
	if err := j.validateSetCustomVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customVersion",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetDependency(val *string) {
	if err := j.validateSetDependencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependency",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetFirstPublishDate(val *string) {
	if err := j.validateSetFirstPublishDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstPublishDate",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetIconId(val *string) {
	if err := j.validateSetIconIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iconId",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetKind(val *string) {
	if err := j.validateSetKindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kind",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetLastPublishDate(val *string) {
	if err := j.validateSetLastPublishDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastPublishDate",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetParentId(val *string) {
	if err := j.validateSetParentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parentId",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetPreviewImages(val *[]*string) {
	if err := j.validateSetPreviewImagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previewImages",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetPreviewImagesDark(val *[]*string) {
	if err := j.validateSetPreviewImagesDarkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"previewImagesDark",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetProviders(val *[]*string) {
	if err := j.validateSetProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providers",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetThreatAnalysisTactics(val *[]*string) {
	if err := j.validateSetThreatAnalysisTacticsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threatAnalysisTactics",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetThreatAnalysisTechniques(val *[]*string) {
	if err := j.validateSetThreatAnalysisTechniquesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"threatAnalysisTechniques",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetVersion(val *string) {
	if err := j.validateSetVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"version",
		val,
	)
}

func (j *jsiiProxy_SentinelMetadata)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

// Generates CDKTF code for importing a SentinelMetadata resource upon running "cdktf plan <stack-name>".
func SentinelMetadata_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSentinelMetadata_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
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
func SentinelMetadata_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelMetadata_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelMetadata_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelMetadata_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SentinelMetadata_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSentinelMetadata_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SentinelMetadata_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.sentinelMetadata.SentinelMetadata",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SentinelMetadata) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SentinelMetadata) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SentinelMetadata) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SentinelMetadata) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelMetadata) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SentinelMetadata) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SentinelMetadata) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SentinelMetadata) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SentinelMetadata) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SentinelMetadata) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SentinelMetadata) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SentinelMetadata) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelMetadata) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SentinelMetadata) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SentinelMetadata) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelMetadata) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SentinelMetadata) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SentinelMetadata) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SentinelMetadata) PutAuthor(value *SentinelMetadataAuthor) {
	if err := s.validatePutAuthorParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putAuthor",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelMetadata) PutCategory(value *SentinelMetadataCategory) {
	if err := s.validatePutCategoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putCategory",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelMetadata) PutSource(value *SentinelMetadataSource) {
	if err := s.validatePutSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSource",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelMetadata) PutSupport(value *SentinelMetadataSupport) {
	if err := s.validatePutSupportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putSupport",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelMetadata) PutTimeouts(value *SentinelMetadataTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetAuthor() {
	_jsii_.InvokeVoid(
		s,
		"resetAuthor",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetCategory() {
	_jsii_.InvokeVoid(
		s,
		"resetCategory",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetContentSchemaVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetContentSchemaVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetCustomVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetCustomVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetDependency() {
	_jsii_.InvokeVoid(
		s,
		"resetDependency",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetFirstPublishDate() {
	_jsii_.InvokeVoid(
		s,
		"resetFirstPublishDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetIconId() {
	_jsii_.InvokeVoid(
		s,
		"resetIconId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetLastPublishDate() {
	_jsii_.InvokeVoid(
		s,
		"resetLastPublishDate",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetPreviewImages() {
	_jsii_.InvokeVoid(
		s,
		"resetPreviewImages",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetPreviewImagesDark() {
	_jsii_.InvokeVoid(
		s,
		"resetPreviewImagesDark",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetProviders() {
	_jsii_.InvokeVoid(
		s,
		"resetProviders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetSource() {
	_jsii_.InvokeVoid(
		s,
		"resetSource",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetSupport() {
	_jsii_.InvokeVoid(
		s,
		"resetSupport",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetThreatAnalysisTactics() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatAnalysisTactics",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetThreatAnalysisTechniques() {
	_jsii_.InvokeVoid(
		s,
		"resetThreatAnalysisTechniques",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) ResetVersion() {
	_jsii_.InvokeVoid(
		s,
		"resetVersion",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SentinelMetadata) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelMetadata) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelMetadata) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelMetadata) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelMetadata) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SentinelMetadata) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

