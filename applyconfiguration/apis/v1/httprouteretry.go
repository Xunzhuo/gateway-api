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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	apisv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// HTTPRouteRetryApplyConfiguration represents a declarative configuration of the HTTPRouteRetry type for use
// with apply.
type HTTPRouteRetryApplyConfiguration struct {
	Codes    []apisv1.HTTPRouteRetryStatusCode `json:"codes,omitempty"`
	Attempts *int                              `json:"attempts,omitempty"`
	Backoff  *apisv1.Duration                  `json:"backoff,omitempty"`
}

// HTTPRouteRetryApplyConfiguration constructs a declarative configuration of the HTTPRouteRetry type for use with
// apply.
func HTTPRouteRetry() *HTTPRouteRetryApplyConfiguration {
	return &HTTPRouteRetryApplyConfiguration{}
}

// WithCodes adds the given value to the Codes field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Codes field.
func (b *HTTPRouteRetryApplyConfiguration) WithCodes(values ...apisv1.HTTPRouteRetryStatusCode) *HTTPRouteRetryApplyConfiguration {
	for i := range values {
		b.Codes = append(b.Codes, values[i])
	}
	return b
}

// WithAttempts sets the Attempts field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Attempts field is set to the value of the last call.
func (b *HTTPRouteRetryApplyConfiguration) WithAttempts(value int) *HTTPRouteRetryApplyConfiguration {
	b.Attempts = &value
	return b
}

// WithBackoff sets the Backoff field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Backoff field is set to the value of the last call.
func (b *HTTPRouteRetryApplyConfiguration) WithBackoff(value apisv1.Duration) *HTTPRouteRetryApplyConfiguration {
	b.Backoff = &value
	return b
}
