package hdinsighthadoopcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/hdinsighthadoopcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightHadoopClusterRolesEdgeNodeOutputReference interface {
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
	HttpsEndpoints() HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList
	HttpsEndpointsInput() interface{}
	InstallScriptAction() HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList
	InstallScriptActionInput() interface{}
	InternalValue() *HdinsightHadoopClusterRolesEdgeNode
	SetInternalValue(val *HdinsightHadoopClusterRolesEdgeNode)
	TargetInstanceCount() *float64
	SetTargetInstanceCount(val *float64)
	TargetInstanceCountInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UninstallScriptActions() HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList
	UninstallScriptActionsInput() interface{}
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutHttpsEndpoints(value interface{})
	PutInstallScriptAction(value interface{})
	PutUninstallScriptActions(value interface{})
	ResetHttpsEndpoints()
	ResetUninstallScriptActions()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightHadoopClusterRolesEdgeNodeOutputReference
type jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) HttpsEndpoints() HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList {
	var returns HdinsightHadoopClusterRolesEdgeNodeHttpsEndpointsList
	_jsii_.Get(
		j,
		"httpsEndpoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) HttpsEndpointsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpsEndpointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InstallScriptAction() HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList {
	var returns HdinsightHadoopClusterRolesEdgeNodeInstallScriptActionList
	_jsii_.Get(
		j,
		"installScriptAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InstallScriptActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"installScriptActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InternalValue() *HdinsightHadoopClusterRolesEdgeNode {
	var returns *HdinsightHadoopClusterRolesEdgeNode
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TargetInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TargetInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"targetInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) UninstallScriptActions() HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList {
	var returns HdinsightHadoopClusterRolesEdgeNodeUninstallScriptActionsList
	_jsii_.Get(
		j,
		"uninstallScriptActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) UninstallScriptActionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"uninstallScriptActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewHdinsightHadoopClusterRolesEdgeNodeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightHadoopClusterRolesEdgeNodeOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightHadoopClusterRolesEdgeNodeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightHadoopClusterRolesEdgeNodeOutputReference_Override(h HdinsightHadoopClusterRolesEdgeNodeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightHadoopCluster.HdinsightHadoopClusterRolesEdgeNodeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetInternalValue(val *HdinsightHadoopClusterRolesEdgeNode) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetTargetInstanceCount(val *float64) {
	if err := j.validateSetTargetInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"targetInstanceCount",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) PutHttpsEndpoints(value interface{}) {
	if err := h.validatePutHttpsEndpointsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putHttpsEndpoints",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) PutInstallScriptAction(value interface{}) {
	if err := h.validatePutInstallScriptActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putInstallScriptAction",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) PutUninstallScriptActions(value interface{}) {
	if err := h.validatePutUninstallScriptActionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putUninstallScriptActions",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ResetHttpsEndpoints() {
	_jsii_.InvokeVoid(
		h,
		"resetHttpsEndpoints",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ResetUninstallScriptActions() {
	_jsii_.InvokeVoid(
		h,
		"resetUninstallScriptActions",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HdinsightHadoopClusterRolesEdgeNodeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

