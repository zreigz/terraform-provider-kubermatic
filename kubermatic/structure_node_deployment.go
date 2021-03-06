package kubermatic

import (
	"github.com/kubermatic/go-kubermatic/models"
)

// flatteners
func flattenNodeDeploymentSpec(in *models.NodeDeploymentSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	if in.Replicas != nil {
		att["replicas"] = *in.Replicas
	}

	if in.Template != nil {
		att["template"] = flattenNodeSpec(in.Template)
	}

	return []interface{}{att}
}

func flattenNodeSpec(in *models.NodeSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	if l := len(in.Labels); l > 0 {
		labels := make(map[string]string, l)
		for key, val := range in.Labels {
			labels[key] = val
		}
		att["labels"] = labels
	}

	if in.OperatingSystem != nil {
		att["operating_system"] = flattenOperatingSystem(in.OperatingSystem)
	}

	if in.Versions != nil {
		att["versions"] = flattenNodeVersion(in.Versions)
	}

	if l := len(in.Taints); l > 0 {
		t := make([]interface{}, l)
		for i, v := range in.Taints {
			t[i] = flattenTaintSpec(v)
		}
		att["taints"] = t
	}

	if in.Cloud != nil {
		att["cloud"] = flattenNodeCloudSpec(in.Cloud)
	}

	return []interface{}{att}
}

func flattenOperatingSystem(in *models.OperatingSystemSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	if in.Ubuntu != nil {
		att["ubuntu"] = flattenUbuntu(in.Ubuntu)
	}

	if in.Centos != nil {
		att["centos"] = flattenCentos(in.Centos)
	}

	if in.ContainerLinux != nil {
		att["container_linux"] = flattenContainerLinux(in.ContainerLinux)
	}

	return []interface{}{att}
}

func flattenUbuntu(in *models.UbuntuSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	att["dist_upgrade_on_boot"] = in.DistUpgradeOnBoot

	return []interface{}{att}
}

func flattenCentos(in *models.CentOSSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	att["dist_upgrade_on_boot"] = in.DistUpgradeOnBoot

	return []interface{}{att}
}

func flattenContainerLinux(in *models.ContainerLinuxSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	att["disable_auto_update"] = in.DisableAutoUpdate

	return []interface{}{att}
}

func flattenNodeVersion(in *models.NodeVersionInfo) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	if in.Kubelet != "" {
		att["kubelet"] = in.Kubelet
	}

	return []interface{}{att}
}

func flattenTaintSpec(in *models.TaintSpec) map[string]interface{} {
	if in == nil {
		return map[string]interface{}{}
	}

	att := make(map[string]interface{})

	if in.Key != "" {
		att["key"] = in.Key
	}

	if in.Value != "" {
		att["value"] = in.Value
	}

	if in.Effect != "" {
		att["effect"] = in.Effect
	}

	return att
}

func flattenNodeCloudSpec(in *models.NodeCloudSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	if in.Aws != nil {
		att["aws"] = flattenAWSNodeSpec(in.Aws)
	}

	// TODO: add all cloud providers

	return []interface{}{att}
}

func flattenAWSNodeSpec(in *models.AWSNodeSpec) []interface{} {
	if in == nil {
		return []interface{}{}
	}

	att := make(map[string]interface{})

	att["assign_public_ip"] = in.AssignPublicIP

	if l := len(in.Tags); l > 0 {
		t := make(map[string]string, l)
		for key, val := range in.Tags {
			t[key] = val
		}
		att["tags"] = t
	}

	if in.AMI != "" {
		att["ami"] = in.AMI
	}

	if in.AvailabilityZone != "" {
		att["availability_zone"] = in.AvailabilityZone
	}

	if in.SubnetID != "" {
		att["subnet_id"] = in.SubnetID
	}

	if in.VolumeType != nil {
		att["volume_type"] = *in.VolumeType
	}

	if in.VolumeSize != nil {
		att["disk_size"] = *in.VolumeSize
	}

	if in.InstanceType != nil {
		att["instance_type"] = *in.InstanceType
	}

	return []interface{}{att}
}

// expanders

