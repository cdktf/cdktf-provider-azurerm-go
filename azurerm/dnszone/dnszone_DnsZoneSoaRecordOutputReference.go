package dnszone

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/dnszone/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DnsZoneSoaRecordOutputReference interface {
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
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	ExpireTime() *float64
	SetExpireTime(val *float64)
	ExpireTimeInput() *float64
	Fqdn() *string
	// Experimental.
	Fqn() *string
	HostName() *string
	SetHostName(val *string)
	HostNameInput() *string
	InternalValue() *DnsZoneSoaRecord
	SetInternalValue(val *DnsZoneSoaRecord)
	MinimumTtl() *float64
	SetMinimumTtl(val *float64)
	MinimumTtlInput() *float64
	RefreshTime() *float64
	SetRefreshTime(val *float64)
	RefreshTimeInput() *float64
	RetryTime() *float64
	SetRetryTime(val *float64)
	RetryTimeInput() *float64
	SerialNumber() *float64
	SetSerialNumber(val *float64)
	SerialNumberInput() *float64
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
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
	ResetExpireTime()
	ResetMinimumTtl()
	ResetRefreshTime()
	ResetRetryTime()
	ResetSerialNumber()
	ResetTags()
	ResetTtl()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DnsZoneSoaRecordOutputReference
type jsiiProxy_DnsZoneSoaRecordOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) ExpireTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) ExpireTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"expireTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) Fqdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) HostName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) HostNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) InternalValue() *DnsZoneSoaRecord {
	var returns *DnsZoneSoaRecord
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) MinimumTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) MinimumTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minimumTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) RefreshTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) RefreshTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"refreshTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) RetryTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) RetryTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) SerialNumber() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) SerialNumberInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"serialNumberInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


func NewDnsZoneSoaRecordOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DnsZoneSoaRecordOutputReference {
	_init_.Initialize()

	if err := validateNewDnsZoneSoaRecordOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DnsZoneSoaRecordOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dnsZone.DnsZoneSoaRecordOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDnsZoneSoaRecordOutputReference_Override(d DnsZoneSoaRecordOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dnsZone.DnsZoneSoaRecordOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetExpireTime(val *float64) {
	if err := j.validateSetExpireTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expireTime",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetHostName(val *string) {
	if err := j.validateSetHostNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostName",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetInternalValue(val *DnsZoneSoaRecord) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetMinimumTtl(val *float64) {
	if err := j.validateSetMinimumTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minimumTtl",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetRefreshTime(val *float64) {
	if err := j.validateSetRefreshTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"refreshTime",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetRetryTime(val *float64) {
	if err := j.validateSetRetryTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryTime",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetSerialNumber(val *float64) {
	if err := j.validateSetSerialNumberParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serialNumber",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DnsZoneSoaRecordOutputReference)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetExpireTime() {
	_jsii_.InvokeVoid(
		d,
		"resetExpireTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetMinimumTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetMinimumTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetRefreshTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRefreshTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetRetryTime() {
	_jsii_.InvokeVoid(
		d,
		"resetRetryTime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetSerialNumber() {
	_jsii_.InvokeVoid(
		d,
		"resetSerialNumber",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		d,
		"resetTags",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ResetTtl() {
	_jsii_.InvokeVoid(
		d,
		"resetTtl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DnsZoneSoaRecordOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

