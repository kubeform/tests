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
	"fmt"
	digitaloceandroplet "kubeform.dev/provider-digitalocean-api/apis/droplet/v1alpha1"
	"kubeform.dev/tests/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
)

var _ = Describe("DigitalOcean", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.DigitalOcean, whichProvider) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichProvider flag is either `all` or `%s` but got `%s`", framework.DigitalOcean, framework.DigitalOcean, whichProvider))
		}
	})
	Describe("Test", func() {
		Context("DigitalOcean", func() {
			var (
				providerRef *core.Secret
				secretName  string
				dropletName string
				droplet     *digitaloceandroplet.Droplet
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				dropletName = f.GetRandomName("")
				providerRef = f.DigitalOceanProviderRef(secretName)
				droplet = f.Droplets(dropletName, secretName)
			})

			It("should create and delete Droplet successfully", func() {
				By("Creating DigitalOceanProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating droplet")
				err = f.CreateDroplet(droplet)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running droplet")
				f.EventuallyDropletRunning(droplet.ObjectMeta).Should(BeTrue())

				By("Deleting droplet")
				err = f.DeleteDroplet(droplet.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting droplet")
				f.EventuallyDropletDeleted(droplet.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})
	})
})
