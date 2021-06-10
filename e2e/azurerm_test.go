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
	azureresource "kubeform.dev/provider-azurerm-api/apis/resource/v1alpha1"
	"kubeform.dev/tests/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
)

var _ = Describe("Azurerm", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.Azurerm, whichProvider) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichProvider flag is either `all` or `%s` but got `%s`", framework.Azurerm, framework.Azurerm, whichProvider))
		}
	})
	Describe("Test", func() {
		Context("Azure", func() {
			var (
				providerRef       *core.Secret
				resourceGroupName string
				secretName        string
				resourceGroup     *azureresource.Group
			)

			BeforeEach(func() {
				resourceGroupName = f.GetRandomName("")
				secretName = f.GetRandomName("secret")
				providerRef = f.AzureProviderRef(secretName)
				resourceGroup = f.ResourceGroup(resourceGroupName, secretName)
			})

			It("should create and delete ResourceGroup successfully", func() {
				By("Creating AzureProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating ResourceGroup")
				err = f.CreateResourceGroup(resourceGroup)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Running ResourceGroup")
				f.EventuallyResourceGroupRunning(resourceGroup.ObjectMeta).Should(BeTrue())

				By("Deleting ResourceGroup")
				err = f.DeleteResourceGroup(resourceGroup.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting ResourceGroup")
				f.EventuallyResourceGroupDeleted(resourceGroup.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})
	})
})
