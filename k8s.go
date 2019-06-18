package kunsul

import (
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	log "github.com/Sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func GetIngresses(config *rest.Config) ([]v1beta1.Ingress, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	ingressList, err := clientset.ExtensionsV1beta1().Ingresses("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	ingressCtrls := ingressList.Items

	log.Debugf("DISCOVERED INGRESSES:> there are %s ingresses in the cluster", len(ingressCtrls))

	for k, v := range ingressCtrls {
		log.Debugf("INGRESS %s:> %s", k, v)
	}

	return ingressCtrls, nil
}

func GetServices(config *rest.Config) ([]v1.Service, error) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	servicesList, err := clientset.CoreV1().Services("").List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	serviceCtrls := servicesList.Items

	log.Debugf("DISCOVERED Services:> there are %s services in the cluster", len(serviceCtrls))

	for k, v := range serviceCtrls {
		log.Debugf("SERVICE %s:> %s", k, v)
	}

	return serviceCtrls, nil
}
