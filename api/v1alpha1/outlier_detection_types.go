// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package v1alpha1

import (
	duration "github.com/golang/protobuf/ptypes/duration"
)

type OutlierDetection struct {
	// Determines whether to distinguish local origin failures from external errors. 
	// If set to true the following configuration parameters are taken into account: 
	// consecutive_local_origin_failure, enforcing_consecutive_local_origin_failure and 
	// enforcing_local_origin_success_rate. Defaults to false.
	SplitExternalLocalOriginErrors bool `json:"split_external_local_origin_errors,omitempty"`
	// The number of consecutive locally originated failures before ejection
	// occurs. Defaults to 5. Parameter takes effect only when split_external_local_origin_errors
	// is set to true.
	ConsecutiveLocalOriginFailures int `json:"consecutive_local_origin_failures,omitempty"`
	// The number of consecutive gateway failures (502, 503, 504 status codes) before a 
	// consecutive gateway failure ejection occurs. Defaults to 5.
	ConsecutiveGatewayErrors int `json:"consecutive_gateway_errors,omitempty"`
	// Number of 5xx errors before a host is ejected from the connection pool.
	// When the upstream host is accessed over an opaque TCP connection, connect
	// timeouts, connection error/failure and request failure events qualify as a
	// 5xx error.
	// This feature defaults to 5 but can be disabled by setting the value to 0.
	//
	// Note that consecutive_gateway_errors and consecutive_5xx_errors can be
	// used separately or together. Because the errors counted by
	// consecutive_gateway_errors are also included in consecutive_5xx_errors,
	// if the value of consecutive_gateway_errors is greater than or equal to
	// the value of consecutive_5xx_errors, consecutive_gateway_errors will have
	// no effect.
	Consecutive_5XxErrors int `json:"consecutive_5xx_errors,omitempty"`
	// Time interval between ejection sweep analysis. format:
	// 1h/1m/1s/1ms. MUST BE >=1ms. Default is 10s.
	Interval *duration.Duration `json:"interval,omitempty"`
	// Minimum ejection duration. A host will remain ejected for a period
	// equal to the product of minimum ejection duration and the number of
	// times the host has been ejected. This technique allows the system to
	// automatically increase the ejection period for unhealthy upstream
	// servers. format: 1h/1m/1s/1ms. MUST BE >=1ms. Default is 30s.
	BaseEjectionTime *duration.Duration `json:"base_ejection_time,omitempty"`
	// Maximum % of hosts in the load balancing pool for the upstream
	// service that can be ejected. Defaults to 10%.
	MaxEjectionPercent int32 `json:"max_ejection_percent,omitempty"`
	// Outlier detection will be enabled as long as the associated load balancing
	// pool has at least min_health_percent hosts in healthy mode. When the
	// percentage of healthy hosts in the load balancing pool drops below this
	// threshold, outlier detection will be disabled and the proxy will load balance
	// across all hosts in the pool (healthy and unhealthy). The threshold can be
	// disabled by setting it to 0%. The default is 0% as it's not typically
	// applicable in k8s environments with few pods per service.
	MinHealthPercent int32 `json:"min_health_percent,omitempty"`
}