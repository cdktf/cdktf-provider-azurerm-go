package hdinsightkafkacluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/hdinsightkafkacluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList
type jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewHdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList {
	_init_.Initialize()

	if err := validateNewHdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewHdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList_Override(h HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		h,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) Get(index *float64) HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsOutputReference {
	if err := h.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsOutputReference

	_jsii_.Invoke(
		h,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (h *jsiiProxy_HdinsightKafkaClusterRolesKafkaManagementNodeScriptActionsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
