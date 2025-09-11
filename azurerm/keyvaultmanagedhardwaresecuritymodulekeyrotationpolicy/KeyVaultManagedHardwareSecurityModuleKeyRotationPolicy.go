// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keyvaultmanagedhardwaresecuritymodulekeyrotationpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/keyvaultmanagedhardwaresecuritymodulekeyrotationpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/key_vault_managed_hardware_security_module_key_rotation_policy azurerm_key_vault_managed_hardware_security_module_key_rotation_policy}.
type KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy interface {
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
	ExpireAfter() *string
	SetExpireAfter(val *string)
	ExpireAfterInput() *string
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
	ManagedHsmKeyId() *string
	SetManagedHsmKeyId(val *string)
	ManagedHsmKeyIdInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TimeAfterCreation() *string
	SetTimeAfterCreation(val *string)
	TimeAfterCreationInput() *string
	TimeBeforeExpiry() *string
	SetTimeBeforeExpiry(val *string)
	TimeBeforeExpiryInput() *string
	Timeouts() KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyTimeoutsOutputReference
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
	PutTimeouts(value *KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyTimeouts)
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTimeAfterCreation()
	ResetTimeBeforeExpiry()
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

// The jsii proxy struct for KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy
type jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ExpireAfter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ExpireAfterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expireAfterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ManagedHsmKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedHsmKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ManagedHsmKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedHsmKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TimeAfterCreation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAfterCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TimeAfterCreationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeAfterCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TimeBeforeExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeBeforeExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TimeBeforeExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeBeforeExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) Timeouts() KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyTimeoutsOutputReference {
	var returns KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/key_vault_managed_hardware_security_module_key_rotation_policy azurerm_key_vault_managed_hardware_security_module_key_rotation_policy} Resource.
func NewKeyVaultManagedHardwareSecurityModuleKeyRotationPolicy(scope constructs.Construct, id *string, config *KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyConfig) KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy {
	_init_.Initialize()

	if err := validateNewKeyVaultManagedHardwareSecurityModuleKeyRotationPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/key_vault_managed_hardware_security_module_key_rotation_policy azurerm_key_vault_managed_hardware_security_module_key_rotation_policy} Resource.
func NewKeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_Override(k KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy, scope constructs.Construct, id *string, config *KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetExpireAfter(val *string) {
	if err := j.validateSetExpireAfterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireAfter",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetManagedHsmKeyId(val *string) {
	if err := j.validateSetManagedHsmKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"managedHsmKeyId",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetTimeAfterCreation(val *string) {
	if err := j.validateSetTimeAfterCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeAfterCreation",
		val,
	)
}

func (j *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy)SetTimeBeforeExpiry(val *string) {
	if err := j.validateSetTimeBeforeExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeBeforeExpiry",
		val,
	)
}

// Generates CDKTF code for importing a KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy resource upon running "cdktf plan <stack-name>".
func KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
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
func KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.keyVaultManagedHardwareSecurityModuleKeyRotationPolicy.KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) PutTimeouts(value *KeyVaultManagedHardwareSecurityModuleKeyRotationPolicyTimeouts) {
	if err := k.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ResetTimeAfterCreation() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeAfterCreation",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ResetTimeBeforeExpiry() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeBeforeExpiry",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KeyVaultManagedHardwareSecurityModuleKeyRotationPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

