package kubernetesfluxconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/kubernetesfluxconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesFluxConfigurationGitRepositoryOutputReference interface {
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
	HttpsCaCertBase64() *string
	SetHttpsCaCertBase64(val *string)
	HttpsCaCertBase64Input() *string
	HttpsKeyBase64() *string
	SetHttpsKeyBase64(val *string)
	HttpsKeyBase64Input() *string
	HttpsUser() *string
	SetHttpsUser(val *string)
	HttpsUserInput() *string
	InternalValue() *KubernetesFluxConfigurationGitRepository
	SetInternalValue(val *KubernetesFluxConfigurationGitRepository)
	LocalAuthReference() *string
	SetLocalAuthReference(val *string)
	LocalAuthReferenceInput() *string
	ReferenceType() *string
	SetReferenceType(val *string)
	ReferenceTypeInput() *string
	ReferenceValue() *string
	SetReferenceValue(val *string)
	ReferenceValueInput() *string
	SshKnownHostsBase64() *string
	SetSshKnownHostsBase64(val *string)
	SshKnownHostsBase64Input() *string
	SshPrivateKeyBase64() *string
	SetSshPrivateKeyBase64(val *string)
	SshPrivateKeyBase64Input() *string
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
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
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
	ResetHttpsCaCertBase64()
	ResetHttpsKeyBase64()
	ResetHttpsUser()
	ResetLocalAuthReference()
	ResetSshKnownHostsBase64()
	ResetSshPrivateKeyBase64()
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

// The jsii proxy struct for KubernetesFluxConfigurationGitRepositoryOutputReference
type jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) HttpsCaCertBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsCaCertBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) HttpsCaCertBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsCaCertBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) HttpsKeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsKeyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) HttpsKeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsKeyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) HttpsUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) HttpsUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) InternalValue() *KubernetesFluxConfigurationGitRepository {
	var returns *KubernetesFluxConfigurationGitRepository
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) LocalAuthReference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) LocalAuthReferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localAuthReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ReferenceValueInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"referenceValueInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) SshKnownHostsBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKnownHostsBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) SshKnownHostsBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshKnownHostsBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) SshPrivateKeyBase64() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPrivateKeyBase64",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) SshPrivateKeyBase64Input() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshPrivateKeyBase64Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) SyncIntervalInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) SyncIntervalInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"syncIntervalInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) TimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) TimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"timeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}


func NewKubernetesFluxConfigurationGitRepositoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesFluxConfigurationGitRepositoryOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesFluxConfigurationGitRepositoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationGitRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesFluxConfigurationGitRepositoryOutputReference_Override(k KubernetesFluxConfigurationGitRepositoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesFluxConfiguration.KubernetesFluxConfigurationGitRepositoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetHttpsCaCertBase64(val *string) {
	if err := j.validateSetHttpsCaCertBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsCaCertBase64",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetHttpsKeyBase64(val *string) {
	if err := j.validateSetHttpsKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsKeyBase64",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetHttpsUser(val *string) {
	if err := j.validateSetHttpsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpsUser",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetInternalValue(val *KubernetesFluxConfigurationGitRepository) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetLocalAuthReference(val *string) {
	if err := j.validateSetLocalAuthReferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localAuthReference",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetReferenceType(val *string) {
	if err := j.validateSetReferenceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceType",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetReferenceValue(val *string) {
	if err := j.validateSetReferenceValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"referenceValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetSshKnownHostsBase64(val *string) {
	if err := j.validateSetSshKnownHostsBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshKnownHostsBase64",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetSshPrivateKeyBase64(val *string) {
	if err := j.validateSetSshPrivateKeyBase64Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sshPrivateKeyBase64",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetSyncIntervalInSeconds(val *float64) {
	if err := j.validateSetSyncIntervalInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"syncIntervalInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetTimeoutInSeconds(val *float64) {
	if err := j.validateSetTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetHttpsCaCertBase64() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpsCaCertBase64",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetHttpsKeyBase64() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpsKeyBase64",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetHttpsUser() {
	_jsii_.InvokeVoid(
		k,
		"resetHttpsUser",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetLocalAuthReference() {
	_jsii_.InvokeVoid(
		k,
		"resetLocalAuthReference",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetSshKnownHostsBase64() {
	_jsii_.InvokeVoid(
		k,
		"resetSshKnownHostsBase64",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetSshPrivateKeyBase64() {
	_jsii_.InvokeVoid(
		k,
		"resetSshPrivateKeyBase64",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetSyncIntervalInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetSyncIntervalInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ResetTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		k,
		"resetTimeoutInSeconds",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesFluxConfigurationGitRepositoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

