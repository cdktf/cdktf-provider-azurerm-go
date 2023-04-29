package iothubdeviceupdateinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/iothubdeviceupdateinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/iothub_device_update_instance azurerm_iothub_device_update_instance}.
type IothubDeviceUpdateInstance interface {
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
	DeviceUpdateAccountId() *string
	SetDeviceUpdateAccountId(val *string)
	DeviceUpdateAccountIdInput() *string
	DiagnosticEnabled() interface{}
	SetDiagnosticEnabled(val interface{})
	DiagnosticEnabledInput() interface{}
	DiagnosticStorageAccount() IothubDeviceUpdateInstanceDiagnosticStorageAccountOutputReference
	DiagnosticStorageAccountInput() *IothubDeviceUpdateInstanceDiagnosticStorageAccount
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
	IothubId() *string
	SetIothubId(val *string)
	IothubIdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	Timeouts() IothubDeviceUpdateInstanceTimeoutsOutputReference
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
	PutDiagnosticStorageAccount(value *IothubDeviceUpdateInstanceDiagnosticStorageAccount)
	PutTimeouts(value *IothubDeviceUpdateInstanceTimeouts)
	ResetDiagnosticEnabled()
	ResetDiagnosticStorageAccount()
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

// The jsii proxy struct for IothubDeviceUpdateInstance
type jsiiProxy_IothubDeviceUpdateInstance struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DeviceUpdateAccountId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceUpdateAccountId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DeviceUpdateAccountIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deviceUpdateAccountIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DiagnosticEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diagnosticEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DiagnosticEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diagnosticEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DiagnosticStorageAccount() IothubDeviceUpdateInstanceDiagnosticStorageAccountOutputReference {
	var returns IothubDeviceUpdateInstanceDiagnosticStorageAccountOutputReference
	_jsii_.Get(
		j,
		"diagnosticStorageAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) DiagnosticStorageAccountInput() *IothubDeviceUpdateInstanceDiagnosticStorageAccount {
	var returns *IothubDeviceUpdateInstanceDiagnosticStorageAccount
	_jsii_.Get(
		j,
		"diagnosticStorageAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) IothubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) IothubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"iothubIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) Timeouts() IothubDeviceUpdateInstanceTimeoutsOutputReference {
	var returns IothubDeviceUpdateInstanceTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubDeviceUpdateInstance) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/iothub_device_update_instance azurerm_iothub_device_update_instance} Resource.
func NewIothubDeviceUpdateInstance(scope constructs.Construct, id *string, config *IothubDeviceUpdateInstanceConfig) IothubDeviceUpdateInstance {
	_init_.Initialize()

	if err := validateNewIothubDeviceUpdateInstanceParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IothubDeviceUpdateInstance{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothubDeviceUpdateInstance.IothubDeviceUpdateInstance",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/iothub_device_update_instance azurerm_iothub_device_update_instance} Resource.
func NewIothubDeviceUpdateInstance_Override(i IothubDeviceUpdateInstance, scope constructs.Construct, id *string, config *IothubDeviceUpdateInstanceConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothubDeviceUpdateInstance.IothubDeviceUpdateInstance",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetDeviceUpdateAccountId(val *string) {
	if err := j.validateSetDeviceUpdateAccountIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deviceUpdateAccountId",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetDiagnosticEnabled(val interface{}) {
	if err := j.validateSetDiagnosticEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diagnosticEnabled",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetIothubId(val *string) {
	if err := j.validateSetIothubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iothubId",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IothubDeviceUpdateInstance)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
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
func IothubDeviceUpdateInstance_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubDeviceUpdateInstance_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubDeviceUpdateInstance.IothubDeviceUpdateInstance",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IothubDeviceUpdateInstance_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubDeviceUpdateInstance_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubDeviceUpdateInstance.IothubDeviceUpdateInstance",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IothubDeviceUpdateInstance_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIothubDeviceUpdateInstance_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.iothubDeviceUpdateInstance.IothubDeviceUpdateInstance",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IothubDeviceUpdateInstance_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.iothubDeviceUpdateInstance.IothubDeviceUpdateInstance",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IothubDeviceUpdateInstance) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) PutDiagnosticStorageAccount(value *IothubDeviceUpdateInstanceDiagnosticStorageAccount) {
	if err := i.validatePutDiagnosticStorageAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putDiagnosticStorageAccount",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) PutTimeouts(value *IothubDeviceUpdateInstanceTimeouts) {
	if err := i.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ResetDiagnosticEnabled() {
	_jsii_.InvokeVoid(
		i,
		"resetDiagnosticEnabled",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ResetDiagnosticStorageAccount() {
	_jsii_.InvokeVoid(
		i,
		"resetDiagnosticStorageAccount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ResetTags() {
	_jsii_.InvokeVoid(
		i,
		"resetTags",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ResetTimeouts() {
	_jsii_.InvokeVoid(
		i,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubDeviceUpdateInstance) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

