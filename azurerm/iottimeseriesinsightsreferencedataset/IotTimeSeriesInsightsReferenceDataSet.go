// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iottimeseriesinsightsreferencedataset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/iottimeseriesinsightsreferencedataset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iot_time_series_insights_reference_data_set azurerm_iot_time_series_insights_reference_data_set}.
type IotTimeSeriesInsightsReferenceDataSet interface {
	cdktf.TerraformResource
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
	DataStringComparisonBehavior() *string
	SetDataStringComparisonBehavior(val *string)
	DataStringComparisonBehaviorInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
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
	KeyProperty() IotTimeSeriesInsightsReferenceDataSetKeyPropertyList
	KeyPropertyInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IotTimeSeriesInsightsReferenceDataSetTimeoutsOutputReference
	TimeoutsInput() interface{}
	TimeSeriesInsightsEnvironmentId() *string
	SetTimeSeriesInsightsEnvironmentId(val *string)
	TimeSeriesInsightsEnvironmentIdInput() *string
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutKeyProperty(value interface{})
	PutTimeouts(value *IotTimeSeriesInsightsReferenceDataSetTimeouts)
	ResetDataStringComparisonBehavior()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTags()
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

// The jsii proxy struct for IotTimeSeriesInsightsReferenceDataSet
type jsiiProxy_IotTimeSeriesInsightsReferenceDataSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) DataStringComparisonBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStringComparisonBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) DataStringComparisonBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStringComparisonBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) KeyProperty() IotTimeSeriesInsightsReferenceDataSetKeyPropertyList {
	var returns IotTimeSeriesInsightsReferenceDataSetKeyPropertyList
	_jsii_.Get(
		j,
		"keyProperty",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) KeyPropertyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"keyPropertyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) Timeouts() IotTimeSeriesInsightsReferenceDataSetTimeoutsOutputReference {
	var returns IotTimeSeriesInsightsReferenceDataSetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TimeSeriesInsightsEnvironmentId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeSeriesInsightsEnvironmentId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) TimeSeriesInsightsEnvironmentIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeSeriesInsightsEnvironmentIdInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iot_time_series_insights_reference_data_set azurerm_iot_time_series_insights_reference_data_set} Resource.
func NewIotTimeSeriesInsightsReferenceDataSet(scope constructs.Construct, id *string, config *IotTimeSeriesInsightsReferenceDataSetConfig) IotTimeSeriesInsightsReferenceDataSet {
	_init_.Initialize()

	if err := validateNewIotTimeSeriesInsightsReferenceDataSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotTimeSeriesInsightsReferenceDataSet{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/iot_time_series_insights_reference_data_set azurerm_iot_time_series_insights_reference_data_set} Resource.
func NewIotTimeSeriesInsightsReferenceDataSet_Override(i IotTimeSeriesInsightsReferenceDataSet, scope constructs.Construct, id *string, config *IotTimeSeriesInsightsReferenceDataSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetDataStringComparisonBehavior(val *string) {
	if err := j.validateSetDataStringComparisonBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataStringComparisonBehavior",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet)SetTimeSeriesInsightsEnvironmentId(val *string) {
	if err := j.validateSetTimeSeriesInsightsEnvironmentIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeSeriesInsightsEnvironmentId",
		val,
	)
}

// Generates CDKTF code for importing a IotTimeSeriesInsightsReferenceDataSet resource upon running "cdktf plan <stack-name>".
func IotTimeSeriesInsightsReferenceDataSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateIotTimeSeriesInsightsReferenceDataSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
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
func IotTimeSeriesInsightsReferenceDataSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotTimeSeriesInsightsReferenceDataSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotTimeSeriesInsightsReferenceDataSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotTimeSeriesInsightsReferenceDataSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IotTimeSeriesInsightsReferenceDataSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIotTimeSeriesInsightsReferenceDataSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IotTimeSeriesInsightsReferenceDataSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.iotTimeSeriesInsightsReferenceDataSet.IotTimeSeriesInsightsReferenceDataSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) AddMoveTarget(moveTarget *string) {
	if err := i.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := i.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) MoveTo(moveTarget *string, index interface{}) {
	if err := i.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) PutKeyProperty(value interface{}) {
	if err := i.validatePutKeyPropertyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putKeyProperty",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) PutTimeouts(value *IotTimeSeriesInsightsReferenceDataSetTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ResetDataStringComparisonBehavior() {
	_jsii_.InvokeVoid(
		i,
		"resetDataStringComparisonBehavior",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotTimeSeriesInsightsReferenceDataSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

