// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mobilenetworksim

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/mobilenetworksim/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mobile_network_sim azurerm_mobile_network_sim}.
type MobileNetworkSim interface {
	cdktf.TerraformResource
	AuthenticationKey() *string
	SetAuthenticationKey(val *string)
	AuthenticationKeyInput() *string
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
	DeviceType() *string
	SetDeviceType(val *string)
	DeviceTypeInput() *string
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
	IntegratedCircuitCardIdentifier() *string
	SetIntegratedCircuitCardIdentifier(val *string)
	IntegratedCircuitCardIdentifierInput() *string
	InternationalMobileSubscriberIdentity() *string
	SetInternationalMobileSubscriberIdentity(val *string)
	InternationalMobileSubscriberIdentityInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MobileNetworkSimGroupId() *string
	SetMobileNetworkSimGroupId(val *string)
	MobileNetworkSimGroupIdInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	OperatorKeyCode() *string
	SetOperatorKeyCode(val *string)
	OperatorKeyCodeInput() *string
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
	SimPolicyId() *string
	SetSimPolicyId(val *string)
	SimPolicyIdInput() *string
	SimState() *string
	StaticIpConfiguration() MobileNetworkSimStaticIpConfigurationList
	StaticIpConfigurationInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() MobileNetworkSimTimeoutsOutputReference
	TimeoutsInput() interface{}
	VendorKeyFingerprint() *string
	VendorName() *string
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
	PutStaticIpConfiguration(value interface{})
	PutTimeouts(value *MobileNetworkSimTimeouts)
	ResetDeviceType()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSimPolicyId()
	ResetStaticIpConfiguration()
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

// The jsii proxy struct for MobileNetworkSim
type jsiiProxy_MobileNetworkSim struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MobileNetworkSim) AuthenticationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) AuthenticationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) DeviceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) DeviceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) IntegratedCircuitCardIdentifier() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integratedCircuitCardIdentifier",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) IntegratedCircuitCardIdentifierInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integratedCircuitCardIdentifierInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) InternationalMobileSubscriberIdentity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internationalMobileSubscriberIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) InternationalMobileSubscriberIdentityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"internationalMobileSubscriberIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) MobileNetworkSimGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkSimGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) MobileNetworkSimGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mobileNetworkSimGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) OperatorKeyCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorKeyCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) OperatorKeyCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operatorKeyCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) SimPolicyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simPolicyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) SimPolicyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simPolicyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) SimState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"simState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) StaticIpConfiguration() MobileNetworkSimStaticIpConfigurationList {
	var returns MobileNetworkSimStaticIpConfigurationList
	_jsii_.Get(
		j,
		"staticIpConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) StaticIpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"staticIpConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) Timeouts() MobileNetworkSimTimeoutsOutputReference {
	var returns MobileNetworkSimTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) VendorKeyFingerprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorKeyFingerprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkSim) VendorName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vendorName",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mobile_network_sim azurerm_mobile_network_sim} Resource.
func NewMobileNetworkSim(scope constructs.Construct, id *string, config *MobileNetworkSimConfig) MobileNetworkSim {
	_init_.Initialize()

	if err := validateNewMobileNetworkSimParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkSim{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/mobile_network_sim azurerm_mobile_network_sim} Resource.
func NewMobileNetworkSim_Override(m MobileNetworkSim, scope constructs.Construct, id *string, config *MobileNetworkSimConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetAuthenticationKey(val *string) {
	if err := j.validateSetAuthenticationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationKey",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetDeviceType(val *string) {
	if err := j.validateSetDeviceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceType",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetIntegratedCircuitCardIdentifier(val *string) {
	if err := j.validateSetIntegratedCircuitCardIdentifierParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integratedCircuitCardIdentifier",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetInternationalMobileSubscriberIdentity(val *string) {
	if err := j.validateSetInternationalMobileSubscriberIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internationalMobileSubscriberIdentity",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetMobileNetworkSimGroupId(val *string) {
	if err := j.validateSetMobileNetworkSimGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mobileNetworkSimGroupId",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetOperatorKeyCode(val *string) {
	if err := j.validateSetOperatorKeyCodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatorKeyCode",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkSim)SetSimPolicyId(val *string) {
	if err := j.validateSetSimPolicyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"simPolicyId",
		val,
	)
}

// Generates CDKTF code for importing a MobileNetworkSim resource upon running "cdktf plan <stack-name>".
func MobileNetworkSim_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateMobileNetworkSim_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
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
func MobileNetworkSim_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkSim_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkSim_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkSim_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func MobileNetworkSim_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateMobileNetworkSim_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MobileNetworkSim_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.mobileNetworkSim.MobileNetworkSim",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MobileNetworkSim) AddMoveTarget(moveTarget *string) {
	if err := m.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (m *jsiiProxy_MobileNetworkSim) AddOverride(path *string, value interface{}) {
	if err := m.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MobileNetworkSim) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkSim) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkSim) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkSim) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkSim) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkSim) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkSim) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkSim) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkSim) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkSim) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSim) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := m.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (m *jsiiProxy_MobileNetworkSim) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkSim) MoveFromId(id *string) {
	if err := m.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveFromId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkSim) MoveTo(moveTarget *string, index interface{}) {
	if err := m.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (m *jsiiProxy_MobileNetworkSim) MoveToId(id *string) {
	if err := m.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"moveToId",
		[]interface{}{id},
	)
}

func (m *jsiiProxy_MobileNetworkSim) OverrideLogicalId(newLogicalId *string) {
	if err := m.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MobileNetworkSim) PutStaticIpConfiguration(value interface{}) {
	if err := m.validatePutStaticIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStaticIpConfiguration",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkSim) PutTimeouts(value *MobileNetworkSimTimeouts) {
	if err := m.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkSim) ResetDeviceType() {
	_jsii_.InvokeVoid(
		m,
		"resetDeviceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSim) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSim) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSim) ResetSimPolicyId() {
	_jsii_.InvokeVoid(
		m,
		"resetSimPolicyId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSim) ResetStaticIpConfiguration() {
	_jsii_.InvokeVoid(
		m,
		"resetStaticIpConfiguration",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSim) ResetTimeouts() {
	_jsii_.InvokeVoid(
		m,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkSim) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSim) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSim) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSim) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSim) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkSim) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

