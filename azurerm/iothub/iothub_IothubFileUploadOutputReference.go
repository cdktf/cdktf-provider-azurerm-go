package iothub

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/iothub/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IothubFileUploadOutputReference interface {
	cdktf.ComplexObject
	AuthenticationType() *string
	SetAuthenticationType(val *string)
	AuthenticationTypeInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectionString() *string
	SetConnectionString(val *string)
	ConnectionStringInput() *string
	ContainerName() *string
	SetContainerName(val *string)
	ContainerNameInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultTtl() *string
	SetDefaultTtl(val *string)
	DefaultTtlInput() *string
	// Experimental.
	Fqn() *string
	IdentityId() *string
	SetIdentityId(val *string)
	IdentityIdInput() *string
	InternalValue() *IothubFileUpload
	SetInternalValue(val *IothubFileUpload)
	LockDuration() *string
	SetLockDuration(val *string)
	LockDurationInput() *string
	MaxDeliveryCount() *float64
	SetMaxDeliveryCount(val *float64)
	MaxDeliveryCountInput() *float64
	Notifications() interface{}
	SetNotifications(val interface{})
	NotificationsInput() interface{}
	SasTtl() *string
	SetSasTtl(val *string)
	SasTtlInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAuthenticationType()
	ResetDefaultTtl()
	ResetIdentityId()
	ResetLockDuration()
	ResetMaxDeliveryCount()
	ResetNotifications()
	ResetSasTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IothubFileUploadOutputReference
type jsiiProxy_IothubFileUploadOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IothubFileUploadOutputReference) AuthenticationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) AuthenticationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) ConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) ConnectionStringInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionStringInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) ContainerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) ContainerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) DefaultTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) DefaultTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) IdentityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) IdentityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"identityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) InternalValue() *IothubFileUpload {
	var returns *IothubFileUpload
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) LockDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) LockDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lockDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) MaxDeliveryCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) MaxDeliveryCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxDeliveryCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) Notifications() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) NotificationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notificationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) SasTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) SasTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IothubFileUploadOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIothubFileUploadOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IothubFileUploadOutputReference {
	_init_.Initialize()

	if err := validateNewIothubFileUploadOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IothubFileUploadOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothub.IothubFileUploadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIothubFileUploadOutputReference_Override(i IothubFileUploadOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iothub.IothubFileUploadOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetAuthenticationType(val *string) {
	if err := j.validateSetAuthenticationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationType",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetConnectionString(val *string) {
	if err := j.validateSetConnectionStringParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionString",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetContainerName(val *string) {
	if err := j.validateSetContainerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerName",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetDefaultTtl(val *string) {
	if err := j.validateSetDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTtl",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetIdentityId(val *string) {
	if err := j.validateSetIdentityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"identityId",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetInternalValue(val *IothubFileUpload) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetLockDuration(val *string) {
	if err := j.validateSetLockDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockDuration",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetMaxDeliveryCount(val *float64) {
	if err := j.validateSetMaxDeliveryCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxDeliveryCount",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetNotifications(val interface{}) {
	if err := j.validateSetNotificationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notifications",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetSasTtl(val *string) {
	if err := j.validateSetSasTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sasTtl",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IothubFileUploadOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubFileUploadOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (i *jsiiProxy_IothubFileUploadOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubFileUploadOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetAuthenticationType() {
	_jsii_.InvokeVoid(
		i,
		"resetAuthenticationType",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetDefaultTtl() {
	_jsii_.InvokeVoid(
		i,
		"resetDefaultTtl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetIdentityId() {
	_jsii_.InvokeVoid(
		i,
		"resetIdentityId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetLockDuration() {
	_jsii_.InvokeVoid(
		i,
		"resetLockDuration",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetMaxDeliveryCount() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxDeliveryCount",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetNotifications() {
	_jsii_.InvokeVoid(
		i,
		"resetNotifications",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ResetSasTtl() {
	_jsii_.InvokeVoid(
		i,
		"resetSasTtl",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IothubFileUploadOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IothubFileUploadOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

