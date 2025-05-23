// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountagreement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/logicappintegrationaccountagreement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/logic_app_integration_account_agreement azurerm_logic_app_integration_account_agreement}.
type LogicAppIntegrationAccountAgreement interface {
	cdktf.TerraformResource
	AgreementType() *string
	SetAgreementType(val *string)
	AgreementTypeInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
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
	GuestIdentity() LogicAppIntegrationAccountAgreementGuestIdentityOutputReference
	GuestIdentityInput() *LogicAppIntegrationAccountAgreementGuestIdentity
	GuestPartnerName() *string
	SetGuestPartnerName(val *string)
	GuestPartnerNameInput() *string
	HostIdentity() LogicAppIntegrationAccountAgreementHostIdentityOutputReference
	HostIdentityInput() *LogicAppIntegrationAccountAgreementHostIdentity
	HostPartnerName() *string
	SetHostPartnerName(val *string)
	HostPartnerNameInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntegrationAccountName() *string
	SetIntegrationAccountName(val *string)
	IntegrationAccountNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Metadata() *map[string]*string
	SetMetadata(val *map[string]*string)
	MetadataInput() *map[string]*string
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
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LogicAppIntegrationAccountAgreementTimeoutsOutputReference
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
	PutGuestIdentity(value *LogicAppIntegrationAccountAgreementGuestIdentity)
	PutHostIdentity(value *LogicAppIntegrationAccountAgreementHostIdentity)
	PutTimeouts(value *LogicAppIntegrationAccountAgreementTimeouts)
	ResetId()
	ResetMetadata()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for LogicAppIntegrationAccountAgreement
type jsiiProxy_LogicAppIntegrationAccountAgreement struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) AgreementType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agreementType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) AgreementTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"agreementTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) GuestIdentity() LogicAppIntegrationAccountAgreementGuestIdentityOutputReference {
	var returns LogicAppIntegrationAccountAgreementGuestIdentityOutputReference
	_jsii_.Get(
		j,
		"guestIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) GuestIdentityInput() *LogicAppIntegrationAccountAgreementGuestIdentity {
	var returns *LogicAppIntegrationAccountAgreementGuestIdentity
	_jsii_.Get(
		j,
		"guestIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) GuestPartnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestPartnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) GuestPartnerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"guestPartnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) HostIdentity() LogicAppIntegrationAccountAgreementHostIdentityOutputReference {
	var returns LogicAppIntegrationAccountAgreementHostIdentityOutputReference
	_jsii_.Get(
		j,
		"hostIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) HostIdentityInput() *LogicAppIntegrationAccountAgreementHostIdentity {
	var returns *LogicAppIntegrationAccountAgreementHostIdentity
	_jsii_.Get(
		j,
		"hostIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) HostPartnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPartnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) HostPartnerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostPartnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) IntegrationAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) IntegrationAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Metadata() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) MetadataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) Timeouts() LogicAppIntegrationAccountAgreementTimeoutsOutputReference {
	var returns LogicAppIntegrationAccountAgreementTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/logic_app_integration_account_agreement azurerm_logic_app_integration_account_agreement} Resource.
func NewLogicAppIntegrationAccountAgreement(scope constructs.Construct, id *string, config *LogicAppIntegrationAccountAgreementConfig) LogicAppIntegrationAccountAgreement {
	_init_.Initialize()

	if err := validateNewLogicAppIntegrationAccountAgreementParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppIntegrationAccountAgreement{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/logic_app_integration_account_agreement azurerm_logic_app_integration_account_agreement} Resource.
func NewLogicAppIntegrationAccountAgreement_Override(l LogicAppIntegrationAccountAgreement, scope constructs.Construct, id *string, config *LogicAppIntegrationAccountAgreementConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetAgreementType(val *string) {
	if err := j.validateSetAgreementTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"agreementType",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetGuestPartnerName(val *string) {
	if err := j.validateSetGuestPartnerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"guestPartnerName",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetHostPartnerName(val *string) {
	if err := j.validateSetHostPartnerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostPartnerName",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetIntegrationAccountName(val *string) {
	if err := j.validateSetIntegrationAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationAccountName",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetMetadata(val *map[string]*string) {
	if err := j.validateSetMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LogicAppIntegrationAccountAgreement)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

// Generates CDKTF code for importing a LogicAppIntegrationAccountAgreement resource upon running "cdktf plan <stack-name>".
func LogicAppIntegrationAccountAgreement_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLogicAppIntegrationAccountAgreement_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
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
func LogicAppIntegrationAccountAgreement_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogicAppIntegrationAccountAgreement_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogicAppIntegrationAccountAgreement_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogicAppIntegrationAccountAgreement_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LogicAppIntegrationAccountAgreement_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLogicAppIntegrationAccountAgreement_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LogicAppIntegrationAccountAgreement_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.logicAppIntegrationAccountAgreement.LogicAppIntegrationAccountAgreement",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) PutGuestIdentity(value *LogicAppIntegrationAccountAgreementGuestIdentity) {
	if err := l.validatePutGuestIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGuestIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) PutHostIdentity(value *LogicAppIntegrationAccountAgreementHostIdentity) {
	if err := l.validatePutHostIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putHostIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) PutTimeouts(value *LogicAppIntegrationAccountAgreementTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ResetMetadata() {
	_jsii_.InvokeVoid(
		l,
		"resetMetadata",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppIntegrationAccountAgreement) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

