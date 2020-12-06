package testsuites

import (
	"github.com/c2devel/aws-ebs-csi-driver/tests/e2e/driver"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clientset "k8s.io/client-go/kubernetes"

	"k8s.io/client-go/util/retry"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DynamicallyProvisionedDeploymentUpdateTest struct {
	CSIDriver     driver.DynamicPVTestDriver
	Deployment    *appsv1.Deployment
	UpdateImageTo string
}

func int32Ptr(i int32) *int32 { return &i }

func (t *DynamicallyProvisionedDeploymentUpdateTest) Run(client clientset.Interface, namespace *v1.Namespace) {

	deploymentsClient := client.AppsV1().Deployments(namespace.Name)

	By("create deployment")
	result, createErr := deploymentsClient.Create(t.Deployment)
	Expect(createErr).ShouldNot(HaveOccurred())

	oldImage := result.Spec.Template.Spec.Containers[0].Image
	deploymentName := result.Name

	By("update deployment")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
		Expect(getErr).ShouldNot(HaveOccurred())
		result.Spec.Replicas = int32Ptr(1)
		result.Spec.Template.Spec.Containers[0].Image = t.UpdateImageTo
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})

	By("verify update completion")
	Expect(retryErr).ShouldNot(HaveOccurred())
	updatedResults, getErr := deploymentsClient.Get("demo-deployment", metav1.GetOptions{})
	Expect(getErr).ShouldNot(HaveOccurred())
	newImage := updatedResults.Spec.Template.Spec.Containers[0].Image

	Expect(newImage).NotTo(ContainSubstring(oldImage))
	Expect(newImage).To(ContainSubstring(t.UpdateImageTo))

	By("delete deployment")
	deletePolicy := metav1.DeletePropagationForeground
	deleteErr := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	Expect(deleteErr).ShouldNot(HaveOccurred())
}
