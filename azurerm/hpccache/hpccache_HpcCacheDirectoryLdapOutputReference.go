package hpccache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/hpccache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HpcCacheDirectoryLdapOutputReference interface {
	cdktf.ComplexObject
	BaseDn() *string
	SetBaseDn(val *string)
	BaseDnInput() *string
	Bind() HpcCacheDirectoryLdapBindOutputReference
	BindInput() *HpcCacheDirectoryLdapBind
	CertificateValidationUri() *string
	SetCertificateValidationUri(val *string)
	CertificateValidationUriInput() *string
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
	DownloadCertificateAutomatically() interface{}
	SetDownloadCertificateAutomatically(val interface{})
	DownloadCertificateAutomaticallyInput() interface{}
	Encrypted() interface{}
	SetEncrypted(val interface{})
	EncryptedInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *HpcCacheDirectoryLdap
	SetInternalValue(val *HpcCacheDirectoryLdap)
	Server() *string
	SetServer(val *string)
	ServerInput() *string
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
	PutBind(value *HpcCacheDirectoryLdapBind)
	ResetBind()
	ResetCertificateValidationUri()
	ResetDownloadCertificateAutomatically()
	ResetEncrypted()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDirectoryLdapOutputReference
type jsiiProxy_HpcCacheDirectoryLdapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) BaseDn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) BaseDnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseDnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Bind() HpcCacheDirectoryLdapBindOutputReference {
	var returns HpcCacheDirectoryLdapBindOutputReference
	_jsii_.Get(
		j,
		"bind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) BindInput() *HpcCacheDirectoryLdapBind {
	var returns *HpcCacheDirectoryLdapBind
	_jsii_.Get(
		j,
		"bindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) CertificateValidationUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateValidationUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) CertificateValidationUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateValidationUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) DownloadCertificateAutomatically() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadCertificateAutomatically",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) DownloadCertificateAutomaticallyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"downloadCertificateAutomaticallyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Encrypted() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encrypted",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) EncryptedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) InternalValue() *HpcCacheDirectoryLdap {
	var returns *HpcCacheDirectoryLdap
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Server() *string {
	var returns *string
	_jsii_.Get(
		j,
		"server",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ServerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDirectoryLdapOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HpcCacheDirectoryLdapOutputReference {
	_init_.Initialize()

	if err := validateNewHpcCacheDirectoryLdapOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HpcCacheDirectoryLdapOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHpcCacheDirectoryLdapOutputReference_Override(h HpcCacheDirectoryLdapOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDirectoryLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetBaseDn(val *string) {
	if err := j.validateSetBaseDnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseDn",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetCertificateValidationUri(val *string) {
	if err := j.validateSetCertificateValidationUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateValidationUri",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetDownloadCertificateAutomatically(val interface{}) {
	if err := j.validateSetDownloadCertificateAutomaticallyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"downloadCertificateAutomatically",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetEncrypted(val interface{}) {
	if err := j.validateSetEncryptedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encrypted",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetInternalValue(val *HpcCacheDirectoryLdap) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetServer(val *string) {
	if err := j.validateSetServerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"server",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDirectoryLdapOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) PutBind(value *HpcCacheDirectoryLdapBind) {
	if err := h.validatePutBindParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putBind",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetBind() {
	_jsii_.InvokeVoid(
		h,
		"resetBind",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetCertificateValidationUri() {
	_jsii_.InvokeVoid(
		h,
		"resetCertificateValidationUri",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetDownloadCertificateAutomatically() {
	_jsii_.InvokeVoid(
		h,
		"resetDownloadCertificateAutomatically",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ResetEncrypted() {
	_jsii_.InvokeVoid(
		h,
		"resetEncrypted",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HpcCacheDirectoryLdapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

