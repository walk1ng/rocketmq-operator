/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import corev1 "k8s.io/api/core/v1"
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const MinimumVersion = "4.2.0"

type BrokerClusterSpec struct {
	BrokerVersion       string                        `json:"brokerVersion"`
	Namesrvs            string                        `json:"namesrvs"`
	Members             int32                         `json:"members, omitempty"`
	BaseBrokerID        uint32                        `json:"baseBrokerId, omitempty"`
	ClusterMode         string                        `json:"clusterMode`
	BrokerProperties    map[string]string             `json:"brokerProperties, omitempty"`
	NodeSelector        map[string]string             `json:"nodeSelector, omitempty"`
	Affinity            *corev1.Affinity              `json:"affinity, omitempty"`
	VolumeClaimTemplate *corev1.PersistentVolumeClaim `json:"volumeClaimTemplate, omitempty"`
	Config              *corev1.LocalObjectReference  `json:"config,omitempty"`
}

type BrokerClusterConditonType string

const BrokerClusterReady BrokerClusterConditonType = "Ready"

type BrokerClusterCondition struct {
	Type               BrokerClusterConditonType
	Status             corev1.ConditionStatus
	LastTransitionTime metav1.Time
	Reason             string
	Message            string
}

type BrokerClusterStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Conditions        []BrokerClusterCondition
}

// +genclient
// +genclient:noStatus
// +resourceName=brokerclusters
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BrokerCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              BrokerClusterSpec   `json:"spec"`
	Status            BrokerClusterStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BrokerClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []BrokerCluster `json:"items"`
}
