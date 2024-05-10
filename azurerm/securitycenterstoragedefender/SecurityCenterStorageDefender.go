// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycenterstoragedefender

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/securitycenterstoragedefender/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/security_center_storage_defender azurerm_security_center_storage_defender}.
type SecurityCenterStorageDefender interface {
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MalwareScanningOnUploadCapGbPerMonth() *float64
	SetMalwareScanningOnUploadCapGbPerMonth(val *float64)
	MalwareScanningOnUploadCapGbPerMonthInput() *float64
	MalwareScanningOnUploadEnabled() interface{}
	SetMalwareScanningOnUploadEnabled(val interface{})
	MalwareScanningOnUploadEnabledInput() interface{}
	// The tree node.
	Node() constructs.Node
	OverrideSubscriptionSettingsEnabled() interface{}
	SetOverrideSubscriptionSettingsEnabled(val interface{})
	OverrideSubscriptionSettingsEnabledInput() interface{}
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
	SensitiveDataDiscoveryEnabled() interface{}
	SetSensitiveDataDiscoveryEnabled(val interface{})
	SensitiveDataDiscoveryEnabledInput() interface{}
	StorageAccountId() *string
	SetStorageAccountId(val *string)
	StorageAccountIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() SecurityCenterStorageDefenderTimeoutsOutputReference
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
	PutTimeouts(value *SecurityCenterStorageDefenderTimeouts)
	ResetId()
	ResetMalwareScanningOnUploadCapGbPerMonth()
	ResetMalwareScanningOnUploadEnabled()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverrideSubscriptionSettingsEnabled()
	ResetSensitiveDataDiscoveryEnabled()
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

// The jsii proxy struct for SecurityCenterStorageDefender
type jsiiProxy_SecurityCenterStorageDefender struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_SecurityCenterStorageDefender) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) MalwareScanningOnUploadCapGbPerMonth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"malwareScanningOnUploadCapGbPerMonth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) MalwareScanningOnUploadCapGbPerMonthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"malwareScanningOnUploadCapGbPerMonthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) MalwareScanningOnUploadEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareScanningOnUploadEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) MalwareScanningOnUploadEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"malwareScanningOnUploadEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) OverrideSubscriptionSettingsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideSubscriptionSettingsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) OverrideSubscriptionSettingsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overrideSubscriptionSettingsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) SensitiveDataDiscoveryEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveDataDiscoveryEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) SensitiveDataDiscoveryEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sensitiveDataDiscoveryEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) StorageAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) StorageAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) Timeouts() SecurityCenterStorageDefenderTimeoutsOutputReference {
	var returns SecurityCenterStorageDefenderTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SecurityCenterStorageDefender) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/security_center_storage_defender azurerm_security_center_storage_defender} Resource.
func NewSecurityCenterStorageDefender(scope constructs.Construct, id *string, config *SecurityCenterStorageDefenderConfig) SecurityCenterStorageDefender {
	_init_.Initialize()

	if err := validateNewSecurityCenterStorageDefenderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_SecurityCenterStorageDefender{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/security_center_storage_defender azurerm_security_center_storage_defender} Resource.
func NewSecurityCenterStorageDefender_Override(s SecurityCenterStorageDefender, scope constructs.Construct, id *string, config *SecurityCenterStorageDefenderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
		[]interface{}{scope, id, config},
		s,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetMalwareScanningOnUploadCapGbPerMonth(val *float64) {
	if err := j.validateSetMalwareScanningOnUploadCapGbPerMonthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"malwareScanningOnUploadCapGbPerMonth",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetMalwareScanningOnUploadEnabled(val interface{}) {
	if err := j.validateSetMalwareScanningOnUploadEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"malwareScanningOnUploadEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetOverrideSubscriptionSettingsEnabled(val interface{}) {
	if err := j.validateSetOverrideSubscriptionSettingsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overrideSubscriptionSettingsEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetSensitiveDataDiscoveryEnabled(val interface{}) {
	if err := j.validateSetSensitiveDataDiscoveryEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sensitiveDataDiscoveryEnabled",
		val,
	)
}

func (j *jsiiProxy_SecurityCenterStorageDefender)SetStorageAccountId(val *string) {
	if err := j.validateSetStorageAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountId",
		val,
	)
}

// Generates CDKTF code for importing a SecurityCenterStorageDefender resource upon running "cdktf plan <stack-name>".
func SecurityCenterStorageDefender_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateSecurityCenterStorageDefender_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
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
func SecurityCenterStorageDefender_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityCenterStorageDefender_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecurityCenterStorageDefender_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityCenterStorageDefender_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func SecurityCenterStorageDefender_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateSecurityCenterStorageDefender_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func SecurityCenterStorageDefender_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.securityCenterStorageDefender.SecurityCenterStorageDefender",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) AddMoveTarget(moveTarget *string) {
	if err := s.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) AddOverride(path *string, value interface{}) {
	if err := s.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := s.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SecurityCenterStorageDefender) MoveFromId(id *string) {
	if err := s.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveFromId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) MoveTo(moveTarget *string, index interface{}) {
	if err := s.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) MoveToId(id *string) {
	if err := s.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"moveToId",
		[]interface{}{id},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) OverrideLogicalId(newLogicalId *string) {
	if err := s.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) PutTimeouts(value *SecurityCenterStorageDefenderTimeouts) {
	if err := s.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetId() {
	_jsii_.InvokeVoid(
		s,
		"resetId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetMalwareScanningOnUploadCapGbPerMonth() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareScanningOnUploadCapGbPerMonth",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetMalwareScanningOnUploadEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetMalwareScanningOnUploadEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetOverrideSubscriptionSettingsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetOverrideSubscriptionSettingsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetSensitiveDataDiscoveryEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetSensitiveDataDiscoveryEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ResetTimeouts() {
	_jsii_.InvokeVoid(
		s,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SecurityCenterStorageDefender) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SecurityCenterStorageDefender) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		s,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

