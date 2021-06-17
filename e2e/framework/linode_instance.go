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

	"kubeform.dev/provider-linode-api/apis/instance/v1alpha1"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"

	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
)

func (i *Invocation) Instance(name string, secretName string) *v1alpha1.Instance {
	image := "linode/ubuntu18.04"
	region := "us-east"

	return &v1alpha1.Instance{
		ObjectMeta: v1.ObjectMeta{
			Name:      name,
			Namespace: i.Namespace(),
			Labels: map[string]string{
				"app": i.app,
			},
		},
		Spec: v1alpha1.InstanceSpec{
			InstanceSpec2: v1alpha1.InstanceSpec2{
				ProviderRef: corev1.LocalObjectReference{
					Name: secretName,
				},
				SecretRef: &corev1.LocalObjectReference{
					Name: InstanceSecretName,
				},
				Image:  &image,
				Label:  &name,
				Region: &region,
			},
		},
	}
}

func (f *Framework) CreateInstance(obj *v1alpha1.Instance) error {
	_, err := f.linodeClient.InstanceV1alpha1().Instances(obj.Namespace).Create(context.TODO(), obj, metav1.CreateOptions{})
	return err
}

func (f *Framework) DeleteInstance(meta metav1.ObjectMeta) error {
	return f.linodeClient.InstanceV1alpha1().Instances(meta.Namespace).Delete(context.TODO(), meta.Name, meta_util.DeleteInForeground())
}

func (f *Framework) EventuallyInstanceRunning(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			instance, err := f.linodeClient.InstanceV1alpha1().Instances(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			Expect(err).NotTo(HaveOccurred())
			return instance.Status.Phase == status.CurrentStatus
		},
		time.Minute*15,
		time.Second*10,
	)
}

func (f *Framework) EventuallyInstanceDeleted(meta metav1.ObjectMeta) GomegaAsyncAssertion {
	return Eventually(
		func() bool {
			_, err := f.linodeClient.InstanceV1alpha1().Instances(meta.Namespace).Get(context.TODO(), meta.Name, metav1.GetOptions{})
			return errors.IsNotFound(err)
		},
		time.Minute*15,
		time.Second*10,
	)
}
