/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package framework

import (
	awsclient "kubeform.dev/provider-aws-api/client/clientset/versioned"
	azurermclient "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	digitaloceanclient "kubeform.dev/provider-digitalocean-api/client/clientset/versioned"
	googleclient "kubeform.dev/provider-google-api/client/clientset/versioned"
	linodeclient "kubeform.dev/provider-linode-api/client/clientset/versioned"

	"github.com/appscode/go/crypto/rand"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	Linode       = "linode"
	Google       = "google"
	AWS          = "aws"
	Azurerm      = "azurerm"
	DigitalOcean = "digitalocean"
	ALL          = "all"
)

type Framework struct {
	restConfig         *rest.Config
	kubeClient         kubernetes.Interface
	linodeClient       linodeclient.Interface
	digitaloceanClient digitaloceanclient.Interface
	googleClient       googleclient.Interface
	awsClient          awsclient.Interface
	azurermClient      azurermclient.Interface
	namespace          string
	name               string
}

func New(
	restConfig *rest.Config,
	kubeClient kubernetes.Interface,
	linodeClient linodeclient.Interface,
	digitaloceanClient digitaloceanclient.Interface,
	googleClient googleclient.Interface,
	awsClient awsclient.Interface,
	azurermClient azurermclient.Interface,
) *Framework {
	return &Framework{
		restConfig:         restConfig,
		kubeClient:         kubeClient,
		linodeClient:       linodeClient,
		digitaloceanClient: digitaloceanClient,
		googleClient:       googleClient,
		awsClient:          awsClient,
		azurermClient:      azurermClient,
		name:               "kfc",
		namespace:          rand.WithUniqSuffix("kubeform"),
	}
}

func (f *Framework) Invoke() *Invocation {
	return &Invocation{
		Framework: f,
		app:       rand.WithUniqSuffix("kfc-e2e"),
	}
}

func (fi *Invocation) KubeClient() kubernetes.Interface {
	return fi.kubeClient
}

func (fi *Invocation) RestConfig() *rest.Config {
	return fi.restConfig
}

func (fi *Invocation) GetRandomName(extraSuffix string) string {
	return rand.WithUniqSuffix(fi.name + extraSuffix)
}

func RunTest(provider, whichProvider string) bool {
	if whichProvider == ALL || provider == whichProvider {
		return true
	} else {
		return false
	}
}

type Invocation struct {
	*Framework
	app string
}
