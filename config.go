package kunsul

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"

	log "github.com/Sirupsen/logrus"
)

var err error

func GetConfig(outside bool) (config *rest.Config) {
	if outside {
		log.Debug("using outside-cluster authentication")
		loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
		configOverrides := &clientcmd.ConfigOverrides{}
		kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
		config, err := kubeConfig.ClientConfig()
		if err != nil {
			log.Error(err)
		}
		log.Debug(config)

		return config
	} else {
		log.Debug("using in-cluster authentication")
		config, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}
		log.Debug(config)

		return config
	}
}
