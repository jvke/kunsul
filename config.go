package kunsul

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	log "github.com/Sirupsen/logrus"
)

var err error

func GetConfig(outside bool) (config *rest.Config, err error) {
	if outside {
		log.Debugf("KUBE:> %s","using outside-cluster authentication")
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		configOverrides := &clientcmd.ConfigOverrides{}
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
		if config, err = kubeConfig.ClientConfig(); err != nil {
			log.Error(err)
			return nil, err
		}
	} else {
		log.Debugf("KUBE:> %s","using in-cluster authentication")
		if config, err = rest.InClusterConfig(); err != nil {
			log.Error(err)
			return nil, err
		}
	}

	log.Debugf("CONFIG:> %s", config)
	return config, nil
}
