/*
Copyright 2017 The Kubernetes Authors.

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

package provider

import (
	"github.com/GoogleCloudPlatform/k8s-stackdriver/event-adapter/pkg/types"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// EventInfo contains info relevant for an Event
type EventInfo struct {
	// GroupResource is the group of the resources
	GroupResource schema.GroupResource
	// Namespaced will be true if the event will be namespaced
	Namespaced bool
	// Event will contain the name of the event
	Event string
}

// EventsProvider is an interface that contains the methods that will provide info for the given events
type EventsProvider interface {
	GetNamespacedEventsByName(namespace, eventName string) (*types.EventValue, error)
	ListAllEvents() (*types.EventValueList, error)
}
