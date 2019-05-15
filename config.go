package kunsul

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	log "github.com/Sirupsen/logrus"
)

var err error

func GetConfig() (config *rest.Config, err error) {


	log.Debugf("KUBE:> %s","try using in-cluster authentication")
	if config, err = rest.InClusterConfig(); err != nil {
		if err != rest.ErrNotInCluster {
			log.Error(err)
			return nil, err
		}
	}

	// else not in cluster
	log.Debugf("KUBE:> %s","try outside-cluster authentication")
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	if config, err = kubeConfig.ClientConfig(); err != nil {
		log.Error(err)
		return nil, err
	}

	log.Debugf("CONFIG:> %s", config)
	return config, nil
}
