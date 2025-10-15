// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mssqljobstep

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mssqljobstep/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/mssql_job_step azurerm_mssql_job_step}.
type MssqlJobStep interface {
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
	InitialRetryIntervalSeconds() *float64
	SetInitialRetryIntervalSeconds(val *float64)
	InitialRetryIntervalSecondsInput() *float64
	JobCredentialId() *string
	SetJobCredentialId(val *string)
	JobCredentialIdInput() *string
	JobId() *string
	SetJobId(val *string)
	JobIdInput() *string
	JobStepIndex() *float64
	SetJobStepIndex(val *float64)
	JobStepIndexInput() *float64
	JobTargetGroupId() *string
	SetJobTargetGroupId(val *string)
	JobTargetGroupIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaximumRetryIntervalSeconds() *float64
	SetMaximumRetryIntervalSeconds(val *float64)
	MaximumRetryIntervalSecondsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OutputTarget() MssqlJobStepOutputTargetOutputReference
	OutputTargetInput() *MssqlJobStepOutputTarget
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
	RetryAttempts() *float64
	SetRetryAttempts(val *float64)
	RetryAttemptsInput() *float64
	RetryIntervalBackoffMultiplier() *float64
	SetRetryIntervalBackoffMultiplier(val *float64)
	RetryIntervalBackoffMultiplierInput() *float64
	SqlScript() *string
	SetSqlScript(val *string)
	SqlScriptInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MssqlJobStepTimeoutsOutputReference
	TimeoutSeconds() *float64
	SetTimeoutSeconds(val *float64)
	TimeoutSecondsInput() *float64
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
	PutOutputTarget(value *MssqlJobStepOutputTarget)
	PutTimeouts(value *MssqlJobStepTimeouts)
	ResetId()
	ResetInitialRetryIntervalSeconds()
	ResetJobCredentialId()
	ResetMaximumRetryIntervalSeconds()
	ResetOutputTarget()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRetryAttempts()
	ResetRetryIntervalBackoffMultiplier()
	ResetTimeouts()
	ResetTimeoutSeconds()
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

// The jsii proxy struct for MssqlJobStep
type jsiiProxy_MssqlJobStep struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MssqlJobStep) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) InitialRetryIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialRetryIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) InitialRetryIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"initialRetryIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobCredentialId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobCredentialId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobCredentialIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobCredentialIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobStepIndex() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobStepIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobStepIndexInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jobStepIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobTargetGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTargetGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) JobTargetGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jobTargetGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) MaximumRetryIntervalSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryIntervalSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) MaximumRetryIntervalSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumRetryIntervalSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) OutputTarget() MssqlJobStepOutputTargetOutputReference {
	var returns MssqlJobStepOutputTargetOutputReference
	_jsii_.Get(
		j,
		"outputTarget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) OutputTargetInput() *MssqlJobStepOutputTarget {
	var returns *MssqlJobStepOutputTarget
	_jsii_.Get(
		j,
		"outputTargetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) RetryAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) RetryAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) RetryIntervalBackoffMultiplier() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalBackoffMultiplier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) RetryIntervalBackoffMultiplierInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalBackoffMultiplierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) SqlScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) SqlScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sqlScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) Timeouts() MssqlJobStepTimeoutsOutputReference {
	var returns MssqlJobStepTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) TimeoutSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) TimeoutSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlJobStep) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/mssql_job_step azurerm_mssql_job_step} Resource.
func NewMssqlJobStep(scope constructs.Construct, id *string, config *MssqlJobStepConfig) MssqlJobStep {
	_init_.Initialize()

	if err := validateNewMssqlJobStepParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlJobStep{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/mssql_job_step azurerm_mssql_job_step} Resource.
func NewMssqlJobStep_Override(m MssqlJobStep, scope constructs.Construct, id *string, config *MssqlJobStepConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetInitialRetryIntervalSeconds(val *float64) {
	if err := j.validateSetInitialRetryIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"initialRetryIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetJobCredentialId(val *string) {
	if err := j.validateSetJobCredentialIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobCredentialId",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetJobId(val *string) {
	if err := j.validateSetJobIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobId",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetJobStepIndex(val *float64) {
	if err := j.validateSetJobStepIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobStepIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetJobTargetGroupId(val *string) {
	if err := j.validateSetJobTargetGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jobTargetGroupId",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetMaximumRetryIntervalSeconds(val *float64) {
	if err := j.validateSetMaximumRetryIntervalSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumRetryIntervalSeconds",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetRetryAttempts(val *float64) {
	if err := j.validateSetRetryAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryAttempts",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetRetryIntervalBackoffMultiplier(val *float64) {
	if err := j.validateSetRetryIntervalBackoffMultiplierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryIntervalBackoffMultiplier",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetSqlScript(val *string) {
	if err := j.validateSetSqlScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sqlScript",
		val,
	)
}

func (j *jsiiProxy_MssqlJobStep)SetTimeoutSeconds(val *float64) {
	if err := j.validateSetTimeoutSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutSeconds",
		val,
	)
}

// Generates CDKTF code for importing a MssqlJobStep resource upon running "cdktf plan <stack-name>".
func MssqlJobStep_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMssqlJobStep_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
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
func MssqlJobStep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlJobStep_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlJobStep_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlJobStep_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MssqlJobStep_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMssqlJobStep_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MssqlJobStep_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mssqlJobStep.MssqlJobStep",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MssqlJobStep) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MssqlJobStep) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MssqlJobStep) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MssqlJobStep) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlJobStep) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MssqlJobStep) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MssqlJobStep) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MssqlJobStep) PutOutputTarget(value *MssqlJobStepOutputTarget) {
	if err := m.validatePutOutputTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putOutputTarget",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlJobStep) PutTimeouts(value *MssqlJobStepTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetInitialRetryIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetInitialRetryIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetJobCredentialId() {
	_jsii_.InvokeVoid(
		m,
		"resetJobCredentialId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetMaximumRetryIntervalSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetMaximumRetryIntervalSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetOutputTarget() {
	_jsii_.InvokeVoid(
		m,
		"resetOutputTarget",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetRetryAttempts() {
	_jsii_.InvokeVoid(
		m,
		"resetRetryAttempts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetRetryIntervalBackoffMultiplier() {
	_jsii_.InvokeVoid(
		m,
		"resetRetryIntervalBackoffMultiplier",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) ResetTimeoutSeconds() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeoutSeconds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlJobStep) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlJobStep) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

