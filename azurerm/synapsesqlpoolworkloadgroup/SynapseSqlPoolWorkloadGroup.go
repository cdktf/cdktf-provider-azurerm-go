// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesqlpoolworkloadgroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/synapsesqlpoolworkloadgroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/synapse_sql_pool_workload_group azurerm_synapse_sql_pool_workload_group}.
type SynapseSqlPoolWorkloadGroup interface {
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
	Importance() *string
	SetImportance(val *string)
	ImportanceInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxResourcePercent() *float64
	SetMaxResourcePercent(val *float64)
	MaxResourcePercentInput() *float64
	MaxResourcePercentPerRequest() *float64
	SetMaxResourcePercentPerRequest(val *float64)
	MaxResourcePercentPerRequestInput() *float64
	MinResourcePercent() *float64
	SetMinResourcePercent(val *float64)
	MinResourcePercentInput() *float64
	MinResourcePercentPerRequest() *float64
	SetMinResourcePercentPerRequest(val *float64)
	MinResourcePercentPerRequestInput() *float64
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
	QueryExecutionTimeoutInSeconds() *float64
	SetQueryExecutionTimeoutInSeconds(val *float64)
	QueryExecutionTimeoutInSecondsInput() *float64
	// Experimental.
	RawOverrides() interface{}
	SqlPoolId() *string
	SetSqlPoolId(val *string)
	SqlPoolIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SynapseSqlPoolWorkloadGroupTimeoutsOutputReference
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
	PutTimeouts(value *SynapseSqlPoolWorkloadGroupTimeouts)
	ResetId()
	ResetImportance()
	ResetMaxResourcePercentPerRequest()
	ResetMinResourcePercentPerRequest()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetQueryExecutionTimeoutInSeconds()
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

// The jsii proxy struct for SynapseSqlPoolWorkloadGroup
type jsiiProxy_SynapseSqlPoolWorkloadGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Importance() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) ImportanceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MaxResourcePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResourcePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MaxResourcePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResourcePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MaxResourcePercentPerRequest() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResourcePercentPerRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MaxResourcePercentPerRequestInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResourcePercentPerRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MinResourcePercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minResourcePercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MinResourcePercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minResourcePercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MinResourcePercentPerRequest() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minResourcePercentPerRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) MinResourcePercentPerRequestInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minResourcePercentPerRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) QueryExecutionTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryExecutionTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) QueryExecutionTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"queryExecutionTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) SqlPoolId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlPoolId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) SqlPoolIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlPoolIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) Timeouts() SynapseSqlPoolWorkloadGroupTimeoutsOutputReference {
	var returns SynapseSqlPoolWorkloadGroupTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/synapse_sql_pool_workload_group azurerm_synapse_sql_pool_workload_group} Resource.
func NewSynapseSqlPoolWorkloadGroup(scope constructs.Construct, id *string, config *SynapseSqlPoolWorkloadGroupConfig) SynapseSqlPoolWorkloadGroup {
	_init_.Initialize()

	if err := validateNewSynapseSqlPoolWorkloadGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SynapseSqlPoolWorkloadGroup{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/synapse_sql_pool_workload_group azurerm_synapse_sql_pool_workload_group} Resource.
func NewSynapseSqlPoolWorkloadGroup_Override(s SynapseSqlPoolWorkloadGroup, scope constructs.Construct, id *string, config *SynapseSqlPoolWorkloadGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetImportance(val *string) {
	if err := j.validateSetImportanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importance",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetMaxResourcePercent(val *float64) {
	if err := j.validateSetMaxResourcePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxResourcePercent",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetMaxResourcePercentPerRequest(val *float64) {
	if err := j.validateSetMaxResourcePercentPerRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxResourcePercentPerRequest",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetMinResourcePercent(val *float64) {
	if err := j.validateSetMinResourcePercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minResourcePercent",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetMinResourcePercentPerRequest(val *float64) {
	if err := j.validateSetMinResourcePercentPerRequestParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minResourcePercentPerRequest",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetQueryExecutionTimeoutInSeconds(val *float64) {
	if err := j.validateSetQueryExecutionTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryExecutionTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SynapseSqlPoolWorkloadGroup)SetSqlPoolId(val *string) {
	if err := j.validateSetSqlPoolIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlPoolId",
		val,
	)
}

// Generates CDKTF code for importing a SynapseSqlPoolWorkloadGroup resource upon running "cdktf plan <stack-name>".
func SynapseSqlPoolWorkloadGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSynapseSqlPoolWorkloadGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
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
func SynapseSqlPoolWorkloadGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseSqlPoolWorkloadGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SynapseSqlPoolWorkloadGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseSqlPoolWorkloadGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SynapseSqlPoolWorkloadGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSynapseSqlPoolWorkloadGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SynapseSqlPoolWorkloadGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.synapseSqlPoolWorkloadGroup.SynapseSqlPoolWorkloadGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) PutTimeouts(value *SynapseSqlPoolWorkloadGroupTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetImportance() {
	_jsii_.InvokeVoid(
		s,
		"resetImportance",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetMaxResourcePercentPerRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxResourcePercentPerRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetMinResourcePercentPerRequest() {
	_jsii_.InvokeVoid(
		s,
		"resetMinResourcePercentPerRequest",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetQueryExecutionTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetQueryExecutionTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SynapseSqlPoolWorkloadGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

