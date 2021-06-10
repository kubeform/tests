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

package e2e_test

import (
	"flag"
	"fmt"
	awsclient "kubeform.dev/provider-aws-api/client/clientset/versioned"
	azurermclient "kubeform.dev/provider-azurerm-api/client/clientset/versioned"
	digitaloceanclient "kubeform.dev/provider-digitalocean-api/client/clientset/versioned"
	googleclient "kubeform.dev/provider-google-api/client/clientset/versioned"
	linodeclient "kubeform.dev/provider-linode-api/client/clientset/versioned"
	"os"
	"path/filepath"
	"testing"
	"time"

	"kubeform.dev/tests/e2e/framework"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	clientSetScheme "k8s.io/client-go/kubernetes/scheme"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/scale/scheme"
	"k8s.io/client-go/util/homedir"
	"kmodules.xyz/client-go/tools/clientcmd"
)

var (
	kubeconfigPath = func() string {
		kubecfg := os.Getenv("KUBECONFIG")
		if kubecfg != "" {
			return kubecfg
		}
		return filepath.Join(homedir.HomeDir(), ".kube", "config")
	}()
	kubeContext   = ""
	whichProvider = func() string {
		whichPrv := os.Getenv("WHICH_PROVIDER")
		if whichPrv != "" {
			return whichPrv
		}
		return "all"
	}()
)

func init() {
	utilruntime.Must(scheme.AddToScheme(clientSetScheme.Scheme))

	flag.StringVar(&kubeconfigPath, "kube-config", kubeconfigPath, "Path to kubeconfig file with authorization information (the master location is set by the master flag).")
	flag.StringVar(&kubeContext, "kube-context", "", "Name of kube context")
	flag.StringVar(&whichProvider, "which-provider", whichProvider, "Define which provider resource you want to create")
}

const (
	TIMEOUT = 20 * time.Minute
)

var (
	root *framework.Framework
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(TIMEOUT)

	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "e2e Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	By("Using kubeconfig from " + kubeconfigPath)
	config, err := clientcmd.BuildConfigFromContext(kubeconfigPath, kubeContext)
	Expect(err).NotTo(HaveOccurred())
	// raise throttling time. ref: https://github.com/appscode/voyager/issues/640
	config.Burst = 100
	config.QPS = 100

	// Clients
	kubeClient := kubernetes.NewForConfigOrDie(config)

	linodeClient := linodeclient.NewForConfigOrDie(config)
	digitaloceanClient := digitaloceanclient.NewForConfigOrDie(config)
	googleClient := googleclient.NewForConfigOrDie(config)
	awsClient := awsclient.NewForConfigOrDie(config)
	azurermClient := azurermclient.NewForConfigOrDie(config)

	// Framework
	root = framework.New(config, kubeClient, linodeClient, digitaloceanClient, googleClient, awsClient, azurermClient)

	// Create namespace
	By("Using namespace " + root.Namespace())
	err = root.CreateNamespace()
	Expect(err).NotTo(HaveOccurred())
	fmt.Println(whichProvider)
	root.EventuallyCRD(whichProvider).Should(Succeed())
})

var _ = AfterSuite(func() {
	By("Deleting Namespace")
	err := root.DeleteNamespace()
	Expect(err).NotTo(HaveOccurred())
})
