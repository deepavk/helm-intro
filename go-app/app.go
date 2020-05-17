package main

import (
	"fmt"
	"html"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func kcontrol() {
	ctx, err, kClient := initKubeClient()
	if err != nil {
		log.Fatal("Init client failed with err %s", err)
	}
	//listPods(kClient, ctx)
	//getSecrets(kClient, ctx)

	// List secrets
	listSecrets(kClient, ctx)
}

func getSecrets(kClient v1.CoreV1Interface, ctx context.Context) {
	sec, err := kClient.Secrets(ns).Get(ctx, "tenant-secret", metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	for name, val := range sec.Data {
		v := string(val)
		fmt.Printf("name :%s , data %s ", name, v)
	}
}

func listSecrets(kClient v1.CoreV1Interface, ctx context.Context) {
	labelsMap := map[string]string{"secretResourceType": "TenantInfo", "tenantID": "t1"}
	labelSelector := metav1.LabelSelector{
		MatchLabels: labelsMap,
	}
	res, _ := kClient.Secrets(ns).List(ctx, metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	})

	fmt.Printf("Secret %+v", res)
}

func listPods(kClient v1.CoreV1Interface, ctx context.Context) {
	pods, err := kClient.Pods(ns).List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//fmt.Printf(pods.String())
	for _, pod := range pods.Items {
		fmt.Printf("Pod name=/%s\n", pod.GetName())
	}
}

func initKubeClient() (context.Context, error, v1.CoreV1Interface) {
	ctx := context.Background()
	kubeconfig := "/home/deepa/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	kClient := clientset.CoreV1()
	return ctx, err, kClient
}

//func (k KubernetesAPI) GetSecrets(labelsMap map[string]string) ([]SecureObject, error) {
//    var secretResources []SecureObject
//
//    labelSelector := metav1.LabelSelector{
//        MatchLabels: labelsMap,
//    }
//
//    res, err := k.kubeClient.CoreV1().Secrets(namespace).List(metav1.ListOptions{
//        LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
//    })
//
//    if err != nil {
//        secretsHandlerLogger.Errorf("Error fetching secrets. Error Details : %v", err.Error())
//        return nil, err
//    }
//
//    for _, item := range res.Items {
//        secretResource := SecureObject{
//            Name:               item.ObjectMeta.Name,
//            ResourceType:       item.ObjectMeta.Labels[SecretResourceType],
//            ResponseBinaryData: item.Data,
//            ResponseStringData: item.StringData,
//        }
//        secretResources = append(secretResources, secretResource)
//    }
//    return secretResources, nil
//}
