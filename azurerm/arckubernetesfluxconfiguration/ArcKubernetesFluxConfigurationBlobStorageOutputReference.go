package arckubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/arckubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ArcKubernetesFluxConfigurationBlobStorageOutputReference interface {
	cdktf.ComplexObject
	AccountKey() *string
	SetAccountKey(val *string)
	AccountKeyInput() *string
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
	ContainerId() *string
	SetContainerId(val *string)
	ContainerIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ArcKubernetesFluxConfigurationBlobStorage
	SetInternalValue(val *ArcKubernetesFluxConfigurationBlobStorage)
	LocalAuthReference() *string
	SetLocalAuthReference(val *string)
	LocalAuthReferenceInput() *string
	SasToken() *string
	SetSasToken(val *string)
	SasTokenInput() *string
	ServicePrincipal() ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference
	ServicePrincipalInput() *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal
	SyncIntervalInSeconds() *float64
	SetSyncIntervalInSeconds(val *float64)
	SyncIntervalInSecondsInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeoutInSeconds() *float64
	SetTimeoutInSeconds(val *float64)
	TimeoutInSecondsInput() *float64
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
	PutServicePrincipal(value *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal)
	ResetAccountKey()
	ResetLocalAuthReference()
	ResetSasToken()
	ResetServicePrincipal()
	ResetSyncIntervalInSeconds()
	ResetTimeoutInSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ArcKubernetesFluxConfigurationBlobStorageOutputReference
type jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) AccountKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) AccountKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accountKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) InternalValue() *ArcKubernetesFluxConfigurationBlobStorage {
	var returns *ArcKubernetesFluxConfigurationBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) LocalAuthReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) LocalAuthReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) SasToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) SasTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ServicePrincipal() ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference {
	var returns ArcKubernetesFluxConfigurationBlobStorageServicePrincipalOutputReference
	_jsii_.Get(
		j,
		"servicePrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ServicePrincipalInput() *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal {
	var returns *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal
	_jsii_.Get(
		j,
		"servicePrincipalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}


func NewArcKubernetesFluxConfigurationBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ArcKubernetesFluxConfigurationBlobStorageOutputReference {
	_init_.Initialize()

	if err := validateNewArcKubernetesFluxConfigurationBlobStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewArcKubernetesFluxConfigurationBlobStorageOutputReference_Override(a ArcKubernetesFluxConfigurationBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.arcKubernetesFluxConfiguration.ArcKubernetesFluxConfigurationBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetAccountKey(val *string) {
	if err := j.validateSetAccountKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accountKey",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetContainerId(val *string) {
	if err := j.validateSetContainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerId",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetInternalValue(val *ArcKubernetesFluxConfigurationBlobStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetLocalAuthReference(val *string) {
	if err := j.validateSetLocalAuthReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthReference",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetSasToken(val *string) {
	if err := j.validateSetSasTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sasToken",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) PutServicePrincipal(value *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal) {
	if err := a.validatePutServicePrincipalParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServicePrincipal",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ResetAccountKey() {
	_jsii_.InvokeVoid(
		a,
		"resetAccountKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ResetLocalAuthReference() {
	_jsii_.InvokeVoid(
		a,
		"resetLocalAuthReference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ResetSasToken() {
	_jsii_.InvokeVoid(
		a,
		"resetSasToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ResetServicePrincipal() {
	_jsii_.InvokeVoid(
		a,
		"resetServicePrincipal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ArcKubernetesFluxConfigurationBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