func expandNodeDeploymentSpec(p []interface{}) *models.NodeDeploymentSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.NodeDeploymentSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["replicas"]; ok {
		obj.Replicas = int32ToPtr(int32(v.(int)))
	}

	if v, ok := in["template"]; ok {
		obj.Template = expandNodeSpec(v.([]interface{}))
	}

	return obj
}

func expandNodeSpec(p []interface{}) *models.NodeSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.NodeSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["labels"]; ok {
		obj.Labels = make(map[string]string)
		for key, val := range v.(map[string]interface{}) {
			obj.Labels[key] = val.(string)
		}
	}

	if v, ok := in["operating_system"]; ok {
		obj.OperatingSystem = expandOperatingSystem(v.([]interface{}))
	}

	if v, ok := in["versions"]; ok {
		obj.Versions = expandNodeVersion(v.([]interface{}))
	}

	if v, ok := in["taints"]; ok {
		for _, t := range v.([]interface{}) {
			obj.Taints = append(obj.Taints, expandTaintSpec(t.(map[string]interface{})))
		}
	}

	if v, ok := in["cloud"]; ok {
		obj.Cloud = expandNodeCloudSpec(v.([]interface{}))
	}

	return obj
}

func expandOperatingSystem(p []interface{}) *models.OperatingSystemSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.OperatingSystemSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["ubuntu"]; ok {
		obj.Ubuntu = expandUbuntu(v.([]interface{}))
	}

	if v, ok := in["centos"]; ok {
		obj.Centos = expandCentos(v.([]interface{}))

	}

	if v, ok := in["container_linux"]; ok {
		obj.ContainerLinux = expandContainerLinux(v.([]interface{}))
	}

	return obj
}

func expandUbuntu(p []interface{}) *models.UbuntuSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.UbuntuSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["dist_upgrade_on_boot"]; ok {
		obj.DistUpgradeOnBoot = v.(bool)
	}

	return obj
}

func expandCentos(p []interface{}) *models.CentOSSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.CentOSSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["dist_upgrade_on_boot"]; ok {
		obj.DistUpgradeOnBoot = v.(bool)
	}

	return obj
}

func expandContainerLinux(p []interface{}) *models.ContainerLinuxSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.ContainerLinuxSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["disable_auto_update"]; ok {
		obj.DisableAutoUpdate = v.(bool)
	}

	return obj
}

func expandNodeVersion(p []interface{}) *models.NodeVersionInfo {
	if len(p) < 1 {
		return nil
	}
	obj := &models.NodeVersionInfo{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["kubelet"]; ok {
		obj.Kubelet = v.(string)
	}

	return obj
}

func expandTaintSpec(in map[string]interface{}) *models.TaintSpec {
	obj := &models.TaintSpec{}

	if v, ok := in["key"]; ok {
		obj.Key = v.(string)
	}

	if v, ok := in["value"]; ok {
		obj.Value = v.(string)
	}

	if v, ok := in["effect"]; ok {
		obj.Effect = v.(string)
	}

	return obj
}

func expandNodeCloudSpec(p []interface{}) *models.NodeCloudSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.NodeCloudSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["aws"]; ok {
		obj.Aws = expandAWSNodeSpec(v.([]interface{}))
	}

	return obj
}

func expandAWSNodeSpec(p []interface{}) *models.AWSNodeSpec {
	if len(p) < 1 {
		return nil
	}
	obj := &models.AWSNodeSpec{}
	if p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["instance_type"]; ok {
		obj.InstanceType = strToPtr(v.(string))
	}

	if v, ok := in["disk_size"]; ok {
		obj.VolumeSize = int64ToPtr(v.(int))
	}

	if v, ok := in["volume_type"]; ok {
		obj.VolumeType = strToPtr(v.(string))
	}

	if v, ok := in["availability_zone"]; ok {
		obj.AvailabilityZone = v.(string)
	}

	if v, ok := in["subnet_id"]; ok {
		obj.SubnetID = v.(string)
	}

	if v, ok := in["assign_public_ip"]; ok {
		obj.AssignPublicIP = v.(bool)
	}

	if v, ok := in["ami"]; ok {
		obj.AMI = v.(string)
	}

	if v, ok := in["tags"]; ok {
		obj.Tags = make(map[string]string)
		for key, val := range v.(map[string]interface{}) {
			obj.Tags[key] = val.(string)
		}
	}

	return obj
}
