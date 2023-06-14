package datafactoryintegrationruntimemanaged

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/datafactoryintegrationruntimemanaged/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_integration_runtime_managed azurerm_data_factory_integration_runtime_managed}.
type DataFactoryIntegrationRuntimeManaged interface {
	cdktf.TerraformResource
	CatalogInfo() DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference
	CatalogInfoInput() *DataFactoryIntegrationRuntimeManagedCatalogInfo
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
	CustomSetupScript() DataFactoryIntegrationRuntimeManagedCustomSetupScriptOutputReference
	CustomSetupScriptInput() *DataFactoryIntegrationRuntimeManagedCustomSetupScript
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
	Edition() *string
	SetEdition(val *string)
	EditionInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxParallelExecutionsPerNode() *float64
	SetMaxParallelExecutionsPerNode(val *float64)
	MaxParallelExecutionsPerNodeInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeSize() *string
	SetNodeSize(val *string)
	NodeSizeInput() *string
	NumberOfNodes() *float64
	SetNumberOfNodes(val *float64)
	NumberOfNodesInput() *float64
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() DataFactoryIntegrationRuntimeManagedTimeoutsOutputReference
	TimeoutsInput() interface{}
	VnetIntegration() DataFactoryIntegrationRuntimeManagedVnetIntegrationOutputReference
	VnetIntegrationInput() *DataFactoryIntegrationRuntimeManagedVnetIntegration
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
	PutCatalogInfo(value *DataFactoryIntegrationRuntimeManagedCatalogInfo)
	PutCustomSetupScript(value *DataFactoryIntegrationRuntimeManagedCustomSetupScript)
	PutTimeouts(value *DataFactoryIntegrationRuntimeManagedTimeouts)
	PutVnetIntegration(value *DataFactoryIntegrationRuntimeManagedVnetIntegration)
	ResetCatalogInfo()
	ResetCustomSetupScript()
	ResetDescription()
	ResetEdition()
	ResetId()
	ResetLicenseType()
	ResetMaxParallelExecutionsPerNode()
	ResetNumberOfNodes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeouts()
	ResetVnetIntegration()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataFactoryIntegrationRuntimeManaged
type jsiiProxy_DataFactoryIntegrationRuntimeManaged struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) CatalogInfo() DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference {
	var returns DataFactoryIntegrationRuntimeManagedCatalogInfoOutputReference
	_jsii_.Get(
		j,
		"catalogInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) CatalogInfoInput() *DataFactoryIntegrationRuntimeManagedCatalogInfo {
	var returns *DataFactoryIntegrationRuntimeManagedCatalogInfo
	_jsii_.Get(
		j,
		"catalogInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) CustomSetupScript() DataFactoryIntegrationRuntimeManagedCustomSetupScriptOutputReference {
	var returns DataFactoryIntegrationRuntimeManagedCustomSetupScriptOutputReference
	_jsii_.Get(
		j,
		"customSetupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) CustomSetupScriptInput() *DataFactoryIntegrationRuntimeManagedCustomSetupScript {
	var returns *DataFactoryIntegrationRuntimeManagedCustomSetupScript
	_jsii_.Get(
		j,
		"customSetupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) DataFactoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) DataFactoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFactoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Edition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) EditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"editionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) MaxParallelExecutionsPerNode() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelExecutionsPerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) MaxParallelExecutionsPerNodeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxParallelExecutionsPerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) NodeSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) NodeSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) NumberOfNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) NumberOfNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numberOfNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) Timeouts() DataFactoryIntegrationRuntimeManagedTimeoutsOutputReference {
	var returns DataFactoryIntegrationRuntimeManagedTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) VnetIntegration() DataFactoryIntegrationRuntimeManagedVnetIntegrationOutputReference {
	var returns DataFactoryIntegrationRuntimeManagedVnetIntegrationOutputReference
	_jsii_.Get(
		j,
		"vnetIntegration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged) VnetIntegrationInput() *DataFactoryIntegrationRuntimeManagedVnetIntegration {
	var returns *DataFactoryIntegrationRuntimeManagedVnetIntegration
	_jsii_.Get(
		j,
		"vnetIntegrationInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_integration_runtime_managed azurerm_data_factory_integration_runtime_managed} Resource.
func NewDataFactoryIntegrationRuntimeManaged(scope constructs.Construct, id *string, config *DataFactoryIntegrationRuntimeManagedConfig) DataFactoryIntegrationRuntimeManaged {
	_init_.Initialize()

	if err := validateNewDataFactoryIntegrationRuntimeManagedParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataFactoryIntegrationRuntimeManaged{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManaged",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/data_factory_integration_runtime_managed azurerm_data_factory_integration_runtime_managed} Resource.
func NewDataFactoryIntegrationRuntimeManaged_Override(d DataFactoryIntegrationRuntimeManaged, scope constructs.Construct, id *string, config *DataFactoryIntegrationRuntimeManagedConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManaged",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetDataFactoryId(val *string) {
	if err := j.validateSetDataFactoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFactoryId",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetEdition(val *string) {
	if err := j.validateSetEditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edition",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetMaxParallelExecutionsPerNode(val *float64) {
	if err := j.validateSetMaxParallelExecutionsPerNodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxParallelExecutionsPerNode",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetNodeSize(val *string) {
	if err := j.validateSetNodeSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nodeSize",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetNumberOfNodes(val *float64) {
	if err := j.validateSetNumberOfNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numberOfNodes",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataFactoryIntegrationRuntimeManaged)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func DataFactoryIntegrationRuntimeManaged_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeManaged_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManaged",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryIntegrationRuntimeManaged_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeManaged_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManaged",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataFactoryIntegrationRuntimeManaged_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataFactoryIntegrationRuntimeManaged_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManaged",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataFactoryIntegrationRuntimeManaged_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.dataFactoryIntegrationRuntimeManaged.DataFactoryIntegrationRuntimeManaged",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) PutCatalogInfo(value *DataFactoryIntegrationRuntimeManagedCatalogInfo) {
	if err := d.validatePutCatalogInfoParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCatalogInfo",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) PutCustomSetupScript(value *DataFactoryIntegrationRuntimeManagedCustomSetupScript) {
	if err := d.validatePutCustomSetupScriptParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCustomSetupScript",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) PutTimeouts(value *DataFactoryIntegrationRuntimeManagedTimeouts) {
	if err := d.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) PutVnetIntegration(value *DataFactoryIntegrationRuntimeManagedVnetIntegration) {
	if err := d.validatePutVnetIntegrationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putVnetIntegration",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetCatalogInfo() {
	_jsii_.InvokeVoid(
		d,
		"resetCatalogInfo",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetCustomSetupScript() {
	_jsii_.InvokeVoid(
		d,
		"resetCustomSetupScript",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetEdition() {
	_jsii_.InvokeVoid(
		d,
		"resetEdition",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetLicenseType() {
	_jsii_.InvokeVoid(
		d,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetMaxParallelExecutionsPerNode() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxParallelExecutionsPerNode",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetNumberOfNodes() {
	_jsii_.InvokeVoid(
		d,
		"resetNumberOfNodes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetTimeouts() {
	_jsii_.InvokeVoid(
		d,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ResetVnetIntegration() {
	_jsii_.InvokeVoid(
		d,
		"resetVnetIntegration",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataFactoryIntegrationRuntimeManaged) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

