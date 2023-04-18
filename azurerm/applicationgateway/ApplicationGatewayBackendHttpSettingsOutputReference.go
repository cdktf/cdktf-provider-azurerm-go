package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewayBackendHttpSettingsOutputReference interface {
	cdktf.ComplexObject
	AffinityCookieName() *string
	SetAffinityCookieName(val *string)
	AffinityCookieNameInput() *string
	AuthenticationCertificate() ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList
	AuthenticationCertificateInput() interface{}
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
	ConnectionDraining() ApplicationGatewayBackendHttpSettingsConnectionDrainingOutputReference
	ConnectionDrainingInput() *ApplicationGatewayBackendHttpSettingsConnectionDraining
	CookieBasedAffinity() *string
	SetCookieBasedAffinity(val *string)
	CookieBasedAffinityInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	PickHostNameFromBackendAddress() interface{}
	SetPickHostNameFromBackendAddress(val interface{})
	PickHostNameFromBackendAddressInput() interface{}
	Port() *float64
	SetPort(val *float64)
	PortInput() *float64
	ProbeId() *string
	ProbeName() *string
	SetProbeName(val *string)
	ProbeNameInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
	RequestTimeout() *float64
	SetRequestTimeout(val *float64)
	RequestTimeoutInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedRootCertificateNames() *[]*string
	SetTrustedRootCertificateNames(val *[]*string)
	TrustedRootCertificateNamesInput() *[]*string
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
	PutAuthenticationCertificate(value interface{})
	PutConnectionDraining(value *ApplicationGatewayBackendHttpSettingsConnectionDraining)
	ResetAffinityCookieName()
	ResetAuthenticationCertificate()
	ResetConnectionDraining()
	ResetHostName()
	ResetPath()
	ResetPickHostNameFromBackendAddress()
	ResetProbeName()
	ResetRequestTimeout()
	ResetTrustedRootCertificateNames()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewayBackendHttpSettingsOutputReference
type jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) AffinityCookieName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinityCookieName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) AffinityCookieNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"affinityCookieNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) AuthenticationCertificate() ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList {
	var returns ApplicationGatewayBackendHttpSettingsAuthenticationCertificateList
	_jsii_.Get(
		j,
		"authenticationCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) AuthenticationCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authenticationCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ConnectionDraining() ApplicationGatewayBackendHttpSettingsConnectionDrainingOutputReference {
	var returns ApplicationGatewayBackendHttpSettingsConnectionDrainingOutputReference
	_jsii_.Get(
		j,
		"connectionDraining",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ConnectionDrainingInput() *ApplicationGatewayBackendHttpSettingsConnectionDraining {
	var returns *ApplicationGatewayBackendHttpSettingsConnectionDraining
	_jsii_.Get(
		j,
		"connectionDrainingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) CookieBasedAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBasedAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) CookieBasedAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieBasedAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) PickHostNameFromBackendAddress() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pickHostNameFromBackendAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) PickHostNameFromBackendAddressInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"pickHostNameFromBackendAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Port() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"port",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) PortInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"portInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ProbeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ProbeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ProbeNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"probeNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) TrustedRootCertificateNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedRootCertificateNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) TrustedRootCertificateNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedRootCertificateNamesInput",
		&returns,
	)
	return returns
}


func NewApplicationGatewayBackendHttpSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewayBackendHttpSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewayBackendHttpSettingsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayBackendHttpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewayBackendHttpSettingsOutputReference_Override(a ApplicationGatewayBackendHttpSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewayBackendHttpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetAffinityCookieName(val *string) {
	if err := j.validateSetAffinityCookieNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"affinityCookieName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetCookieBasedAffinity(val *string) {
	if err := j.validateSetCookieBasedAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieBasedAffinity",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetPickHostNameFromBackendAddress(val interface{}) {
	if err := j.validateSetPickHostNameFromBackendAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pickHostNameFromBackendAddress",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetPort(val *float64) {
	if err := j.validateSetPortParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"port",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetProbeName(val *string) {
	if err := j.validateSetProbeNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"probeName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetRequestTimeout(val *float64) {
	if err := j.validateSetRequestTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference)SetTrustedRootCertificateNames(val *[]*string) {
	if err := j.validateSetTrustedRootCertificateNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedRootCertificateNames",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) PutAuthenticationCertificate(value interface{}) {
	if err := a.validatePutAuthenticationCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAuthenticationCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) PutConnectionDraining(value *ApplicationGatewayBackendHttpSettingsConnectionDraining) {
	if err := a.validatePutConnectionDrainingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putConnectionDraining",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetAffinityCookieName() {
	_jsii_.InvokeVoid(
		a,
		"resetAffinityCookieName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetAuthenticationCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetConnectionDraining() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionDraining",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetHostName() {
	_jsii_.InvokeVoid(
		a,
		"resetHostName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetPickHostNameFromBackendAddress() {
	_jsii_.InvokeVoid(
		a,
		"resetPickHostNameFromBackendAddress",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetProbeName() {
	_jsii_.InvokeVoid(
		a,
		"resetProbeName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ResetTrustedRootCertificateNames() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedRootCertificateNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationGatewayBackendHttpSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

