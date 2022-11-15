package apimanagement

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/apimanagement/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementHostnameConfigurationProxyOutputReference interface {
	cdktf.ComplexObject
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	CertificatePassword() *string
	SetCertificatePassword(val *string)
	CertificatePasswordInput() *string
	CertificateSource() *string
	CertificateStatus() *string
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
	DefaultSslBinding() interface{}
	SetDefaultSslBinding(val interface{})
	DefaultSslBindingInput() interface{}
	Expiry() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	KeyVaultId() *string
	SetKeyVaultId(val *string)
	KeyVaultIdInput() *string
	NegotiateClientCertificate() interface{}
	SetNegotiateClientCertificate(val interface{})
	NegotiateClientCertificateInput() interface{}
	SslKeyvaultIdentityClientId() *string
	SetSslKeyvaultIdentityClientId(val *string)
	SslKeyvaultIdentityClientIdInput() *string
	Subject() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Thumbprint() *string
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
	ResetCertificate()
	ResetCertificatePassword()
	ResetDefaultSslBinding()
	ResetKeyVaultId()
	ResetNegotiateClientCertificate()
	ResetSslKeyvaultIdentityClientId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementHostnameConfigurationProxyOutputReference
type jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificateSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CertificateStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) DefaultSslBinding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSslBinding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) DefaultSslBindingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultSslBindingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) KeyVaultId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) KeyVaultIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyVaultIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) NegotiateClientCertificate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) NegotiateClientCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"negotiateClientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SslKeyvaultIdentityClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) SslKeyvaultIdentityClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sslKeyvaultIdentityClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Subject() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subject",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Thumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"thumbprint",
		&returns,
	)
	return returns
}


func NewApiManagementHostnameConfigurationProxyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApiManagementHostnameConfigurationProxyOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementHostnameConfigurationProxyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationProxyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApiManagementHostnameConfigurationProxyOutputReference_Override(a ApiManagementHostnameConfigurationProxyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagement.ApiManagementHostnameConfigurationProxyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetCertificatePassword(val *string) {
	if err := j.validateSetCertificatePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificatePassword",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetDefaultSslBinding(val interface{}) {
	if err := j.validateSetDefaultSslBindingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultSslBinding",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetKeyVaultId(val *string) {
	if err := j.validateSetKeyVaultIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyVaultId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetNegotiateClientCertificate(val interface{}) {
	if err := j.validateSetNegotiateClientCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"negotiateClientCertificate",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetSslKeyvaultIdentityClientId(val *string) {
	if err := j.validateSetSslKeyvaultIdentityClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sslKeyvaultIdentityClientId",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetDefaultSslBinding() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultSslBinding",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetKeyVaultId() {
	_jsii_.InvokeVoid(
		a,
		"resetKeyVaultId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetNegotiateClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetNegotiateClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ResetSslKeyvaultIdentityClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetSslKeyvaultIdentityClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiManagementHostnameConfigurationProxyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

