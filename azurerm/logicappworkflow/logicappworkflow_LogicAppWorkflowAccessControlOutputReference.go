package logicappworkflow

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/logicappworkflow/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppWorkflowAccessControlOutputReference interface {
	cdktf.ComplexObject
	Action() LogicAppWorkflowAccessControlActionOutputReference
	ActionInput() *LogicAppWorkflowAccessControlAction
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
	Content() LogicAppWorkflowAccessControlContentOutputReference
	ContentInput() *LogicAppWorkflowAccessControlContent
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *LogicAppWorkflowAccessControl
	SetInternalValue(val *LogicAppWorkflowAccessControl)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Trigger() LogicAppWorkflowAccessControlTriggerOutputReference
	TriggerInput() *LogicAppWorkflowAccessControlTrigger
	WorkflowManagement() LogicAppWorkflowAccessControlWorkflowManagementOutputReference
	WorkflowManagementInput() *LogicAppWorkflowAccessControlWorkflowManagement
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
	PutAction(value *LogicAppWorkflowAccessControlAction)
	PutContent(value *LogicAppWorkflowAccessControlContent)
	PutTrigger(value *LogicAppWorkflowAccessControlTrigger)
	PutWorkflowManagement(value *LogicAppWorkflowAccessControlWorkflowManagement)
	ResetAction()
	ResetContent()
	ResetTrigger()
	ResetWorkflowManagement()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LogicAppWorkflowAccessControlOutputReference
type jsiiProxy_LogicAppWorkflowAccessControlOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) Action() LogicAppWorkflowAccessControlActionOutputReference {
	var returns LogicAppWorkflowAccessControlActionOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ActionInput() *LogicAppWorkflowAccessControlAction {
	var returns *LogicAppWorkflowAccessControlAction
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) Content() LogicAppWorkflowAccessControlContentOutputReference {
	var returns LogicAppWorkflowAccessControlContentOutputReference
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ContentInput() *LogicAppWorkflowAccessControlContent {
	var returns *LogicAppWorkflowAccessControlContent
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) InternalValue() *LogicAppWorkflowAccessControl {
	var returns *LogicAppWorkflowAccessControl
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) Trigger() LogicAppWorkflowAccessControlTriggerOutputReference {
	var returns LogicAppWorkflowAccessControlTriggerOutputReference
	_jsii_.Get(
		j,
		"trigger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) TriggerInput() *LogicAppWorkflowAccessControlTrigger {
	var returns *LogicAppWorkflowAccessControlTrigger
	_jsii_.Get(
		j,
		"triggerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) WorkflowManagement() LogicAppWorkflowAccessControlWorkflowManagementOutputReference {
	var returns LogicAppWorkflowAccessControlWorkflowManagementOutputReference
	_jsii_.Get(
		j,
		"workflowManagement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) WorkflowManagementInput() *LogicAppWorkflowAccessControlWorkflowManagement {
	var returns *LogicAppWorkflowAccessControlWorkflowManagement
	_jsii_.Get(
		j,
		"workflowManagementInput",
		&returns,
	)
	return returns
}


func NewLogicAppWorkflowAccessControlOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogicAppWorkflowAccessControlOutputReference {
	_init_.Initialize()

	if err := validateNewLogicAppWorkflowAccessControlOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppWorkflowAccessControlOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppWorkflow.LogicAppWorkflowAccessControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogicAppWorkflowAccessControlOutputReference_Override(l LogicAppWorkflowAccessControlOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppWorkflow.LogicAppWorkflowAccessControlOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference)SetInternalValue(val *LogicAppWorkflowAccessControl) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppWorkflowAccessControlOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) PutAction(value *LogicAppWorkflowAccessControlAction) {
	if err := l.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) PutContent(value *LogicAppWorkflowAccessControlContent) {
	if err := l.validatePutContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putContent",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) PutTrigger(value *LogicAppWorkflowAccessControlTrigger) {
	if err := l.validatePutTriggerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTrigger",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) PutWorkflowManagement(value *LogicAppWorkflowAccessControlWorkflowManagement) {
	if err := l.validatePutWorkflowManagementParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putWorkflowManagement",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ResetAction() {
	_jsii_.InvokeVoid(
		l,
		"resetAction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ResetContent() {
	_jsii_.InvokeVoid(
		l,
		"resetContent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ResetTrigger() {
	_jsii_.InvokeVoid(
		l,
		"resetTrigger",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ResetWorkflowManagement() {
	_jsii_.InvokeVoid(
		l,
		"resetWorkflowManagement",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppWorkflowAccessControlOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

