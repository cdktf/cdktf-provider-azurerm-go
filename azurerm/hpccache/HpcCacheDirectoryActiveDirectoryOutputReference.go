package hpccache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/hpccache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HpcCacheDirectoryActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	CacheNetbiosName() *string
	SetCacheNetbiosName(val *string)
	CacheNetbiosNameInput() *string
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
	DnsPrimaryIp() *string
	SetDnsPrimaryIp(val *string)
	DnsPrimaryIpInput() *string
	DnsSecondaryIp() *string
	SetDnsSecondaryIp(val *string)
	DnsSecondaryIpInput() *string
	DomainName() *string
	SetDomainName(val *string)
	DomainNameInput() *string
	DomainNetbiosName() *string
	SetDomainNetbiosName(val *string)
	DomainNetbiosNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *HpcCacheDirectoryActiveDirectory
	SetInternalValue(val *HpcCacheDirectoryActiveDirectory)
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetDnsSecondaryIp()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDirectoryActiveDirectoryOutputReference
type jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) CacheNetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNetbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) CacheNetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheNetbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsPrimaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrimaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsPrimaryIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsPrimaryIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsSecondaryIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSecondaryIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DnsSecondaryIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsSecondaryIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainNetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNetbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) DomainNetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainNetbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) InternalValue() *HpcCacheDirectoryActiveDirectory {
	var returns *HpcCacheDirectoryActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewHpcCacheDirectoryActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDirectoryActiveDirectoryOutputReference {
	_init_.Initialize()

	if err := validateNewHpcCacheDirectoryActiveDirectoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDirectoryActiveDirectoryOutputReference_Override(h HpcCacheDirectoryActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetCacheNetbiosName(val *string) {
	if err := j.validateSetCacheNetbiosNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheNetbiosName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetDnsPrimaryIp(val *string) {
	if err := j.validateSetDnsPrimaryIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsPrimaryIp",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetDnsSecondaryIp(val *string) {
	if err := j.validateSetDnsSecondaryIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSecondaryIp",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetDomainName(val *string) {
	if err := j.validateSetDomainNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetDomainNetbiosName(val *string) {
	if err := j.validateSetDomainNetbiosNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domainNetbiosName",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetInternalValue(val *HpcCacheDirectoryActiveDirectory) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ResetDnsSecondaryIp() {
	_jsii_.InvokeVoid(
		h,
		"resetDnsSecondaryIp",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

