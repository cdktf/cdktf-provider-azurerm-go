package hdinsighthbasecluster


type HdinsightHbaseClusterNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#connection_direction HdinsightHbaseCluster#connection_direction}.
	ConnectionDirection *string `field:"optional" json:"connectionDirection" yaml:"connectionDirection"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/hdinsight_hbase_cluster#private_link_enabled HdinsightHbaseCluster#private_link_enabled}.
	PrivateLinkEnabled interface{} `field:"optional" json:"privateLinkEnabled" yaml:"privateLinkEnabled"`
}
