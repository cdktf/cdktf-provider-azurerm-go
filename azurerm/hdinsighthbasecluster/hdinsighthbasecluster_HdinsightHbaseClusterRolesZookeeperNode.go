package hdinsighthbasecluster


type HdinsightHbaseClusterRolesZookeeperNode struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#username HdinsightHbaseCluster#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#vm_size HdinsightHbaseCluster#vm_size}.
	VmSize *string `field:"required" json:"vmSize" yaml:"vmSize"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#password HdinsightHbaseCluster#password}.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#ssh_keys HdinsightHbaseCluster#ssh_keys}.
	SshKeys *[]*string `field:"optional" json:"sshKeys" yaml:"sshKeys"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#subnet_id HdinsightHbaseCluster#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#virtual_network_id HdinsightHbaseCluster#virtual_network_id}.
	VirtualNetworkId *string `field:"optional" json:"virtualNetworkId" yaml:"virtualNetworkId"`
}
