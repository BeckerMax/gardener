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

package shootevent

import (
	"context"
	"github.com/gardener/gardener/pkg/api"
	"github.com/gardener/gardener/pkg/apis/core"
	"github.com/gardener/gardener/pkg/apis/core/validation"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/rest"
	"k8s.io/apiserver/pkg/storage/names"
)

// TODO can't I just completely reuse the eventStrategy from K8s?

// Strategy defines the strategy for storing seeds.
type Strategy struct {
	runtime.ObjectTyper
	names.NameGenerator

	CloudProfiles rest.StandardStorage
}

// NewStrategy defines the storage strategy for ShootEvents.
func NewStrategy(cloudProfiles rest.StandardStorage) Strategy {
	return Strategy{api.Scheme, names.SimpleNameGenerator, cloudProfiles}
}

// NamespaceScoped returns true if the object must be within a namespace.
func (Strategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate is invoked on create before validation to normalize
// the object.  For example: remove fields that are not to be persisted,
// sort order-insensitive list fields, etc.  This should not remove fields
// whose presence would be considered a validation error.
//
// Often implemented as a type check and an initailization or clearing of
// status. Clear the status because status changes are internal. External
// callers of an api (users) should not be setting an initial status on
// newly created objects.
func (s Strategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (s Strategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

// Validate returns an ErrorList with validation errors or nil.  Validate
// is invoked after default fields in the object have been filled in
// before the object is persisted.  This method should not mutate the
// object.
func  (Strategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	event := obj.(*core.ShootEvent)
	return validation.ValidateShootEvent(event)
}

// Canonicalize allows an object to be mutated into a canonical form. This
// ensures that code that operates on these objects can rely on the common
// form for things like comparison.  Canonicalize is invoked after
// validation has succeeded but before the object has been persisted.
// This method may mutate the object. Often implemented as a type check or
// empty method.
func (Strategy) Canonicalize(obj runtime.Object) {
}

func (Strategy) AllowCreateOnUpdate() bool {
	return true
}

func (Strategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	event := obj.(*core.ShootEvent)
	return validation.ValidateShootEvent(event)
}

func (Strategy) AllowUnconditionalUpdate() bool {
	return true
}

// TODO do I need the following? Its in the original K8s event
// // GetAttrs returns labels and fields of a given object for filtering purposes.
// func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
// 	event, ok := obj.(*api.Event)
// 	if !ok {
// 		return nil, nil, fmt.Errorf("not an event")
// 	}
// 	return labels.Set(event.Labels), ToSelectableFields(event), nil
// }

// // Matcher returns a selection predicate for a given label and field selector.
// func Matcher(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
// 	return storage.SelectionPredicate{
// 		Label:    label,
// 		Field:    field,
// 		GetAttrs: GetAttrs,
// 	}
// }

// // ToSelectableFields returns a field set that represents the object.
// func ToSelectableFields(event *api.Event) fields.Set {
// 	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(&event.ObjectMeta, true)
// 	specificFieldsSet := fields.Set{
// 		"involvedObject.kind":            event.InvolvedObject.Kind,
// 		"involvedObject.namespace":       event.InvolvedObject.Namespace,
// 		"involvedObject.name":            event.InvolvedObject.Name,
// 		"involvedObject.uid":             string(event.InvolvedObject.UID),
// 		"involvedObject.apiVersion":      event.InvolvedObject.APIVersion,
// 		"involvedObject.resourceVersion": event.InvolvedObject.ResourceVersion,
// 		"involvedObject.fieldPath":       event.InvolvedObject.FieldPath,
// 		"reason":                         event.Reason,
// 		"source":                         event.Source.Component,
// 		"type":                           event.Type,
// 	}
// 	return generic.MergeFieldsSets(objectMetaFieldsSet, specificFieldsSet)
// }
