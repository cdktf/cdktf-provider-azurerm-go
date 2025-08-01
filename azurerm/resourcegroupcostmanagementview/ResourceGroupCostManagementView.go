// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/resourcegroupcostmanagementview/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/resource_group_cost_management_view azurerm_resource_group_cost_management_view}.
type ResourceGroupCostManagementView interface {
	cdktf.TerraformResource
	Accumulated() interface{}
	SetAccumulated(val interface{})
	AccumulatedInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ChartType() *string
	SetChartType(val *string)
	ChartTypeInput() *string
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
	Dataset() ResourceGroupCostManagementViewDatasetOutputReference
	DatasetInput() *ResourceGroupCostManagementViewDataset
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisplayName() *string
	SetDisplayName(val *string)
	DisplayNameInput() *string
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
	Kpi() ResourceGroupCostManagementViewKpiList
	KpiInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	Pivot() ResourceGroupCostManagementViewPivotList
	PivotInput() interface{}
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
	ReportType() *string
	SetReportType(val *string)
	ReportTypeInput() *string
	ResourceGroupId() *string
	SetResourceGroupId(val *string)
	ResourceGroupIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeframe() *string
	SetTimeframe(val *string)
	TimeframeInput() *string
	Timeouts() ResourceGroupCostManagementViewTimeoutsOutputReference
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
	PutDataset(value *ResourceGroupCostManagementViewDataset)
	PutKpi(value interface{})
	PutPivot(value interface{})
	PutTimeouts(value *ResourceGroupCostManagementViewTimeouts)
	ResetId()
	ResetKpi()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPivot()
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

// The jsii proxy struct for ResourceGroupCostManagementView
type jsiiProxy_ResourceGroupCostManagementView struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Accumulated() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accumulated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) AccumulatedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"accumulatedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ChartType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chartType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ChartTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"chartTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Dataset() ResourceGroupCostManagementViewDatasetOutputReference {
	var returns ResourceGroupCostManagementViewDatasetOutputReference
	_jsii_.Get(
		j,
		"dataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) DatasetInput() *ResourceGroupCostManagementViewDataset {
	var returns *ResourceGroupCostManagementViewDataset
	_jsii_.Get(
		j,
		"datasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) DisplayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) DisplayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"displayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Kpi() ResourceGroupCostManagementViewKpiList {
	var returns ResourceGroupCostManagementViewKpiList
	_jsii_.Get(
		j,
		"kpi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) KpiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"kpiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Pivot() ResourceGroupCostManagementViewPivotList {
	var returns ResourceGroupCostManagementViewPivotList
	_jsii_.Get(
		j,
		"pivot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) PivotInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pivotInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ReportType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ReportTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ResourceGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) ResourceGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Timeframe() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) TimeframeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeframeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) Timeouts() ResourceGroupCostManagementViewTimeoutsOutputReference {
	var returns ResourceGroupCostManagementViewTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementView) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/resource_group_cost_management_view azurerm_resource_group_cost_management_view} Resource.
func NewResourceGroupCostManagementView(scope constructs.Construct, id *string, config *ResourceGroupCostManagementViewConfig) ResourceGroupCostManagementView {
	_init_.Initialize()

	if err := validateNewResourceGroupCostManagementViewParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceGroupCostManagementView{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/resource_group_cost_management_view azurerm_resource_group_cost_management_view} Resource.
func NewResourceGroupCostManagementView_Override(r ResourceGroupCostManagementView, scope constructs.Construct, id *string, config *ResourceGroupCostManagementViewConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetAccumulated(val interface{}) {
	if err := j.validateSetAccumulatedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accumulated",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetChartType(val *string) {
	if err := j.validateSetChartTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"chartType",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetDisplayName(val *string) {
	if err := j.validateSetDisplayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"displayName",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetReportType(val *string) {
	if err := j.validateSetReportTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportType",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetResourceGroupId(val *string) {
	if err := j.validateSetResourceGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupId",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementView)SetTimeframe(val *string) {
	if err := j.validateSetTimeframeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeframe",
		val,
	)
}

// Generates CDKTF code for importing a ResourceGroupCostManagementView resource upon running "cdktf plan <stack-name>".
func ResourceGroupCostManagementView_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateResourceGroupCostManagementView_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
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
func ResourceGroupCostManagementView_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceGroupCostManagementView_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceGroupCostManagementView_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceGroupCostManagementView_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ResourceGroupCostManagementView_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateResourceGroupCostManagementView_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ResourceGroupCostManagementView_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementView",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) AddMoveTarget(moveTarget *string) {
	if err := r.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := r.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) MoveFromId(id *string) {
	if err := r.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveFromId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) MoveTo(moveTarget *string, index interface{}) {
	if err := r.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) MoveToId(id *string) {
	if err := r.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"moveToId",
		[]interface{}{id},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) PutDataset(value *ResourceGroupCostManagementViewDataset) {
	if err := r.validatePutDatasetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putDataset",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) PutKpi(value interface{}) {
	if err := r.validatePutKpiParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putKpi",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) PutPivot(value interface{}) {
	if err := r.validatePutPivotParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPivot",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) PutTimeouts(value *ResourceGroupCostManagementViewTimeouts) {
	if err := r.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ResetKpi() {
	_jsii_.InvokeVoid(
		r,
		"resetKpi",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ResetPivot() {
	_jsii_.InvokeVoid(
		r,
		"resetPivot",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ResetTimeouts() {
	_jsii_.InvokeVoid(
		r,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementView) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementView) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

