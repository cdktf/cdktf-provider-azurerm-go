package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayFrontendIpConfigurationOutputReference interface {
	cdktf.ComplexObject
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddressAllocation() *string
	SetPrivateIpAddressAllocation(val *string)
	PrivateIpAddressAllocationInput() *string
	PrivateIpAddressInput() *string
	PrivateLinkConfigurationId() *string
	PrivateLinkConfigurationName() *string
	SetPrivateLinkConfigurationName(val *string)
	PrivateLinkConfigurationNameInput() *string
	PublicIpAddressId() *string
	SetPublicIpAddressId(val *string)
	PublicIpAddressIdInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
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
	ResetPrivateIpAddress()
	ResetPrivateIpAddressAllocation()
	ResetPrivateLinkConfigurationName()
	ResetPublicIpAddressId()
	ResetSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayFrontendIpConfigurationOutputReference
type jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateIpAddressAllocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateIpAddressAllocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateLinkConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateLinkConfigurationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkConfigurationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PrivateLinkConfigurationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateLinkConfigurationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) PublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewayFrontendIpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayFrontendIpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayFrontendIpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayFrontendIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayFrontendIpConfigurationOutputReference_Override(a ApplicationGatewayFrontendIpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayFrontendIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetPrivateIpAddressAllocation(val *string) {
	if err := j.validateSetPrivateIpAddressAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressAllocation",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetPrivateLinkConfigurationName(val *string) {
	if err := j.validateSetPrivateLinkConfigurationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateLinkConfigurationName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetPublicIpAddressId(val *string) {
	if err := j.validateSetPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ResetPrivateIpAddressAllocation() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateIpAddressAllocation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ResetPrivateLinkConfigurationName() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateLinkConfigurationName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ResetPublicIpAddressId() {
	_jsii_.InvokeVoid(
		a,
		"resetPublicIpAddressId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayFrontendIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

