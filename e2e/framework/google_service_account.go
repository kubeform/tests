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
	"time"

	base "kubeform.dev/apimachinery/api/v1alpha1"
	"kubeform.dev/provider-google-api/apis/service/v1alpha1"

	. "github.com/onsi/gomega"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
)

func (i *Invocation) ServiceAccount(name string, secretName string) *v1alpha1.Account {
	project := "appscode-testing"

	return &v1alpha1.Account{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.AccountSpec{
			AccountSpec2: v1alpha1.AccountSpec2{
				AccountID: &name,
				ProviderRef: v12.LocalObjectReference{
					Name: secretName,
				},
				DisplayName: &name,
				Project:     &project,
			},
		},
	}
}

func (f *Framework) CreateServiceAccount(obj *v1alpha1.Account) error {
	_, err := f.googleClient.ServiceV1alpha1().Accounts(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) DeleteServiceAccount(meta metav1.ObjectMeta) error {
	return f.googleClient.ServiceV1alpha1().Accounts(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInForeground())
}

func (f *Framework) EventuallyServiceAccountRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			serviceAccount, err := f.googleClient.ServiceV1alpha1().Accounts(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return serviceAccount.Status.Phase == base.PhaseRunning
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyServiceAccountDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.googleClient.ServiceV1alpha1().Accounts(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
