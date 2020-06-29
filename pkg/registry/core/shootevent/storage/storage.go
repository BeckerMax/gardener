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

package storage

import (
	"github.com/gardener/gardener/pkg/registry/core/shootevent"

	"github.com/gardener/gardener/pkg/apis/core"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// REST implements a RESTStorage for shootEvents against etcd.
type REST struct {
	*genericregistry.Store
}

// ShootEvent implements the storage for ShootEvents
// TODO Polish comments for public stuff
type ShootEvent struct {
	ShootEvent *REST
}

// NewStorage creates a new ShootEvent object.
func NewStorage(optsGetter generic.RESTOptionsGetter, cloudProfiles rest.StandardStorage) ShootEvent {
	// TODO Can I do this inline?
	shootEventRest := NewREST(optsGetter, cloudProfiles)

	// TODO why all these indirections
	return ShootEvent{
		ShootEvent: shootEventRest,
	}
}

// NewREST returns a RESTStorage object that will work with Shootevent objects.
func NewREST(optsGetter generic.RESTOptionsGetter, cloudProfiles rest.StandardStorage) *REST {
	// TODO implement NewStrategy. Function could be put in shootevent package
	strategy := shootevent.NewStrategy(cloudProfiles)
	// statusStrategy := seed.NewStatusStrategy(cloudProfiles)

	store := &genericregistry.Store{
		NewFunc:                  func() runtime.Object { return &core.ShootEvent{} },
		NewListFunc:              func() runtime.Object { return &core.ShootEventList{} },
		DefaultQualifiedResource: core.Resource("shootevents"),
		EnableGarbageCollection:  true,

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: newTableConvertor(),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter}
	if err := store.CompleteWithOptions(options); err != nil {
		panic(err)
	}

	return &REST{store}
}