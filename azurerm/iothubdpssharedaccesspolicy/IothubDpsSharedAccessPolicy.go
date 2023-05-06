package iothubdpssharedaccesspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/iothubdpssharedaccesspolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps_shared_access_policy azurerm_iothub_dps_shared_access_policy}.
type IothubDpsSharedAccessPolicy interface {
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
	EnrollmentRead() interface{}
	SetEnrollmentRead(val interface{})
	EnrollmentReadInput() interface{}
	EnrollmentWrite() interface{}
	SetEnrollmentWrite(val interface{})
	EnrollmentWriteInput() interface{}
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
	IothubDpsName() *string
	SetIothubDpsName(val *string)
	IothubDpsNameInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	PrimaryConnectionString() *string
	PrimaryKey() *string
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
	RegistrationRead() interface{}
	SetRegistrationRead(val interface{})
	RegistrationReadInput() interface{}
	RegistrationWrite() interface{}
	SetRegistrationWrite(val interface{})
	RegistrationWriteInput() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	SecondaryConnectionString() *string
	SecondaryKey() *string
	ServiceConfig() interface{}
	SetServiceConfig(val interface{})
	ServiceConfigInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() IothubDpsSharedAccessPolicyTimeoutsOutputReference
	TimeoutsInput() interface{}
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
	PutTimeouts(value *IothubDpsSharedAccessPolicyTimeouts)
	ResetEnrollmentRead()
	ResetEnrollmentWrite()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRegistrationRead()
	ResetRegistrationWrite()
	ResetServiceConfig()
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

// The jsii proxy struct for IothubDpsSharedAccessPolicy
type jsiiProxy_IothubDpsSharedAccessPolicy struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) EnrollmentRead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrollmentRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) EnrollmentReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrollmentReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) EnrollmentWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrollmentWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) EnrollmentWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enrollmentWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) IothubDpsName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubDpsName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) IothubDpsNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubDpsNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) PrimaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) PrimaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) RegistrationRead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrationRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) RegistrationReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrationReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) RegistrationWrite() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrationWrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) RegistrationWriteInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"registrationWriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) SecondaryConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) SecondaryKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) ServiceConfig() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) ServiceConfigInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serviceConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) Timeouts() IothubDpsSharedAccessPolicyTimeoutsOutputReference {
	var returns IothubDpsSharedAccessPolicyTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps_shared_access_policy azurerm_iothub_dps_shared_access_policy} Resource.
func NewIothubDpsSharedAccessPolicy(scope constructs.Construct, id *string, config *IothubDpsSharedAccessPolicyConfig) IothubDpsSharedAccessPolicy {
	_init_.Initialize()

	if err := validateNewIothubDpsSharedAccessPolicyParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IothubDpsSharedAccessPolicy{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothubDpsSharedAccessPolicy.IothubDpsSharedAccessPolicy",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/iothub_dps_shared_access_policy azurerm_iothub_dps_shared_access_policy} Resource.
func NewIothubDpsSharedAccessPolicy_Override(i IothubDpsSharedAccessPolicy, scope constructs.Construct, id *string, config *IothubDpsSharedAccessPolicyConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothubDpsSharedAccessPolicy.IothubDpsSharedAccessPolicy",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetEnrollmentRead(val interface{}) {
	if err := j.validateSetEnrollmentReadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enrollmentRead",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetEnrollmentWrite(val interface{}) {
	if err := j.validateSetEnrollmentWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enrollmentWrite",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetIothubDpsName(val *string) {
	if err := j.validateSetIothubDpsNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iothubDpsName",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetRegistrationRead(val interface{}) {
	if err := j.validateSetRegistrationReadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrationRead",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetRegistrationWrite(val interface{}) {
	if err := j.validateSetRegistrationWriteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"registrationWrite",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_IothubDpsSharedAccessPolicy)SetServiceConfig(val interface{}) {
	if err := j.validateSetServiceConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceConfig",
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
func IothubDpsSharedAccessPolicy_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubDpsSharedAccessPolicy_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubDpsSharedAccessPolicy.IothubDpsSharedAccessPolicy",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IothubDpsSharedAccessPolicy_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubDpsSharedAccessPolicy_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubDpsSharedAccessPolicy.IothubDpsSharedAccessPolicy",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IothubDpsSharedAccessPolicy_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubDpsSharedAccessPolicy_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubDpsSharedAccessPolicy.IothubDpsSharedAccessPolicy",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IothubDpsSharedAccessPolicy_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.iothubDpsSharedAccessPolicy.IothubDpsSharedAccessPolicy",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) PutTimeouts(value *IothubDpsSharedAccessPolicyTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetEnrollmentRead() {
	_jsii_.InvokeVoid(
		i,
		"resetEnrollmentRead",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetEnrollmentWrite() {
	_jsii_.InvokeVoid(
		i,
		"resetEnrollmentWrite",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetRegistrationRead() {
	_jsii_.InvokeVoid(
		i,
		"resetRegistrationRead",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetRegistrationWrite() {
	_jsii_.InvokeVoid(
		i,
		"resetRegistrationWrite",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetServiceConfig() {
	_jsii_.InvokeVoid(
		i,
		"resetServiceConfig",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubDpsSharedAccessPolicy) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

