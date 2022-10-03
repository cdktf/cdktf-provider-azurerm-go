package datafactorydatasetdelimitedtext

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/datafactorydatasetdelimitedtext/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text azurerm_data_factory_dataset_delimited_text}.
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) Count() *float64 {
	var returns *float64
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text azurerm_data_factory_dataset_delimited_text} Resource.
func NewDataFactoryDatasetDelimitedText(scope constructs.Construct, id *string, config *DataFactoryDatasetDelimitedTextConfig) DataFactoryDatasetDelimitedText {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedText{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text azurerm_data_factory_dataset_delimited_text} Resource.
func NewDataFactoryDatasetDelimitedText_Override(d DataFactoryDatasetDelimitedText, scope constructs.Construct, id *string, config *DataFactoryDatasetDelimitedTextConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetAdditionalProperties(val *map[string]*string) {
	_jsii_.Set(
		j,
		"additionalProperties",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetAnnotations(val *[]*string) {
	_jsii_.Set(
		j,
		"annotations",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetColumnDelimiter(val *string) {
	_jsii_.Set(
		j,
		"columnDelimiter",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetCompressionCodec(val *string) {
	_jsii_.Set(
		j,
		"compressionCodec",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetCompressionLevel(val *string) {
	_jsii_.Set(
		j,
		"compressionLevel",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetDataFactoryId(val *string) {
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetEncoding(val *string) {
	_jsii_.Set(
		j,
		"encoding",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetEscapeCharacter(val *string) {
	_jsii_.Set(
		j,
		"escapeCharacter",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetFirstRowAsHeader(val interface{}) {
	_jsii_.Set(
		j,
		"firstRowAsHeader",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetFolder(val *string) {
	_jsii_.Set(
		j,
		"folder",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetLinkedServiceName(val *string) {
	_jsii_.Set(
		j,
		"linkedServiceName",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetNullValue(val *string) {
	_jsii_.Set(
		j,
		"nullValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetParameters(val *map[string]*string) {
	_jsii_.Set(
		j,
		"parameters",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetQuoteCharacter(val *string) {
	_jsii_.Set(
		j,
		"quoteCharacter",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedText) SetRowDelimiter(val *string) {
	_jsii_.Set(
		j,
		"rowDelimiter",
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
func DataFactoryDatasetDelimitedText_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedText",
		"isConstruct",
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

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutAzureBlobFsLocation(value *DataFactoryDatasetDelimitedTextAzureBlobFsLocation) {
	_jsii_.InvokeVoid(
		d,
		"putAzureBlobFsLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutAzureBlobStorageLocation(value *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation) {
	_jsii_.InvokeVoid(
		d,
		"putAzureBlobStorageLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutHttpServerLocation(value *DataFactoryDatasetDelimitedTextHttpServerLocation) {
	_jsii_.InvokeVoid(
		d,
		"putHttpServerLocation",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutSchemaColumn(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putSchemaColumn",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedText) PutTimeouts(value *DataFactoryDatasetDelimitedTextTimeouts) {
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

type DataFactoryDatasetDelimitedTextAzureBlobFsLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#file_system DataFactoryDatasetDelimitedText#file_system}.
	FileSystem *string `field:"required" json:"fileSystem" yaml:"fileSystem"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#filename DataFactoryDatasetDelimitedText#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#path DataFactoryDatasetDelimitedText#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

type DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference interface {
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
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	FileSystem() *string
	SetFileSystem(val *string)
	FileSystemInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataFactoryDatasetDelimitedTextAzureBlobFsLocation
	SetInternalValue(val *DataFactoryDatasetDelimitedTextAzureBlobFsLocation)
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ResetFilename()
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference
type jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) FileSystem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) FileSystemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) InternalValue() *DataFactoryDatasetDelimitedTextAzureBlobFsLocation {
	var returns *DataFactoryDatasetDelimitedTextAzureBlobFsLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference_Override(d DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetFileSystem(val *string) {
	_jsii_.Set(
		j,
		"fileSystem",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetInternalValue(val *DataFactoryDatasetDelimitedTextAzureBlobFsLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) ResetFilename() {
	_jsii_.InvokeVoid(
		d,
		"resetFilename",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		d,
		"resetPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobFsLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryDatasetDelimitedTextAzureBlobStorageLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#container DataFactoryDatasetDelimitedText#container}.
	Container *string `field:"required" json:"container" yaml:"container"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#dynamic_container_enabled DataFactoryDatasetDelimitedText#dynamic_container_enabled}.
	DynamicContainerEnabled interface{} `field:"optional" json:"dynamicContainerEnabled" yaml:"dynamicContainerEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#dynamic_filename_enabled DataFactoryDatasetDelimitedText#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#dynamic_path_enabled DataFactoryDatasetDelimitedText#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#filename DataFactoryDatasetDelimitedText#filename}.
	Filename *string `field:"optional" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#path DataFactoryDatasetDelimitedText#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

type DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference interface {
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
	Container() *string
	SetContainer(val *string)
	ContainerInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DynamicContainerEnabled() interface{}
	SetDynamicContainerEnabled(val interface{})
	DynamicContainerEnabledInput() interface{}
	DynamicFilenameEnabled() interface{}
	SetDynamicFilenameEnabled(val interface{})
	DynamicFilenameEnabledInput() interface{}
	DynamicPathEnabled() interface{}
	SetDynamicPathEnabled(val interface{})
	DynamicPathEnabledInput() interface{}
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation
	SetInternalValue(val *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation)
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ResetDynamicContainerEnabled()
	ResetDynamicFilenameEnabled()
	ResetDynamicPathEnabled()
	ResetFilename()
	ResetPath()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference
type jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) Container() *string {
	var returns *string
	_jsii_.Get(
		j,
		"container",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ContainerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) DynamicContainerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicContainerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) DynamicContainerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicContainerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) DynamicFilenameEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) DynamicFilenameEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) DynamicPathEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) DynamicPathEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) InternalValue() *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation {
	var returns *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference_Override(d DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetContainer(val *string) {
	_jsii_.Set(
		j,
		"container",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetDynamicContainerEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dynamicContainerEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetDynamicFilenameEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dynamicFilenameEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetDynamicPathEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dynamicPathEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetInternalValue(val *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ResetDynamicContainerEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicContainerEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ResetDynamicFilenameEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicFilenameEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ResetDynamicPathEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicPathEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ResetFilename() {
	_jsii_.InvokeVoid(
		d,
		"resetFilename",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		d,
		"resetPath",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextAzureBlobStorageLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryDatasetDelimitedTextConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#data_factory_id DataFactoryDatasetDelimitedText#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#linked_service_name DataFactoryDatasetDelimitedText#linked_service_name}.
	LinkedServiceName *string `field:"required" json:"linkedServiceName" yaml:"linkedServiceName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#name DataFactoryDatasetDelimitedText#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#additional_properties DataFactoryDatasetDelimitedText#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#annotations DataFactoryDatasetDelimitedText#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// azure_blob_fs_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#azure_blob_fs_location DataFactoryDatasetDelimitedText#azure_blob_fs_location}
	AzureBlobFsLocation *DataFactoryDatasetDelimitedTextAzureBlobFsLocation `field:"optional" json:"azureBlobFsLocation" yaml:"azureBlobFsLocation"`
	// azure_blob_storage_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#azure_blob_storage_location DataFactoryDatasetDelimitedText#azure_blob_storage_location}
	AzureBlobStorageLocation *DataFactoryDatasetDelimitedTextAzureBlobStorageLocation `field:"optional" json:"azureBlobStorageLocation" yaml:"azureBlobStorageLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#column_delimiter DataFactoryDatasetDelimitedText#column_delimiter}.
	ColumnDelimiter *string `field:"optional" json:"columnDelimiter" yaml:"columnDelimiter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#compression_codec DataFactoryDatasetDelimitedText#compression_codec}.
	CompressionCodec *string `field:"optional" json:"compressionCodec" yaml:"compressionCodec"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#compression_level DataFactoryDatasetDelimitedText#compression_level}.
	CompressionLevel *string `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#description DataFactoryDatasetDelimitedText#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#encoding DataFactoryDatasetDelimitedText#encoding}.
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#escape_character DataFactoryDatasetDelimitedText#escape_character}.
	EscapeCharacter *string `field:"optional" json:"escapeCharacter" yaml:"escapeCharacter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#first_row_as_header DataFactoryDatasetDelimitedText#first_row_as_header}.
	FirstRowAsHeader interface{} `field:"optional" json:"firstRowAsHeader" yaml:"firstRowAsHeader"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#folder DataFactoryDatasetDelimitedText#folder}.
	Folder *string `field:"optional" json:"folder" yaml:"folder"`
	// http_server_location block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#http_server_location DataFactoryDatasetDelimitedText#http_server_location}
	HttpServerLocation *DataFactoryDatasetDelimitedTextHttpServerLocation `field:"optional" json:"httpServerLocation" yaml:"httpServerLocation"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#id DataFactoryDatasetDelimitedText#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#null_value DataFactoryDatasetDelimitedText#null_value}.
	NullValue *string `field:"optional" json:"nullValue" yaml:"nullValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#parameters DataFactoryDatasetDelimitedText#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#quote_character DataFactoryDatasetDelimitedText#quote_character}.
	QuoteCharacter *string `field:"optional" json:"quoteCharacter" yaml:"quoteCharacter"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#row_delimiter DataFactoryDatasetDelimitedText#row_delimiter}.
	RowDelimiter *string `field:"optional" json:"rowDelimiter" yaml:"rowDelimiter"`
	// schema_column block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#schema_column DataFactoryDatasetDelimitedText#schema_column}
	SchemaColumn interface{} `field:"optional" json:"schemaColumn" yaml:"schemaColumn"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#timeouts DataFactoryDatasetDelimitedText#timeouts}
	Timeouts *DataFactoryDatasetDelimitedTextTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

type DataFactoryDatasetDelimitedTextHttpServerLocation struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#filename DataFactoryDatasetDelimitedText#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#path DataFactoryDatasetDelimitedText#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#relative_url DataFactoryDatasetDelimitedText#relative_url}.
	RelativeUrl *string `field:"required" json:"relativeUrl" yaml:"relativeUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#dynamic_filename_enabled DataFactoryDatasetDelimitedText#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#dynamic_path_enabled DataFactoryDatasetDelimitedText#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
}

type DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference interface {
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
	DynamicFilenameEnabled() interface{}
	SetDynamicFilenameEnabled(val interface{})
	DynamicFilenameEnabledInput() interface{}
	DynamicPathEnabled() interface{}
	SetDynamicPathEnabled(val interface{})
	DynamicPathEnabledInput() interface{}
	Filename() *string
	SetFilename(val *string)
	FilenameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *DataFactoryDatasetDelimitedTextHttpServerLocation
	SetInternalValue(val *DataFactoryDatasetDelimitedTextHttpServerLocation)
	Path() *string
	SetPath(val *string)
	PathInput() *string
	RelativeUrl() *string
	SetRelativeUrl(val *string)
	RelativeUrlInput() *string
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
	ResetDynamicFilenameEnabled()
	ResetDynamicPathEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference
type jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) DynamicFilenameEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) DynamicFilenameEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicFilenameEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) DynamicPathEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) DynamicPathEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dynamicPathEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) Filename() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filename",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) FilenameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filenameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) InternalValue() *DataFactoryDatasetDelimitedTextHttpServerLocation {
	var returns *DataFactoryDatasetDelimitedTextHttpServerLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) RelativeUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) RelativeUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"relativeUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetDelimitedTextHttpServerLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetDelimitedTextHttpServerLocationOutputReference_Override(d DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetDynamicFilenameEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dynamicFilenameEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetDynamicPathEnabled(val interface{}) {
	_jsii_.Set(
		j,
		"dynamicPathEnabled",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetFilename(val *string) {
	_jsii_.Set(
		j,
		"filename",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetInternalValue(val *DataFactoryDatasetDelimitedTextHttpServerLocation) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetRelativeUrl(val *string) {
	_jsii_.Set(
		j,
		"relativeUrl",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) ResetDynamicFilenameEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicFilenameEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) ResetDynamicPathEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetDynamicPathEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextHttpServerLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryDatasetDelimitedTextSchemaColumn struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#name DataFactoryDatasetDelimitedText#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#description DataFactoryDatasetDelimitedText#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#type DataFactoryDatasetDelimitedText#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

type DataFactoryDatasetDelimitedTextSchemaColumnList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) DataFactoryDatasetDelimitedTextSchemaColumnOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetDelimitedTextSchemaColumnList
type jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetDelimitedTextSchemaColumnList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) DataFactoryDatasetDelimitedTextSchemaColumnList {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextSchemaColumnList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetDelimitedTextSchemaColumnList_Override(d DataFactoryDatasetDelimitedTextSchemaColumnList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextSchemaColumnList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) Get(index *float64) DataFactoryDatasetDelimitedTextSchemaColumnOutputReference {
	var returns DataFactoryDatasetDelimitedTextSchemaColumnOutputReference

	_jsii_.Invoke(
		d,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryDatasetDelimitedTextSchemaColumnOutputReference interface {
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	ResetDescription()
	ResetType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataFactoryDatasetDelimitedTextSchemaColumnOutputReference
type jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetDelimitedTextSchemaColumnOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataFactoryDatasetDelimitedTextSchemaColumnOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextSchemaColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetDelimitedTextSchemaColumnOutputReference_Override(d DataFactoryDatasetDelimitedTextSchemaColumnOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextSchemaColumnOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) ResetType() {
	_jsii_.InvokeVoid(
		d,
		"resetType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextSchemaColumnOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

type DataFactoryDatasetDelimitedTextTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#create DataFactoryDatasetDelimitedText#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#delete DataFactoryDatasetDelimitedText#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#read DataFactoryDatasetDelimitedText#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/data_factory_dataset_delimited_text#update DataFactoryDatasetDelimitedText#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

type DataFactoryDatasetDelimitedTextTimeoutsOutputReference interface {
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

// The jsii proxy struct for DataFactoryDatasetDelimitedTextTimeoutsOutputReference
type jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) Create() *string {
	var returns *string
	_jsii_.Get(
		j,
		"create",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) CreateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) Delete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"delete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) DeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) Read() *string {
	var returns *string
	_jsii_.Get(
		j,
		"read",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ReadInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) Update() *string {
	var returns *string
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) UpdateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewDataFactoryDatasetDelimitedTextTimeoutsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataFactoryDatasetDelimitedTextTimeoutsOutputReference {
	_init_.Initialize()

	j := jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataFactoryDatasetDelimitedTextTimeoutsOutputReference_Override(d DataFactoryDatasetDelimitedTextTimeoutsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryDatasetDelimitedText.DataFactoryDatasetDelimitedTextTimeoutsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetCreate(val *string) {
	_jsii_.Set(
		j,
		"create",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetDelete(val *string) {
	_jsii_.Set(
		j,
		"delete",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetRead(val *string) {
	_jsii_.Set(
		j,
		"read",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) SetUpdate(val *string) {
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ResetCreate() {
	_jsii_.InvokeVoid(
		d,
		"resetCreate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ResetDelete() {
	_jsii_.InvokeVoid(
		d,
		"resetDelete",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ResetRead() {
	_jsii_.InvokeVoid(
		d,
		"resetRead",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		d,
		"resetUpdate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryDatasetDelimitedTextTimeoutsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

