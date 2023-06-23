package activedirectorydomainservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/activedirectorydomainservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActiveDirectoryDomainServiceSecureLdapOutputReference interface {
	cdktf.ComplexObject
	CertificateExpiry() *string
	CertificateThumbprint() *string
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExternalAccessEnabled() interface{}
	SetExternalAccessEnabled(val interface{})
	ExternalAccessEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *ActiveDirectoryDomainServiceSecureLdap
	SetInternalValue(val *ActiveDirectoryDomainServiceSecureLdap)
	PfxCertificate() *string
	SetPfxCertificate(val *string)
	PfxCertificateInput() *string
	PfxCertificatePassword() *string
	SetPfxCertificatePassword(val *string)
	PfxCertificatePasswordInput() *string
	PublicCertificate() *string
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
	ResetExternalAccessEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActiveDirectoryDomainServiceSecureLdapOutputReference
type jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) CertificateExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) CertificateThumbprint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateThumbprint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ExternalAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ExternalAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) InternalValue() *ActiveDirectoryDomainServiceSecureLdap {
	var returns *ActiveDirectoryDomainServiceSecureLdap
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PfxCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfxCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) PublicCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewActiveDirectoryDomainServiceSecureLdapOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActiveDirectoryDomainServiceSecureLdapOutputReference {
	_init_.Initialize()

	if err := validateNewActiveDirectoryDomainServiceSecureLdapOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecureLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActiveDirectoryDomainServiceSecureLdapOutputReference_Override(a ActiveDirectoryDomainServiceSecureLdapOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.activeDirectoryDomainService.ActiveDirectoryDomainServiceSecureLdapOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetExternalAccessEnabled(val interface{}) {
	if err := j.validateSetExternalAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"externalAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetInternalValue(val *ActiveDirectoryDomainServiceSecureLdap) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetPfxCertificate(val *string) {
	if err := j.validateSetPfxCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfxCertificate",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetPfxCertificatePassword(val *string) {
	if err := j.validateSetPfxCertificatePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfxCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ResetExternalAccessEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetExternalAccessEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ActiveDirectoryDomainServiceSecureLdapOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

