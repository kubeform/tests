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
	"context"
	"errors"
	"time"

	. "github.com/onsi/gomega"
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (f *Framework) EventuallyCRD(whichProvider string) GomegaAsyncAssertion {
	return Eventually(
		func() error {
			if whichProvider == "all" || whichProvider == "google" {
				// Check ServiceAccount CRD
				if _, err := f.googleClient.ServiceV1alpha1().Accounts(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD ServiceAccount is not ready")
				}
			}

			if whichProvider == "all" || whichProvider == "azurerm" {
				// Check ResourceGroup CRD
				if _, err := f.azurermClient.ResourceV1alpha1().Groups(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD ResourceGroup is not ready")
				}
			}

			if whichProvider == "all" || whichProvider == "aws" {
				// Check S3Buckets CRD
				if _, err := f.awsClient.S3V1alpha1().Buckets(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD S3Buckets is not ready")
				}
			}

			if whichProvider == "all" || whichProvider == "linode" {
				// Check Instances CRD
				if _, err := f.linodeClient.InstanceV1alpha1().Instances(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD Instances is not ready")
				}
			}

			if whichProvider == "all" || whichProvider == "digitalocean" {
				// Check Droplets CRD
				if _, err := f.digitaloceanClient.DropletV1alpha1().Droplets(core.NamespaceAll).List(context.TODO(), metav1.ListOptions{}); err != nil {
					return errors.New("CRD Droplets is not ready")
				}
			}

			return nil
		},
		time.Minute*2,
		time.Second*10,
	)
}
