package arckubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/arckubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference interface {
	cdktf.ComplexObject
	ClientCertificateBase64() *string
	SetClientCertificateBase64(val *string)
	ClientCertificateBase64Input() *string
	ClientCertificatePassword() *string
	SetClientCertificatePassword(val *string)
	ClientCertificatePasswordInput() *string
	ClientCertificateSendChain() interface{}
	SetClientCertificateSendChain(val interface{})
	ClientCertificateSendChainInput() interface{}
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretInput() *string
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
	InternalValue() *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal
	SetInternalValue(val *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal)
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
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
	ResetClientCertificateBase64()
	ResetClientCertificatePassword()
	ResetClientCertificateSendChain()
	ResetClientSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference
type jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientCertificateBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientCertificateBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientCertificateSendChain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateSendChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientCertificateSendChainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientCertificateSendChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) InternalValue() *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal {
	var returns *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference {
	_init_.Initialize()

	if err := validateNewArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference_Override(a ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetClientCertificateBase64(val *string) {
	if err := j.validateSetClientCertificateBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateBase64",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetClientCertificatePassword(val *string) {
	if err := j.validateSetClientCertificatePasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetClientCertificateSendChain(val interface{}) {
	if err := j.validateSetClientCertificateSendChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientCertificateSendChain",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetClientId(val *string) {
	if err := j.validateSetClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetClientSecret(val *string) {
	if err := j.validateSetClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetInternalValue(val *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetTenantId(val *string) {
	if err := j.validateSetTenantIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ResetClientCertificateBase64() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateBase64",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ResetClientCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ResetClientCertificateSendChain() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificateSendChain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

