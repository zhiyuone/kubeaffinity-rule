package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:webhook:path=/mutate-v1alpha1-pod,mutating=true,failurePolicy=Ignore,groups="",resources=pods,verbs=create;update,versions=v1,name=m.affinities.rule.newbis.top

// AffinityAnnotationKey is the annotation key which affinities will be used.
const AffinityAnnotationKey string = "rule.newbis.top/affinity"

// PodAffinityRuler manager pod affinity
type PodAffinityRuler struct {
	Client  client.Client
	decoder *admission.Decoder
}

// InjectDecoder injects the decoder.
func (a *PodAffinityRuler) InjectDecoder(d *admission.Decoder) error {
	a.decoder = d
	return nil
}

// Handle manage pod affinity by affinity rules.
func (a *PodAffinityRuler) Handle(ctx context.Context, req admission.Request) admission.Response {
	pod := &corev1.Pod{}

	err := a.decoder.Decode(req, pod)
	if err != nil {
		return admission.Errored(http.StatusBadRequest, err)
	}

	var ns corev1.Namespace
	if pod.Namespace != "" {
		a.Client.Get(ctx, types.NamespacedName{Name: pod.Namespace}, &ns)
	}

	if !mutationRequired(&ns.ObjectMeta) && !mutationRequired(&pod.ObjectMeta) {
		return admission.Allowed(fmt.Sprintf("Not affinity rule require in %s/%s annotation", pod.Namespace, pod.Name))
	}

	// TODO

	marshaledPod, err := json.Marshal(pod)
	if err != nil {
		return admission.Errored(http.StatusInternalServerError, err)
	}
	return admission.PatchResponseFromRaw(req.Object.Raw, marshaledPod)
}

func mutationRequired(metadata *metav1.ObjectMeta) bool {
	if _, ok := metadata.Annotations[AffinityAnnotationKey]; ok {
		return true
	}
	return false
}

func mergeAffinity() {

}
