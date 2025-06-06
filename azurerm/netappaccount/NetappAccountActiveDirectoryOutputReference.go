// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappaccount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/netappaccount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappAccountActiveDirectoryOutputReference interface {
	cdktf.ComplexObject
	AesEncryptionEnabled() interface{}
	SetAesEncryptionEnabled(val interface{})
	AesEncryptionEnabledInput() interface{}
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
	DnsServers() *[]*string
	SetDnsServers(val *[]*string)
	DnsServersInput() *[]*string
	Domain() *string
	SetDomain(val *string)
	DomainInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetappAccountActiveDirectory
	SetInternalValue(val *NetappAccountActiveDirectory)
	KerberosAdName() *string
	SetKerberosAdName(val *string)
	KerberosAdNameInput() *string
	KerberosKdcIp() *string
	SetKerberosKdcIp(val *string)
	KerberosKdcIpInput() *string
	LdapOverTlsEnabled() interface{}
	SetLdapOverTlsEnabled(val interface{})
	LdapOverTlsEnabledInput() interface{}
	LdapSigningEnabled() interface{}
	SetLdapSigningEnabled(val interface{})
	LdapSigningEnabledInput() interface{}
	LocalNfsUsersWithLdapAllowed() interface{}
	SetLocalNfsUsersWithLdapAllowed(val interface{})
	LocalNfsUsersWithLdapAllowedInput() interface{}
	OrganizationalUnit() *string
	SetOrganizationalUnit(val *string)
	OrganizationalUnitInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ServerRootCaCertificate() *string
	SetServerRootCaCertificate(val *string)
	ServerRootCaCertificateInput() *string
	SiteName() *string
	SetSiteName(val *string)
	SiteNameInput() *string
	SmbServerName() *string
	SetSmbServerName(val *string)
	SmbServerNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
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
	ResetAesEncryptionEnabled()
	ResetKerberosAdName()
	ResetKerberosKdcIp()
	ResetLdapOverTlsEnabled()
	ResetLdapSigningEnabled()
	ResetLocalNfsUsersWithLdapAllowed()
	ResetOrganizationalUnit()
	ResetServerRootCaCertificate()
	ResetSiteName()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappAccountActiveDirectoryOutputReference
type jsiiProxy_NetappAccountActiveDirectoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) AesEncryptionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aesEncryptionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) AesEncryptionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aesEncryptionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) DnsServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) DnsServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dnsServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) Domain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) DomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"domainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) InternalValue() *NetappAccountActiveDirectory {
	var returns *NetappAccountActiveDirectory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) KerberosAdName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosAdName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) KerberosAdNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosAdNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) KerberosKdcIp() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKdcIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) KerberosKdcIpInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kerberosKdcIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) LdapOverTlsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapOverTlsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) LdapOverTlsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapOverTlsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) LdapSigningEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapSigningEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) LdapSigningEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ldapSigningEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) LocalNfsUsersWithLdapAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localNfsUsersWithLdapAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) LocalNfsUsersWithLdapAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localNfsUsersWithLdapAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) OrganizationalUnit() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) OrganizationalUnitInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationalUnitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ServerRootCaCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverRootCaCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ServerRootCaCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverRootCaCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) SiteName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) SiteNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"siteNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) SmbServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) SmbServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"smbServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


func NewNetappAccountActiveDirectoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappAccountActiveDirectoryOutputReference {
	_init_.Initialize()

	if err := validateNewNetappAccountActiveDirectoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappAccountActiveDirectoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappAccount.NetappAccountActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappAccountActiveDirectoryOutputReference_Override(n NetappAccountActiveDirectoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappAccount.NetappAccountActiveDirectoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetAesEncryptionEnabled(val interface{}) {
	if err := j.validateSetAesEncryptionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"aesEncryptionEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetDnsServers(val *[]*string) {
	if err := j.validateSetDnsServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsServers",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetDomain(val *string) {
	if err := j.validateSetDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"domain",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetInternalValue(val *NetappAccountActiveDirectory) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetKerberosAdName(val *string) {
	if err := j.validateSetKerberosAdNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosAdName",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetKerberosKdcIp(val *string) {
	if err := j.validateSetKerberosKdcIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kerberosKdcIp",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetLdapOverTlsEnabled(val interface{}) {
	if err := j.validateSetLdapOverTlsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapOverTlsEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetLdapSigningEnabled(val interface{}) {
	if err := j.validateSetLdapSigningEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ldapSigningEnabled",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetLocalNfsUsersWithLdapAllowed(val interface{}) {
	if err := j.validateSetLocalNfsUsersWithLdapAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localNfsUsersWithLdapAllowed",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetOrganizationalUnit(val *string) {
	if err := j.validateSetOrganizationalUnitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"organizationalUnit",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetServerRootCaCertificate(val *string) {
	if err := j.validateSetServerRootCaCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverRootCaCertificate",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetSiteName(val *string) {
	if err := j.validateSetSiteNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"siteName",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetSmbServerName(val *string) {
	if err := j.validateSetSmbServerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"smbServerName",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetappAccountActiveDirectoryOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetAesEncryptionEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetAesEncryptionEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetKerberosAdName() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberosAdName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetKerberosKdcIp() {
	_jsii_.InvokeVoid(
		n,
		"resetKerberosKdcIp",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetLdapOverTlsEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetLdapOverTlsEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetLdapSigningEnabled() {
	_jsii_.InvokeVoid(
		n,
		"resetLdapSigningEnabled",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetLocalNfsUsersWithLdapAllowed() {
	_jsii_.InvokeVoid(
		n,
		"resetLocalNfsUsersWithLdapAllowed",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetOrganizationalUnit() {
	_jsii_.InvokeVoid(
		n,
		"resetOrganizationalUnit",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetServerRootCaCertificate() {
	_jsii_.InvokeVoid(
		n,
		"resetServerRootCaCertificate",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ResetSiteName() {
	_jsii_.InvokeVoid(
		n,
		"resetSiteName",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappAccountActiveDirectoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

