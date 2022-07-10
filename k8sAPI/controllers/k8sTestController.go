package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	_ "k8sAPI/docs" // docs is generated by Swag CLI, you have to import it.
	"k8sAPI/structs"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
)

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @Tags test api
// @Summary k8s Deployment Restart
// @Description deployment restart
// @Param data body structs.InParam true "파라미터 설명"
// @Accept json
// @Produce json
// @Router /k8s/test-restart [post]
func ReStart(c *gin.Context) {
	var input structs.InParam

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inputPodName string = input.PodName
	deploymentsClient := Config().AppsV1().Deployments(K8sNamespace)
	data := fmt.Sprintf(`{"spec":{"template":{"metadata":{"annotations":{"kubectl.kubernetes.io/restartedAt":"%s"}}}}}`, time.Now().String())
	temp, err := deploymentsClient.Patch(context.Background(), inputPodName, types.StrategicMergePatchType, []byte(data), metav1.PatchOptions{FieldManager: "kubectl-rollout"})
	fmt.Println(temp)
	if err != nil {
		return
	}
}

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @Tags test api
// @Summary k8s New Namespace
// @Description create Namespace
// @Param data body structs.ServiceStartParam true "파라미터 설명"
// @Accept json
// @Produce json
// @Router /k8s/test-new-namespace [post]
func NewNamespace(c *gin.Context) {
	var input structs.InParam

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var appName string = input.PodName
	nsName := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: appName,
		},
	}

	Config().CoreV1().Namespaces().Create(context.Background(), nsName, metav1.CreateOptions{})
}

/* 아래 항목이 swagger에 의해 문서화 된다. */
// @Tags test api
// @Summary Summary k8s Deployment Service Create
// @Description create pod & service
// @name CreateKubeService
// @Accept json
// @Produce json
// @Router /k8s/test-create [post]
// @Param data body structs.InParam true "파라미터 설명"
func TempCreate(c *gin.Context) {
	var input structs.InParam

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var inputPodNm string = input.PodName
	deploymentsClient := Config().AppsV1().Deployments(K8sNamespace)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: inputPodNm,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": inputPodNm,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": inputPodNm,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "web",
							Image: "codercom/code-server:latest",
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
	// Create Deployment
	fmt.Println("Creating deployment...")
	result, err := deploymentsClient.Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		c.JSON(400, gin.H{"error": "deployment Create Fail"})
		panic(err)
	}
	fmt.Printf("Created deployment %q.\n", result.GetObjectMeta().GetName())
	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: inputPodNm,
		},
		Spec: apiv1.ServiceSpec{
			Selector: map[string]string{
				"app": inputPodNm,
			},
			Type: "NodePort",
			Ports: []apiv1.ServicePort{
				{
					Port: 8080,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Int,
						IntVal: 8080,
					},
				},
			},
		},
	}

	// Run Service
	fmt.Println("Running service...")
	svc, err := Config().CoreV1().Services(K8sNamespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		c.JSON(400, gin.H{"error": "service execute Fail"})
		// panic(err)
		return
	}
	fmt.Printf("Service running  %q.\n", svc.GetObjectMeta().GetName())

}
