// Copyright (c) 2018 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ShootEvent is a report of an event on a Shoot resource. It generally denotes some state change.
type ShootEvent struct {
	metav1.TypeMeta
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta

	// Required. Time when this Event was first observed.
	EventTime metav1.MicroTime

	// What action was taken/failed regarding to the regarding object.
	// +optional
	Action string

	// Why the action was taken.
	Reason string

	// Type of this event (Normal, Warning), new types could be added in the
	// future.
	// +optional
	Type string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ShootEventList is a list of Event objects.
type ShootEventList struct {
	metav1.TypeMeta
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta

	// Items is a list of schema objects.
	Items []ShootEvent
}

// TODO: think if I want to have we need  an Event series as well, like https://github.com/kubernetes/kubernetes/blob/90237ce89a7c252c44162b841a37d85c7d2ef71f/staging/src/k8s.io/api/events/v1beta1/types.go#L95
