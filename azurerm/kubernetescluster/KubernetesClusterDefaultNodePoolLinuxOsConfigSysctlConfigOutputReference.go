// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference interface {
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
	FsAioMaxNr() *float64
	SetFsAioMaxNr(val *float64)
	FsAioMaxNrInput() *float64
	FsFileMax() *float64
	SetFsFileMax(val *float64)
	FsFileMaxInput() *float64
	FsInotifyMaxUserWatches() *float64
	SetFsInotifyMaxUserWatches(val *float64)
	FsInotifyMaxUserWatchesInput() *float64
	FsNrOpen() *float64
	SetFsNrOpen(val *float64)
	FsNrOpenInput() *float64
	InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig)
	KernelThreadsMax() *float64
	SetKernelThreadsMax(val *float64)
	KernelThreadsMaxInput() *float64
	NetCoreNetdevMaxBacklog() *float64
	SetNetCoreNetdevMaxBacklog(val *float64)
	NetCoreNetdevMaxBacklogInput() *float64
	NetCoreOptmemMax() *float64
	SetNetCoreOptmemMax(val *float64)
	NetCoreOptmemMaxInput() *float64
	NetCoreRmemDefault() *float64
	SetNetCoreRmemDefault(val *float64)
	NetCoreRmemDefaultInput() *float64
	NetCoreRmemMax() *float64
	SetNetCoreRmemMax(val *float64)
	NetCoreRmemMaxInput() *float64
	NetCoreSomaxconn() *float64
	SetNetCoreSomaxconn(val *float64)
	NetCoreSomaxconnInput() *float64
	NetCoreWmemDefault() *float64
	SetNetCoreWmemDefault(val *float64)
	NetCoreWmemDefaultInput() *float64
	NetCoreWmemMax() *float64
	SetNetCoreWmemMax(val *float64)
	NetCoreWmemMaxInput() *float64
	NetIpv4IpLocalPortRangeMax() *float64
	SetNetIpv4IpLocalPortRangeMax(val *float64)
	NetIpv4IpLocalPortRangeMaxInput() *float64
	NetIpv4IpLocalPortRangeMin() *float64
	SetNetIpv4IpLocalPortRangeMin(val *float64)
	NetIpv4IpLocalPortRangeMinInput() *float64
	NetIpv4NeighDefaultGcThresh1() *float64
	SetNetIpv4NeighDefaultGcThresh1(val *float64)
	NetIpv4NeighDefaultGcThresh1Input() *float64
	NetIpv4NeighDefaultGcThresh2() *float64
	SetNetIpv4NeighDefaultGcThresh2(val *float64)
	NetIpv4NeighDefaultGcThresh2Input() *float64
	NetIpv4NeighDefaultGcThresh3() *float64
	SetNetIpv4NeighDefaultGcThresh3(val *float64)
	NetIpv4NeighDefaultGcThresh3Input() *float64
	NetIpv4TcpFinTimeout() *float64
	SetNetIpv4TcpFinTimeout(val *float64)
	NetIpv4TcpFinTimeoutInput() *float64
	NetIpv4TcpKeepaliveIntvl() *float64
	SetNetIpv4TcpKeepaliveIntvl(val *float64)
	NetIpv4TcpKeepaliveIntvlInput() *float64
	NetIpv4TcpKeepaliveProbes() *float64
	SetNetIpv4TcpKeepaliveProbes(val *float64)
	NetIpv4TcpKeepaliveProbesInput() *float64
	NetIpv4TcpKeepaliveTime() *float64
	SetNetIpv4TcpKeepaliveTime(val *float64)
	NetIpv4TcpKeepaliveTimeInput() *float64
	NetIpv4TcpMaxSynBacklog() *float64
	SetNetIpv4TcpMaxSynBacklog(val *float64)
	NetIpv4TcpMaxSynBacklogInput() *float64
	NetIpv4TcpMaxTwBuckets() *float64
	SetNetIpv4TcpMaxTwBuckets(val *float64)
	NetIpv4TcpMaxTwBucketsInput() *float64
	NetIpv4TcpTwReuse() interface{}
	SetNetIpv4TcpTwReuse(val interface{})
	NetIpv4TcpTwReuseInput() interface{}
	NetNetfilterNfConntrackBuckets() *float64
	SetNetNetfilterNfConntrackBuckets(val *float64)
	NetNetfilterNfConntrackBucketsInput() *float64
	NetNetfilterNfConntrackMax() *float64
	SetNetNetfilterNfConntrackMax(val *float64)
	NetNetfilterNfConntrackMaxInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmMaxMapCount() *float64
	SetVmMaxMapCount(val *float64)
	VmMaxMapCountInput() *float64
	VmSwappiness() *float64
	SetVmSwappiness(val *float64)
	VmSwappinessInput() *float64
	VmVfsCachePressure() *float64
	SetVmVfsCachePressure(val *float64)
	VmVfsCachePressureInput() *float64
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
	ResetFsAioMaxNr()
	ResetFsFileMax()
	ResetFsInotifyMaxUserWatches()
	ResetFsNrOpen()
	ResetKernelThreadsMax()
	ResetNetCoreNetdevMaxBacklog()
	ResetNetCoreOptmemMax()
	ResetNetCoreRmemDefault()
	ResetNetCoreRmemMax()
	ResetNetCoreSomaxconn()
	ResetNetCoreWmemDefault()
	ResetNetCoreWmemMax()
	ResetNetIpv4IpLocalPortRangeMax()
	ResetNetIpv4IpLocalPortRangeMin()
	ResetNetIpv4NeighDefaultGcThresh1()
	ResetNetIpv4NeighDefaultGcThresh2()
	ResetNetIpv4NeighDefaultGcThresh3()
	ResetNetIpv4TcpFinTimeout()
	ResetNetIpv4TcpKeepaliveIntvl()
	ResetNetIpv4TcpKeepaliveProbes()
	ResetNetIpv4TcpKeepaliveTime()
	ResetNetIpv4TcpMaxSynBacklog()
	ResetNetIpv4TcpMaxTwBuckets()
	ResetNetIpv4TcpTwReuse()
	ResetNetNetfilterNfConntrackBuckets()
	ResetNetNetfilterNfConntrackMax()
	ResetVmMaxMapCount()
	ResetVmSwappiness()
	ResetVmVfsCachePressure()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsAioMaxNr() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsAioMaxNr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsAioMaxNrInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsAioMaxNrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsFileMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsFileMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsFileMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsFileMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsInotifyMaxUserWatches() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsInotifyMaxUserWatches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsInotifyMaxUserWatchesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsInotifyMaxUserWatchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsNrOpen() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsNrOpen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) FsNrOpenInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fsNrOpenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig {
	var returns *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) KernelThreadsMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kernelThreadsMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) KernelThreadsMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"kernelThreadsMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreNetdevMaxBacklog() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreNetdevMaxBacklog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreNetdevMaxBacklogInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreNetdevMaxBacklogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreOptmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreOptmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreOptmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreOptmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreRmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreRmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreSomaxconn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreSomaxconn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreSomaxconnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreSomaxconnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemDefault() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemDefault",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemDefaultInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemDefaultInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetCoreWmemMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netCoreWmemMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4IpLocalPortRangeMinInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4IpLocalPortRangeMinInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh1() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh1",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh1Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh1Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh2() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh2",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh2Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh2Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh3() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh3",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4NeighDefaultGcThresh3Input() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4NeighDefaultGcThresh3Input",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpFinTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpFinTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpFinTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpFinTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveIntvl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveIntvl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveIntvlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveIntvlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveProbes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveProbes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveProbesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveProbesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpKeepaliveTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpKeepaliveTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxSynBacklog() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxSynBacklog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxSynBacklogInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxSynBacklogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxTwBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxTwBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpMaxTwBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netIpv4TcpMaxTwBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpTwReuse() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netIpv4TcpTwReuse",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetIpv4TcpTwReuseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"netIpv4TcpTwReuseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackBuckets() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackBucketsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackMax() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) NetNetfilterNfConntrackMaxInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"netNetfilterNfConntrackMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmMaxMapCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmMaxMapCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmMaxMapCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmMaxMapCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmSwappiness() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmSwappiness",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmSwappinessInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmSwappinessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmVfsCachePressure() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmVfsCachePressure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) VmVfsCachePressureInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmVfsCachePressureInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference_Override(k KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetFsAioMaxNr(val *float64) {
	if err := j.validateSetFsAioMaxNrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsAioMaxNr",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetFsFileMax(val *float64) {
	if err := j.validateSetFsFileMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsFileMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetFsInotifyMaxUserWatches(val *float64) {
	if err := j.validateSetFsInotifyMaxUserWatchesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsInotifyMaxUserWatches",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetFsNrOpen(val *float64) {
	if err := j.validateSetFsNrOpenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fsNrOpen",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetInternalValue(val *KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetKernelThreadsMax(val *float64) {
	if err := j.validateSetKernelThreadsMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kernelThreadsMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreNetdevMaxBacklog(val *float64) {
	if err := j.validateSetNetCoreNetdevMaxBacklogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreNetdevMaxBacklog",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreOptmemMax(val *float64) {
	if err := j.validateSetNetCoreOptmemMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreOptmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreRmemDefault(val *float64) {
	if err := j.validateSetNetCoreRmemDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreRmemDefault",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreRmemMax(val *float64) {
	if err := j.validateSetNetCoreRmemMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreRmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreSomaxconn(val *float64) {
	if err := j.validateSetNetCoreSomaxconnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreSomaxconn",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreWmemDefault(val *float64) {
	if err := j.validateSetNetCoreWmemDefaultParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreWmemDefault",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetCoreWmemMax(val *float64) {
	if err := j.validateSetNetCoreWmemMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netCoreWmemMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4IpLocalPortRangeMax(val *float64) {
	if err := j.validateSetNetIpv4IpLocalPortRangeMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4IpLocalPortRangeMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4IpLocalPortRangeMin(val *float64) {
	if err := j.validateSetNetIpv4IpLocalPortRangeMinParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4IpLocalPortRangeMin",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4NeighDefaultGcThresh1(val *float64) {
	if err := j.validateSetNetIpv4NeighDefaultGcThresh1Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh1",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4NeighDefaultGcThresh2(val *float64) {
	if err := j.validateSetNetIpv4NeighDefaultGcThresh2Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh2",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4NeighDefaultGcThresh3(val *float64) {
	if err := j.validateSetNetIpv4NeighDefaultGcThresh3Parameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4NeighDefaultGcThresh3",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpFinTimeout(val *float64) {
	if err := j.validateSetNetIpv4TcpFinTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpFinTimeout",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpKeepaliveIntvl(val *float64) {
	if err := j.validateSetNetIpv4TcpKeepaliveIntvlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveIntvl",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpKeepaliveProbes(val *float64) {
	if err := j.validateSetNetIpv4TcpKeepaliveProbesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveProbes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpKeepaliveTime(val *float64) {
	if err := j.validateSetNetIpv4TcpKeepaliveTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpKeepaliveTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpMaxSynBacklog(val *float64) {
	if err := j.validateSetNetIpv4TcpMaxSynBacklogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpMaxSynBacklog",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpMaxTwBuckets(val *float64) {
	if err := j.validateSetNetIpv4TcpMaxTwBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpMaxTwBuckets",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetIpv4TcpTwReuse(val interface{}) {
	if err := j.validateSetNetIpv4TcpTwReuseParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netIpv4TcpTwReuse",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetNetfilterNfConntrackBuckets(val *float64) {
	if err := j.validateSetNetNetfilterNfConntrackBucketsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netNetfilterNfConntrackBuckets",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetNetNetfilterNfConntrackMax(val *float64) {
	if err := j.validateSetNetNetfilterNfConntrackMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netNetfilterNfConntrackMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetVmMaxMapCount(val *float64) {
	if err := j.validateSetVmMaxMapCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmMaxMapCount",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetVmSwappiness(val *float64) {
	if err := j.validateSetVmSwappinessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSwappiness",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference)SetVmVfsCachePressure(val *float64) {
	if err := j.validateSetVmVfsCachePressureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmVfsCachePressure",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsAioMaxNr() {
	_jsii_.InvokeVoid(
		k,
		"resetFsAioMaxNr",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsFileMax() {
	_jsii_.InvokeVoid(
		k,
		"resetFsFileMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsInotifyMaxUserWatches() {
	_jsii_.InvokeVoid(
		k,
		"resetFsInotifyMaxUserWatches",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetFsNrOpen() {
	_jsii_.InvokeVoid(
		k,
		"resetFsNrOpen",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetKernelThreadsMax() {
	_jsii_.InvokeVoid(
		k,
		"resetKernelThreadsMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreNetdevMaxBacklog() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreNetdevMaxBacklog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreOptmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreOptmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreRmemDefault() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreRmemDefault",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreRmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreRmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreSomaxconn() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreSomaxconn",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreWmemDefault() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreWmemDefault",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetCoreWmemMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetCoreWmemMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4IpLocalPortRangeMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4IpLocalPortRangeMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4IpLocalPortRangeMin() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4IpLocalPortRangeMin",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh1() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh1",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh2() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh2",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4NeighDefaultGcThresh3() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4NeighDefaultGcThresh3",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpFinTimeout() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpFinTimeout",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveIntvl() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveIntvl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveProbes() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveProbes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpKeepaliveTime() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpKeepaliveTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpMaxSynBacklog() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpMaxSynBacklog",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpMaxTwBuckets() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpMaxTwBuckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetIpv4TcpTwReuse() {
	_jsii_.InvokeVoid(
		k,
		"resetNetIpv4TcpTwReuse",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetNetfilterNfConntrackBuckets() {
	_jsii_.InvokeVoid(
		k,
		"resetNetNetfilterNfConntrackBuckets",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetNetNetfilterNfConntrackMax() {
	_jsii_.InvokeVoid(
		k,
		"resetNetNetfilterNfConntrackMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmMaxMapCount() {
	_jsii_.InvokeVoid(
		k,
		"resetVmMaxMapCount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmSwappiness() {
	_jsii_.InvokeVoid(
		k,
		"resetVmSwappiness",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ResetVmVfsCachePressure() {
	_jsii_.InvokeVoid(
		k,
		"resetVmVfsCachePressure",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolLinuxOsConfigSysctlConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

