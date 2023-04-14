package machinelearninginferencecluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/machinelearninginferencecluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningInferenceClusterSslOutputReference interface {
	cdktf.ComplexObject
	Cert() *string
	SetCert(val *string)
	CertInput() *string
	Cname() *string
	SetCname(val *string)
	CnameInput() *string
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
	InternalValue() *MachineLearningInferenceClusterSsl
	SetInternalValue(val *MachineLearningInferenceClusterSsl)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	LeafDomainLabel() *string
	SetLeafDomainLabel(val *string)
	LeafDomainLabelInput() *string
	OverwriteExistingDomain() interface{}
	SetOverwriteExistingDomain(val interface{})
	OverwriteExistingDomainInput() interface{}
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
	ResetCert()
	ResetCname()
	ResetKey()
	ResetLeafDomainLabel()
	ResetOverwriteExistingDomain()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MachineLearningInferenceClusterSslOutputReference
type jsiiProxy_MachineLearningInferenceClusterSslOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) Cert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) CertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) Cname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) CnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) InternalValue() *MachineLearningInferenceClusterSsl {
	var returns *MachineLearningInferenceClusterSsl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) LeafDomainLabel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leafDomainLabel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) LeafDomainLabelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leafDomainLabelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) OverwriteExistingDomain() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteExistingDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) OverwriteExistingDomainInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteExistingDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMachineLearningInferenceClusterSslOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MachineLearningInferenceClusterSslOutputReference {
	_init_.Initialize()

	if err := validateNewMachineLearningInferenceClusterSslOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MachineLearningInferenceClusterSslOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningInferenceCluster.MachineLearningInferenceClusterSslOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMachineLearningInferenceClusterSslOutputReference_Override(m MachineLearningInferenceClusterSslOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.machineLearningInferenceCluster.MachineLearningInferenceClusterSslOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetCert(val *string) {
	if err := j.validateSetCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cert",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetCname(val *string) {
	if err := j.validateSetCnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cname",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetInternalValue(val *MachineLearningInferenceClusterSsl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetLeafDomainLabel(val *string) {
	if err := j.validateSetLeafDomainLabelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafDomainLabel",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetOverwriteExistingDomain(val interface{}) {
	if err := j.validateSetOverwriteExistingDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteExistingDomain",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MachineLearningInferenceClusterSslOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ResetCert() {
	_jsii_.InvokeVoid(
		m,
		"resetCert",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ResetCname() {
	_jsii_.InvokeVoid(
		m,
		"resetCname",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		m,
		"resetKey",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ResetLeafDomainLabel() {
	_jsii_.InvokeVoid(
		m,
		"resetLeafDomainLabel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ResetOverwriteExistingDomain() {
	_jsii_.InvokeVoid(
		m,
		"resetOverwriteExistingDomain",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MachineLearningInferenceClusterSslOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

