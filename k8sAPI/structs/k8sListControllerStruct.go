package structs

import (
	apiv1 "k8s.io/api/core/v1"
)

type PodInfo struct {
	Namespace string
	Name      string
	PortInfo  []apiv1.ContainerPort
}

type ServiceInfo struct {
	Namespace string
	Name      string
	SpecInfo  apiv1.ServiceSpec
}

type DeployInfo struct {
	Namespace     string
	Name          string
	ContainerInfo []apiv1.Container
}

type AllInfo struct {
	PodInfoList     []AllinfoUnderPodInfo
	ServiceInfoList []AllinfoUnderServiceInfo
	DeployInfoList  []AllinfoUnderDeployInfo
}

type AllinfoUnderPodInfo struct {
	Namespace string
	Name      string
}

type AllinfoUnderServiceInfo struct {
	Namespace string
	Name      string
}

type AllinfoUnderDeployInfo struct {
	Namespace string
	Name      string
}
