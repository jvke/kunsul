package kunsul

import (
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"

	log "github.com/sirupsen/logrus"
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
