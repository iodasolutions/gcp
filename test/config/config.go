package config

var ProjectId = "ageless-domain-414706"
var Zone = "europe-west9-b"
var InstanceType = "e2-medium"
var InstanceDiskSize int64 = 20
var ImageProject = "ubuntu-os-cloud"
var ImageFamily = "ubuntu-2004-lts"
var VolumeSize int64 = 100
var VolumeType = "pd-standard" //pd-standard,pd-ssd,pd-balanced
var Network = "default"
var User = "ubuntu"

func String(v string) *string {
	return &v
}
