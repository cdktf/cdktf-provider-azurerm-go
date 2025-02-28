// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetdelimitedtext

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/datafactorydatasetdelimitedtext/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_delimited_text azurerm_data_factory_dataset_delimited_text}.
type DataFactoryDatasetDelimitedText interface {
	cdktf.TerraformResource
	AdditionalProperties() *map[string]*string
	SetAdditionalProperties(val *map[string]*string)
	AdditionalPropertiesInput() *map[string]*string
	Annotations() *[]*string
	SetAnnotations(val *[]*string)
	AnnotationsInput() *[]*string
	AzureBlobFsLocation() DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference
	AzureBlobFsLocationInput() *DataFactoryDatasetDelimitedTextAzureBlobFsLocation
	AzureBlobStorageLocation() DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference
	AzureBlobStorageLocationInput() *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ColumnDelimiter() *string
	SetColumnDelimiter(val *string)
	ColumnDelimiterInput() *string
	CompressionCodec() *string
	SetCompressionCodec(val *string)
	CompressionCodecInput() *string
	CompressionLevel() *string
	SetCompressionLevel(val *string)
	CompressionLevelInput() *string
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
	DataFactoryId() *string
	SetDataFactoryId(val *string)
	DataFactoryIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Encoding() *string
	SetEncoding(val *string)
	EncodingInput() *string
	EscapeCharacter() *string
	SetEscapeCharacter(val *string)
	EscapeCharacterInput() *string
	FirstRowAsHeader() interface{}
	SetFirstRowAsHeader(val interface{})
	FirstRowAsHeaderInput() interface{}
	Folder() *string
	SetFolder(val *string)
	FolderInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpServerLocation() DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference
	HttpServerLocationInput() *DataFactoryDatasetDelimitedTextHttpServerLocation
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LinkedServiceName() *string
	SetLinkedServiceName(val *string)
	LinkedServiceNameInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NullValue() *string
	SetNullValue(val *string)
	NullValueInput() *string
	Parameters() *map[string]*string
	SetParameters(val *map[string]*string)
	ParametersInput() *map[string]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QuoteCharacter() *string
	SetQuoteCharacter(val *string)
	QuoteCharacterInput() *string
	// Experimental.
	RawOverrides() interface{}
	RowDelimiter() *string
	SetRowDelimiter(val *string)
	RowDelimiterInput() *string
	SchemaColumn() DataFactoryDatasetDelimitedTextSchemaColumnList
	SchemaColumnInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryDatasetDelimitedTextTimeoutsOutputReference
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
	PutAzureBlobFsLocation(value *DataFactoryDatasetDelimitedTextAzureBlobFsLocation)
	PutAzureBlobStorageLocation(value *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation)
	PutHttpServerLocation(value *DataFactoryDatasetDelimitedTextHttpServerLocation)
	PutSchemaColumn(value interface{})
	PutTimeouts(value *DataFactoryDatasetDelimitedTextTimeouts)
	ResetAdditionalProperties()
	ResetAnnotations()
	ResetAzureBlobFsLocation()
	ResetAzureBlobStorageLocation()
	ResetColumnDelimiter()
	ResetCompressionCodec()
	ResetCompressionLevel()
	ResetDescription()
	ResetEncoding()
	ResetEscapeCharacter()
	ResetFirstRowAsHeader()
	ResetFolder()
	ResetHttpServerLocation()
	ResetId()
	ResetNullValue()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetParameters()
	ResetQuoteCharacter()
	ResetRowDelimiter()
	ResetSchemaColumn()
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

// The jsii proxy struct for DataFactoryDatasetDelimitedText
type jsiiProxy_DataFactoryDatasetDelimitedText struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AdditionalProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AdditionalPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Annotations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AnnotationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"annotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AzureBlobFsLocation() DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference {
	var returns DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference
	_jsii_.Get(
		j,
		"azureBlobFsLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AzureBlobFsLocationInput() *DataFactoryDatasetDelimitedTextAzureBlobFsLocation {
	var returns *DataFactoryDatasetDelimitedTextAzureBlobFsLocation
	_jsii_.Get(
		j,
		"azureBlobFsLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AzureBlobStorageLocation() DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference {
	var returns DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference
	_jsii_.Get(
		j,
		"azureBlobStorageLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) AzureBlobStorageLocationInput() *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation {
	var returns *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation
	_jsii_.Get(
		j,
		"azureBlobStorageLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) ColumnDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) ColumnDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"columnDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) CompressionCodec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionCodec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) CompressionCodecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionCodecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) CompressionLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) CompressionLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Encoding() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encoding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) EncodingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) EscapeCharacter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeCharacter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) EscapeCharacterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"escapeCharacterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) FirstRowAsHeader() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstRowAsHeader",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) FirstRowAsHeaderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"firstRowAsHeaderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Folder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) FolderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"folderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) HttpServerLocation() DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference {
	var returns DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference
	_jsii_.Get(
		j,
		"httpServerLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) HttpServerLocationInput() *DataFactoryDatasetDelimitedTextHttpServerLocation {
	var returns *DataFactoryDatasetDelimitedTextHttpServerLocation
	_jsii_.Get(
		j,
		"httpServerLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) LinkedServiceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) LinkedServiceNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linkedServiceNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) NullValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) NullValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nullValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Parameters() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) ParametersInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"parametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) QuoteCharacter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteCharacter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) QuoteCharacterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"quoteCharacterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) RowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) RowDelimiterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rowDelimiterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SchemaColumn() DataFactoryDatasetDelimitedTextSchemaColumnList {
	var returns DataFactoryDatasetDelimitedTextSchemaColumnList
	_jsii_.Get(
		j,
		"schemaColumn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SchemaColumnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"schemaColumnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Timeouts() DataFactoryDatasetDelimitedTextTimeoutsOutputReference {
	var returns DataFactoryDatasetDelimitedTextTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_delimited_text azurerm_data_factory_dataset_delimited_text} Resource.
func NewDataFactoryDatasetDelimitedText(scope constructs.Construct, id *string, config *DataFactoryDatasetDelimitedTextConfig) DataFactoryDatasetDelimitedText {
	_init_.Initialize()

	if err := validateNewDataFactoryDatasetDelimitedTextParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryDatasetDelimitedText{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/data_factory_dataset_delimited_text azurerm_data_factory_dataset_delimited_text} Resource.
func NewDataFactoryDatasetDelimitedText_Override(d DataFactoryDatasetDelimitedText, scope constructs.Construct, id *string, config *DataFactoryDatasetDelimitedTextConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetAdditionalProperties(val *map[string]*string) {
	if err := j.validateSetAdditionalPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetAnnotations(val *[]*string) {
	if err := j.validateSetAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetColumnDelimiter(val *string) {
	if err := j.validateSetColumnDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"columnDelimiter",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetCompressionCodec(val *string) {
	if err := j.validateSetCompressionCodecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionCodec",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetCompressionLevel(val *string) {
	if err := j.validateSetCompressionLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionLevel",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetDataFactoryId(val *string) {
	if err := j.validateSetDataFactoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetEncoding(val *string) {
	if err := j.validateSetEncodingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetEscapeCharacter(val *string) {
	if err := j.validateSetEscapeCharacterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"escapeCharacter",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetFirstRowAsHeader(val interface{}) {
	if err := j.validateSetFirstRowAsHeaderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"firstRowAsHeader",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetFolder(val *string) {
	if err := j.validateSetFolderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetLinkedServiceName(val *string) {
	if err := j.validateSetLinkedServiceNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linkedServiceName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetNullValue(val *string) {
	if err := j.validateSetNullValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nullValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetParameters(val *map[string]*string) {
	if err := j.validateSetParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetQuoteCharacter(val *string) {
	if err := j.validateSetQuoteCharacterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"quoteCharacter",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText)SetRowDelimiter(val *string) {
	if err := j.validateSetRowDelimiterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rowDelimiter",
		val,
	)
}

// Generates CDKTF code for importing a DataFactoryDatasetDelimitedText resource upon running "cdktf plan <stack-name>".
func DataFactoryDatasetDelimitedText_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateDataFactoryDatasetDelimitedText_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
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
func DataFactoryDatasetDelimitedText_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryDatasetDelimitedText_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryDatasetDelimitedText_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryDatasetDelimitedText_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryDatasetDelimitedText_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryDatasetDelimitedText_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryDatasetDelimitedText_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) AddMoveTarget(moveTarget *string) {
	if err := d.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := d.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) MoveFromId(id *string) {
	if err := d.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveFromId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) MoveTo(moveTarget *string, index interface{}) {
	if err := d.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) MoveToId(id *string) {
	if err := d.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"moveToId",
		[]interface{}{id},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutAzureBlobFsLocation(value *DataFactoryDatasetDelimitedTextAzureBlobFsLocation) {
	if err := d.validatePutAzureBlobFsLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureBlobFsLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutAzureBlobStorageLocation(value *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation) {
	if err := d.validatePutAzureBlobStorageLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putAzureBlobStorageLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutHttpServerLocation(value *DataFactoryDatasetDelimitedTextHttpServerLocation) {
	if err := d.validatePutHttpServerLocationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHttpServerLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutSchemaColumn(value interface{}) {
	if err := d.validatePutSchemaColumnParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSchemaColumn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutTimeouts(value *DataFactoryDatasetDelimitedTextTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetAdditionalProperties() {
	_jsii_.InvokeVoid(
		d,
		"resetAdditionalProperties",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetAnnotations() {
	_jsii_.InvokeVoid(
		d,
		"resetAnnotations",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetAzureBlobFsLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureBlobFsLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetAzureBlobStorageLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetAzureBlobStorageLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetColumnDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetColumnDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetCompressionCodec() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionCodec",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetCompressionLevel() {
	_jsii_.InvokeVoid(
		d,
		"resetCompressionLevel",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetEncoding() {
	_jsii_.InvokeVoid(
		d,
		"resetEncoding",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetEscapeCharacter() {
	_jsii_.InvokeVoid(
		d,
		"resetEscapeCharacter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetFirstRowAsHeader() {
	_jsii_.InvokeVoid(
		d,
		"resetFirstRowAsHeader",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetFolder() {
	_jsii_.InvokeVoid(
		d,
		"resetFolder",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetHttpServerLocation() {
	_jsii_.InvokeVoid(
		d,
		"resetHttpServerLocation",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetNullValue() {
	_jsii_.InvokeVoid(
		d,
		"resetNullValue",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetParameters() {
	_jsii_.InvokeVoid(
		d,
		"resetParameters",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetQuoteCharacter() {
	_jsii_.InvokeVoid(
		d,
		"resetQuoteCharacter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetRowDelimiter() {
	_jsii_.InvokeVoid(
		d,
		"resetRowDelimiter",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetSchemaColumn() {
	_jsii_.InvokeVoid(
		d,
		"resetSchemaColumn",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

