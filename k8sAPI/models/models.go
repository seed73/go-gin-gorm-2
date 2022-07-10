// models/book.go

package models

import (
	"time"
)

type Task struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	AssingedTo string    `json:"assignedTo"`
	Task       string    `json:"task"`
	Expired    time.Time `json:"expired"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Pod_List struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	UserName   string    `json:"user_name"`
	UserPass   []byte    `json:"user_pass"`
	Label      string    `json:"label"`
	PodName    string    `json:"pod_name"`
	PortNum    string    `json:"port_num"`
	ImageName1 string    `json:"image_name1"`
	ImageName2 string    `json:"image_name2"`
	ImageName3 string    `json:"image_name3"`
	ImageName4 string    `json:"image_name4"`
	Option     string    `json:"option"`
	OptionVer  string    `json:"optionVer"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Pod_Detail struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	PodName     string    `json:"pod_name"`
	Port        int       `json:"port"`
	PortName    string    `json:"port_name"`
	PodProtocol string    `json:"pod_protocol"`
	Namespace   string    `json:"namespace"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Pod_Option_List struct {
	ID            uint      `json:"id" gorm:"primary_key"`
	OptionName    string    `json:"option_name"`
	OptionKorName string    `json:"option_kor_name"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Pod_Option_Detail struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	OptionName string    `json:"option_name"`
	OptionVer  string    `json:"option_ver"`
	OptionId   string    `json:"option_id"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Service_List struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	ServiceName string    `json:"service_name"`
	PodName     string    `json:"pod_name"`
	PortNum     string    `json:"port_num"`
	ServiceType string    `json:"service_type"`
	Namespace   string    `json:"namespace"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Service_Detail struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	ServiceName string    `json:"service_name"`
	PortName    string    `json:"port_name"`
	Port        int       `json:"port"`
	Namespace   string    `json:"namespace"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type K8s_Image_List struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	ImageName string    `json:"image_name"`
	Size      string    `json:"size"`
	OptionId  string    `json:"option_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
