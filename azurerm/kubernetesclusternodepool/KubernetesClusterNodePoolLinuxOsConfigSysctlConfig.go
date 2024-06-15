// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesclusternodepool


type KubernetesClusterNodePoolLinuxOsConfigSysctlConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#fs_aio_max_nr KubernetesClusterNodePool#fs_aio_max_nr}.
	FsAioMaxNr *float64 `field:"optional" json:"fsAioMaxNr" yaml:"fsAioMaxNr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#fs_file_max KubernetesClusterNodePool#fs_file_max}.
	FsFileMax *float64 `field:"optional" json:"fsFileMax" yaml:"fsFileMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#fs_inotify_max_user_watches KubernetesClusterNodePool#fs_inotify_max_user_watches}.
	FsInotifyMaxUserWatches *float64 `field:"optional" json:"fsInotifyMaxUserWatches" yaml:"fsInotifyMaxUserWatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#fs_nr_open KubernetesClusterNodePool#fs_nr_open}.
	FsNrOpen *float64 `field:"optional" json:"fsNrOpen" yaml:"fsNrOpen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#kernel_threads_max KubernetesClusterNodePool#kernel_threads_max}.
	KernelThreadsMax *float64 `field:"optional" json:"kernelThreadsMax" yaml:"kernelThreadsMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_netdev_max_backlog KubernetesClusterNodePool#net_core_netdev_max_backlog}.
	NetCoreNetdevMaxBacklog *float64 `field:"optional" json:"netCoreNetdevMaxBacklog" yaml:"netCoreNetdevMaxBacklog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_optmem_max KubernetesClusterNodePool#net_core_optmem_max}.
	NetCoreOptmemMax *float64 `field:"optional" json:"netCoreOptmemMax" yaml:"netCoreOptmemMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_rmem_default KubernetesClusterNodePool#net_core_rmem_default}.
	NetCoreRmemDefault *float64 `field:"optional" json:"netCoreRmemDefault" yaml:"netCoreRmemDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_rmem_max KubernetesClusterNodePool#net_core_rmem_max}.
	NetCoreRmemMax *float64 `field:"optional" json:"netCoreRmemMax" yaml:"netCoreRmemMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_somaxconn KubernetesClusterNodePool#net_core_somaxconn}.
	NetCoreSomaxconn *float64 `field:"optional" json:"netCoreSomaxconn" yaml:"netCoreSomaxconn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_wmem_default KubernetesClusterNodePool#net_core_wmem_default}.
	NetCoreWmemDefault *float64 `field:"optional" json:"netCoreWmemDefault" yaml:"netCoreWmemDefault"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_core_wmem_max KubernetesClusterNodePool#net_core_wmem_max}.
	NetCoreWmemMax *float64 `field:"optional" json:"netCoreWmemMax" yaml:"netCoreWmemMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_ip_local_port_range_max KubernetesClusterNodePool#net_ipv4_ip_local_port_range_max}.
	NetIpv4IpLocalPortRangeMax *float64 `field:"optional" json:"netIpv4IpLocalPortRangeMax" yaml:"netIpv4IpLocalPortRangeMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_ip_local_port_range_min KubernetesClusterNodePool#net_ipv4_ip_local_port_range_min}.
	NetIpv4IpLocalPortRangeMin *float64 `field:"optional" json:"netIpv4IpLocalPortRangeMin" yaml:"netIpv4IpLocalPortRangeMin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_neigh_default_gc_thresh1 KubernetesClusterNodePool#net_ipv4_neigh_default_gc_thresh1}.
	NetIpv4NeighDefaultGcThresh1 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh1" yaml:"netIpv4NeighDefaultGcThresh1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_neigh_default_gc_thresh2 KubernetesClusterNodePool#net_ipv4_neigh_default_gc_thresh2}.
	NetIpv4NeighDefaultGcThresh2 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh2" yaml:"netIpv4NeighDefaultGcThresh2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_neigh_default_gc_thresh3 KubernetesClusterNodePool#net_ipv4_neigh_default_gc_thresh3}.
	NetIpv4NeighDefaultGcThresh3 *float64 `field:"optional" json:"netIpv4NeighDefaultGcThresh3" yaml:"netIpv4NeighDefaultGcThresh3"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_fin_timeout KubernetesClusterNodePool#net_ipv4_tcp_fin_timeout}.
	NetIpv4TcpFinTimeout *float64 `field:"optional" json:"netIpv4TcpFinTimeout" yaml:"netIpv4TcpFinTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_keepalive_intvl KubernetesClusterNodePool#net_ipv4_tcp_keepalive_intvl}.
	NetIpv4TcpKeepaliveIntvl *float64 `field:"optional" json:"netIpv4TcpKeepaliveIntvl" yaml:"netIpv4TcpKeepaliveIntvl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_keepalive_probes KubernetesClusterNodePool#net_ipv4_tcp_keepalive_probes}.
	NetIpv4TcpKeepaliveProbes *float64 `field:"optional" json:"netIpv4TcpKeepaliveProbes" yaml:"netIpv4TcpKeepaliveProbes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_keepalive_time KubernetesClusterNodePool#net_ipv4_tcp_keepalive_time}.
	NetIpv4TcpKeepaliveTime *float64 `field:"optional" json:"netIpv4TcpKeepaliveTime" yaml:"netIpv4TcpKeepaliveTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_max_syn_backlog KubernetesClusterNodePool#net_ipv4_tcp_max_syn_backlog}.
	NetIpv4TcpMaxSynBacklog *float64 `field:"optional" json:"netIpv4TcpMaxSynBacklog" yaml:"netIpv4TcpMaxSynBacklog"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_max_tw_buckets KubernetesClusterNodePool#net_ipv4_tcp_max_tw_buckets}.
	NetIpv4TcpMaxTwBuckets *float64 `field:"optional" json:"netIpv4TcpMaxTwBuckets" yaml:"netIpv4TcpMaxTwBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_ipv4_tcp_tw_reuse KubernetesClusterNodePool#net_ipv4_tcp_tw_reuse}.
	NetIpv4TcpTwReuse interface{} `field:"optional" json:"netIpv4TcpTwReuse" yaml:"netIpv4TcpTwReuse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_netfilter_nf_conntrack_buckets KubernetesClusterNodePool#net_netfilter_nf_conntrack_buckets}.
	NetNetfilterNfConntrackBuckets *float64 `field:"optional" json:"netNetfilterNfConntrackBuckets" yaml:"netNetfilterNfConntrackBuckets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#net_netfilter_nf_conntrack_max KubernetesClusterNodePool#net_netfilter_nf_conntrack_max}.
	NetNetfilterNfConntrackMax *float64 `field:"optional" json:"netNetfilterNfConntrackMax" yaml:"netNetfilterNfConntrackMax"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#vm_max_map_count KubernetesClusterNodePool#vm_max_map_count}.
	VmMaxMapCount *float64 `field:"optional" json:"vmMaxMapCount" yaml:"vmMaxMapCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#vm_swappiness KubernetesClusterNodePool#vm_swappiness}.
	VmSwappiness *float64 `field:"optional" json:"vmSwappiness" yaml:"vmSwappiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.108.0/docs/resources/kubernetes_cluster_node_pool#vm_vfs_cache_pressure KubernetesClusterNodePool#vm_vfs_cache_pressure}.
	VmVfsCachePressure *float64 `field:"optional" json:"vmVfsCachePressure" yaml:"vmVfsCachePressure"`
}

