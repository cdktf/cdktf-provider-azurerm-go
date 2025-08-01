// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storagecontainerimmutabilitypolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/storagecontainerimmutabilitypolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/storage_container_immutability_policy azurerm_storage_container_immutability_policy}.
type StorageContainerImmutabilityPolicy interface {
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
	ImmutabilityPeriodInDays() *float64
	SetImmutabilityPeriodInDays(val *float64)
	ImmutabilityPeriodInDaysInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locked() interface{}
	SetLocked(val interface{})
	LockedInput() interface{}
	// The tree node.
	Node() constructs.Node
	ProtectedAppendWritesAllEnabled() interface{}
	SetProtectedAppendWritesAllEnabled(val interface{})
	ProtectedAppendWritesAllEnabledInput() interface{}
	ProtectedAppendWritesEnabled() interface{}
	SetProtectedAppendWritesEnabled(val interface{})
	ProtectedAppendWritesEnabledInput() interface{}
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
	StorageContainerResourceManagerId() *string
	SetStorageContainerResourceManagerId(val *string)
	StorageContainerResourceManagerIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() StorageContainerImmutabilityPolicyTimeoutsOutputReference
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
	PutTimeouts(value *StorageContainerImmutabilityPolicyTimeouts)
	ResetId()
	ResetLocked()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetProtectedAppendWritesAllEnabled()
	ResetProtectedAppendWritesEnabled()
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

// The jsii proxy struct for StorageContainerImmutabilityPolicy
type jsiiProxy_StorageContainerImmutabilityPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ImmutabilityPeriodInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"immutabilityPeriodInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ImmutabilityPeriodInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"immutabilityPeriodInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Locked() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"locked",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) LockedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ProtectedAppendWritesAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedAppendWritesAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ProtectedAppendWritesAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedAppendWritesAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ProtectedAppendWritesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedAppendWritesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) ProtectedAppendWritesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedAppendWritesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) StorageContainerResourceManagerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerResourceManagerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) StorageContainerResourceManagerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageContainerResourceManagerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) Timeouts() StorageContainerImmutabilityPolicyTimeoutsOutputReference {
	var returns StorageContainerImmutabilityPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/storage_container_immutability_policy azurerm_storage_container_immutability_policy} Resource.
func NewStorageContainerImmutabilityPolicy(scope constructs.Construct, id *string, config *StorageContainerImmutabilityPolicyConfig) StorageContainerImmutabilityPolicy {
	_init_.Initialize()

	if err := validateNewStorageContainerImmutabilityPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_StorageContainerImmutabilityPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/storage_container_immutability_policy azurerm_storage_container_immutability_policy} Resource.
func NewStorageContainerImmutabilityPolicy_Override(s StorageContainerImmutabilityPolicy, scope constructs.Construct, id *string, config *StorageContainerImmutabilityPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetImmutabilityPeriodInDays(val *float64) {
	if err := j.validateSetImmutabilityPeriodInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"immutabilityPeriodInDays",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetLocked(val interface{}) {
	if err := j.validateSetLockedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locked",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetProtectedAppendWritesAllEnabled(val interface{}) {
	if err := j.validateSetProtectedAppendWritesAllEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedAppendWritesAllEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetProtectedAppendWritesEnabled(val interface{}) {
	if err := j.validateSetProtectedAppendWritesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protectedAppendWritesEnabled",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_StorageContainerImmutabilityPolicy)SetStorageContainerResourceManagerId(val *string) {
	if err := j.validateSetStorageContainerResourceManagerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageContainerResourceManagerId",
		val,
	)
}

// Generates CDKTF code for importing a StorageContainerImmutabilityPolicy resource upon running "cdktf plan <stack-name>".
func StorageContainerImmutabilityPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateStorageContainerImmutabilityPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
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
func StorageContainerImmutabilityPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageContainerImmutabilityPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageContainerImmutabilityPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageContainerImmutabilityPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func StorageContainerImmutabilityPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateStorageContainerImmutabilityPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func StorageContainerImmutabilityPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.storageContainerImmutabilityPolicy.StorageContainerImmutabilityPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) PutTimeouts(value *StorageContainerImmutabilityPolicyTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ResetLocked() {
	_jsii_.InvokeVoid(
		s,
		"resetLocked",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ResetProtectedAppendWritesAllEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetProtectedAppendWritesAllEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ResetProtectedAppendWritesEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetProtectedAppendWritesEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_StorageContainerImmutabilityPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

