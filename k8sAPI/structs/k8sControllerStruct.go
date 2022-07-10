package structs

type ImageNameInput struct {
	ImageName []string `json:"assignedTo" example:"20210101"`
}

type CreatePodServiceParam struct {
	PortNum string `json:"portNum" example:"2"`
	Option  string `json:"option" example:"go"`
	Ver     string `json:"Ver" example:"1.15.15"`
	Port    []int  `json:"port" example:"8080,7070"`
	Label   string `json:"label" example:"생성할 VSCODE명칭"`
}

type DeletePodServiceParam struct {
	PodName string `json:"podName" example:"shiftone-go-1.15.15-20220218"`
}

type PodPortInfoRes struct {
	Port int
}

type ServiceStartParam struct {
	PodName string `json:"podName" example:"shiftone-go-1.15.15-20220218"`
}
