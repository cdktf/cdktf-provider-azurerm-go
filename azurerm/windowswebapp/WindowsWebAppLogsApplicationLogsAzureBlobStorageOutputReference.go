package windowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/windowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference interface {
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
	InternalValue() *WindowsWebAppLogsApplicationLogsAzureBlobStorage
	SetInternalValue(val *WindowsWebAppLogsApplicationLogsAzureBlobStorage)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	RetentionInDays() *float64
	SetRetentionInDays(val *float64)
	RetentionInDaysInput() *float64
	SasUrl() *string
	SetSasUrl(val *string)
	SasUrlInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference
type jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) InternalValue() *WindowsWebAppLogsApplicationLogsAzureBlobStorage {
	var returns *WindowsWebAppLogsApplicationLogsAzureBlobStorage
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) RetentionInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) RetentionInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retentionInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SasUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) SasUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sasUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference_Override(w WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetInternalValue(val *WindowsWebAppLogsApplicationLogsAzureBlobStorage) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetRetentionInDays(val *float64) {
	if err := j.validateSetRetentionInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retentionInDays",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetSasUrl(val *string) {
	if err := j.validateSetSasUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sasUrl",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppLogsApplicationLogsAzureBlobStorageOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

