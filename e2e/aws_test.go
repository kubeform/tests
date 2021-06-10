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
	awss3 "kubeform.dev/provider-aws-api/apis/s3/v1alpha1"
	"kubeform.dev/tests/e2e/framework"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
)

var _ = Describe("AWS", func() {
	var (
		err error
		f   *framework.Invocation
	)
	BeforeEach(func() {
		f = root.Invoke()
		if !framework.RunTest(framework.AWS, whichProvider) {
			Skip(fmt.Sprintf("`%s` test is applied only when whichProvider flag is either `all` or `%s` but got `%s`", framework.AWS, framework.AWS, whichProvider))
		}
	})
	Describe("Test", func() {
		Context("AWS", func() {
			var (
				providerRef    *core.Secret
				secretName     string
				dbInstanceName string
				s3Bucket       *awss3.Bucket
			)

			BeforeEach(func() {
				secretName = f.GetRandomName("secret")
				dbInstanceName = f.GetRandomName("")
				providerRef = f.AwsProviderRef(secretName)
				s3Bucket = f.S3Bucket(dbInstanceName, secretName)
			})

			It("should create and delete s3 bucket successfully", func() {
				By("Creating AwsProviderRef")
				err = f.CreateSecret(providerRef)
				Expect(err).NotTo(HaveOccurred())

				By("Creating S3Bucket")
				err = f.CreateS3Bucket(s3Bucket)
				Expect(err).NotTo(HaveOccurred())

				By("Waiting for Running S3 Bucket")
				f.EventuallyS3BucketRunning(s3Bucket.ObjectMeta).Should(BeTrue())

				By("Deleting S3 Bucket")
				err = f.DeleteS3Bucket(s3Bucket.ObjectMeta)
				Expect(err).NotTo(HaveOccurred())

				By("Wait for Deleting S3 Bucket")
				f.EventuallyS3BucketDeleted(s3Bucket.ObjectMeta).Should(BeTrue())

				By("Deleting secret")
				err = f.DeleteSecret(providerRef.ObjectMeta)
			})
		})
	})
})
