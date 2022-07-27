# toDoGolang

# 실행법
 go run ./main.go


# table 생성


CREATE TABLE k8sapi.task (
	id INT auto_increment primary key NOT NULL,
	assignedTo VARCHAR(100) NULL,
	task VARCHAR(100) NULL,
	expired VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;



CREATE TABLE k8sapi.k8s_pod_lists (
	id INT auto_increment primary key NOT NULL,
	user_name VARCHAR(100) NULL,
	user_pass VARCHAR(100) NULL,
	pod_name VARCHAR(100) NULL,
	port_num VARCHAR(100) NULL,
	image_name1 VARCHAR(100) NULL,
	image_name2 VARCHAR(100) NULL,
	image_name3 VARCHAR(100) NULL,
	image_name4 VARCHAR(100) NULL,
	option VARCHAR(100) NULL,
	optionVer VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


CREATE TABLE k8sapi.k8s_pod_details (
	id INT auto_increment primary key NOT NULL,
	pod_name VARCHAR(100) NULL,
	port VARCHAR(100) NULL,
	port_name VARCHAR(100) NULL,
	pod_protocol VARCHAR(100) NULL,
	namespace VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


CREATE TABLE k8sapi.k8s_pod_option_lists (
	id INT auto_increment primary key NOT NULL,
	option_name VARCHAR(100) NULL,
	option_kor_name VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


CREATE TABLE k8sapi.k8s_pod_option_detail (
	id INT auto_increment primary key NOT NULL,
	option_name VARCHAR(100) NULL,
	option_ver VARCHAR(100) NULL,
	option_id VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


CREATE TABLE k8sapi.k8s_service_lists (
	id INT auto_increment primary key NOT NULL,
	service_name VARCHAR(100) NULL,
	pod_name VARCHAR(100) NULL,
	port_num VARCHAR(100) NULL,
	service_type VARCHAR(100) NULL,
	namespace VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


CREATE TABLE k8sapi.k8s_service_details (
	id INT auto_increment primary key NOT NULL,
	service_name VARCHAR(100) NULL,
	pod_name VARCHAR(100) NULL,
	port_num VARCHAR(100) NULL,
	namespace VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


CREATE TABLE k8sapi.k8s_image_lists (
	id INT auto_increment primary key NOT NULL,
	image_name VARCHAR(100) NULL,
	size VARCHAR(100) NULL,
	option_id VARCHAR(100) NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
)
ENGINE=InnoDB
DEFAULT CHARSET=utf8mb3
COLLATE=utf8mb3_general_ci;


### swagger
go get -u github.com/swaggo/swag/cmd/swag
go install github.com/swaggo/swag/cmd/swag

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

swag init

http://localhost:8080/swagger/index.html - docs 확인