package kunsul

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	log "github.com/Sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func GetIngresses(config *rest.Config) []v1beta1.Ingress {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	ingressList, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	ingressCtrls := ingressList.Items

	log.Debugf("there are %s ingresses in the cluster", len(ingressCtrls))

	for k, v := range ingressCtrls {
		log.Debugf("key:", k, "value:", v)
	}

	return ingressCtrls
}
