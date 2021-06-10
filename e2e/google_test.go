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
	googleservice "kubeform.dev/provider-google-api/apis/service/v1alpha1"
	"kubeform.dev/tests/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
)

var _ = Describe("Google", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.Google, whichProvider) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichProvider flag is either `all` or `%s` but got `%s`", framework.Google, framework.Google, whichProvider))
		}
	})
	Describe("Test", func() {
		Context("Google", func() {
			var (
				providerRef        *core.Secret
				serviceAccountName string
				secretName         string
				serviceAccount     *googleservice.Account
			)

			BeforeEach(func() {
				serviceAccountName = f.GetRandomName("")
				secretName = f.GetRandomName("secret")
				providerRef = f.GoogleProviderRef(secretName)
				serviceAccount = f.ServiceAccount(serviceAccountName, secretName)
			})

			It("should create and delete service account successfully", func() {
				By("Creating GoogleProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating ServiceAccount")
				err = f.CreateServiceAccount(serviceAccount)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Running ServiceAccount")
				f.EventuallyServiceAccountRunning(serviceAccount.ObjectMeta).Should(BeTrue())

				By("Deleting service account")
				err = f.DeleteServiceAccount(serviceAccount.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Deleting ServiceAccount")
				f.EventuallyServiceAccountDeleted(serviceAccount.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})
	})
})
