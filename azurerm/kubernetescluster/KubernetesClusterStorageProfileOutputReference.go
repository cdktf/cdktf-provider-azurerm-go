package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterStorageProfileOutputReference interface {
	cdktf.ComplexObject
	BlobDriverEnabled() interface{}
	SetBlobDriverEnabled(val interface{})
	BlobDriverEnabledInput() interface{}
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
	DiskDriverEnabled() interface{}
	SetDiskDriverEnabled(val interface{})
	DiskDriverEnabledInput() interface{}
	DiskDriverVersion() *string
	SetDiskDriverVersion(val *string)
	DiskDriverVersionInput() *string
	FileDriverEnabled() interface{}
	SetFileDriverEnabled(val interface{})
	FileDriverEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterStorageProfile
	SetInternalValue(val *KubernetesClusterStorageProfile)
	SnapshotControllerEnabled() interface{}
	SetSnapshotControllerEnabled(val interface{})
	SnapshotControllerEnabledInput() interface{}
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
	ResetBlobDriverEnabled()
	ResetDiskDriverEnabled()
	ResetDiskDriverVersion()
	ResetFileDriverEnabled()
	ResetSnapshotControllerEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterStorageProfileOutputReference
type jsiiProxy_KubernetesClusterStorageProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) BlobDriverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blobDriverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) BlobDriverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blobDriverEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) DiskDriverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskDriverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) DiskDriverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"diskDriverEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) DiskDriverVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskDriverVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) DiskDriverVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskDriverVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) FileDriverEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileDriverEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) FileDriverEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileDriverEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) InternalValue() *KubernetesClusterStorageProfile {
	var returns *KubernetesClusterStorageProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) SnapshotControllerEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotControllerEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) SnapshotControllerEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snapshotControllerEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterStorageProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterStorageProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterStorageProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterStorageProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterStorageProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterStorageProfileOutputReference_Override(k KubernetesClusterStorageProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterStorageProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetBlobDriverEnabled(val interface{}) {
	if err := j.validateSetBlobDriverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blobDriverEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetDiskDriverEnabled(val interface{}) {
	if err := j.validateSetDiskDriverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskDriverEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetDiskDriverVersion(val *string) {
	if err := j.validateSetDiskDriverVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskDriverVersion",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetFileDriverEnabled(val interface{}) {
	if err := j.validateSetFileDriverEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileDriverEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetInternalValue(val *KubernetesClusterStorageProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetSnapshotControllerEnabled(val interface{}) {
	if err := j.validateSetSnapshotControllerEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snapshotControllerEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterStorageProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ResetBlobDriverEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetBlobDriverEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ResetDiskDriverEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetDiskDriverEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ResetDiskDriverVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetDiskDriverVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ResetFileDriverEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetFileDriverEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ResetSnapshotControllerEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetSnapshotControllerEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterStorageProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

