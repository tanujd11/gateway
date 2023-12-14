// Copyright Envoy Gateway Authors
// SPDX-License-Identifier: Apache-2.0
// The full text of the Apache license is available in the LICENSE file at
// the root of the repo.

package proxy

import (
	"fmt"
	"os"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"
	autoscalingv2 "k8s.io/api/autoscaling/v2"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/utils/pointer"
	"sigs.k8s.io/yaml"

	egv1a1 "github.com/envoyproxy/gateway/api/v1alpha1"
	"github.com/envoyproxy/gateway/internal/envoygateway/config"
	"github.com/envoyproxy/gateway/internal/gatewayapi"
	"github.com/envoyproxy/gateway/internal/ir"
	"github.com/envoyproxy/gateway/internal/utils/ptr"
)

const (
	// envoyHTTPPort is the container port number of Envoy's HTTP endpoint.
	envoyHTTPPort = int32(8080)
	// envoyHTTPSPort is the container port number of Envoy's HTTPS endpoint.
	envoyHTTPSPort = int32(8443)
)

func newTestInfra() *ir.Infra {
	i := ir.NewInfra()

	i.Proxy.GetProxyMetadata().Labels[gatewayapi.OwningGatewayNamespaceLabel] = "default"
	i.Proxy.GetProxyMetadata().Labels[gatewayapi.OwningGatewayNameLabel] = i.Proxy.Name
	i.Proxy.Listeners = []ir.ProxyListener{
		{
			Ports: []ir.ListenerPort{
				{
					Name:          "EnvoyHTTPPort",
					Protocol:      ir.TCPProtocolType,
					ContainerPort: envoyHTTPPort,
				},
				{
					Name:          "EnvoyHTTPSPort",
					Protocol:      ir.TCPProtocolType,
					ContainerPort: envoyHTTPSPort,
				},
			},
		},
	}

	return i
}

func TestDeployment(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	cases := []struct {
		caseName     string
		infra        *ir.Infra
		deploy       *egv1a1.KubernetesDeploymentSpec
		proxyLogging map[egv1a1.ProxyLogComponent]egv1a1.LogLevel
		bootstrap    string
		telemetry    *egv1a1.ProxyTelemetry
		concurrency  *int32
	}{
		{
			caseName: "default",
			infra:    newTestInfra(),
			deploy:   nil,
		},
		{
			caseName: "custom",
			infra:    newTestInfra(),
			deploy: &egv1a1.KubernetesDeploymentSpec{
				Replicas: pointer.Int32(2),
				Strategy: egv1a1.DefaultKubernetesDeploymentStrategy(),
				Pod: &egv1a1.KubernetesPodSpec{
					Annotations: map[string]string{
						"prometheus.io/scrape": "true",
					},
					Labels: map[string]string{
						"foo.bar": "custom-label",
					},
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: pointer.Int64(1000),
					},
					HostNetwork: true,
				},
				Container: &egv1a1.KubernetesContainerSpec{
					Image: pointer.String("envoyproxy/envoy:v1.2.3"),
					Resources: &corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("400m"),
							corev1.ResourceMemory: resource.MustParse("2Gi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("200m"),
							corev1.ResourceMemory: resource.MustParse("1Gi"),
						},
					},
					SecurityContext: &corev1.SecurityContext{
						Privileged: pointer.Bool(true),
					},
				},
			},
		},
		{
			caseName:  "bootstrap",
			infra:     newTestInfra(),
			deploy:    nil,
			bootstrap: `test bootstrap config`,
		},
		{
			caseName: "extension-env",
			infra:    newTestInfra(),
			deploy: &egv1a1.KubernetesDeploymentSpec{
				Replicas: pointer.Int32(2),
				Strategy: egv1a1.DefaultKubernetesDeploymentStrategy(),
				Pod: &egv1a1.KubernetesPodSpec{
					Annotations: map[string]string{
						"prometheus.io/scrape": "true",
					},
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: pointer.Int64(1000),
					},
				},
				Container: &egv1a1.KubernetesContainerSpec{
					Env: []corev1.EnvVar{
						{
							Name:  "env_a",
							Value: "env_a_value",
						},
						{
							Name:  "env_b",
							Value: "env_b_value",
						},
					},
					Image: pointer.String("envoyproxy/envoy:v1.2.3"),
					Resources: &corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("400m"),
							corev1.ResourceMemory: resource.MustParse("2Gi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("200m"),
							corev1.ResourceMemory: resource.MustParse("1Gi"),
						},
					},
					SecurityContext: &corev1.SecurityContext{
						Privileged: pointer.Bool(true),
					},
				},
			},
		},
		{
			caseName: "default-env",
			infra:    newTestInfra(),
			deploy: &egv1a1.KubernetesDeploymentSpec{
				Replicas: pointer.Int32(2),
				Strategy: egv1a1.DefaultKubernetesDeploymentStrategy(),
				Pod: &egv1a1.KubernetesPodSpec{
					Annotations: map[string]string{
						"prometheus.io/scrape": "true",
					},
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: pointer.Int64(1000),
					},
				},
				Container: &egv1a1.KubernetesContainerSpec{
					Env:   nil,
					Image: pointer.String("envoyproxy/envoy:v1.2.3"),
					Resources: &corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("400m"),
							corev1.ResourceMemory: resource.MustParse("2Gi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("200m"),
							corev1.ResourceMemory: resource.MustParse("1Gi"),
						},
					},
					SecurityContext: &corev1.SecurityContext{
						Privileged: pointer.Bool(true),
					},
				},
			},
		},
		{
			caseName: "volumes",
			infra:    newTestInfra(),
			deploy: &egv1a1.KubernetesDeploymentSpec{
				Replicas: pointer.Int32(2),
				Strategy: egv1a1.DefaultKubernetesDeploymentStrategy(),
				Pod: &egv1a1.KubernetesPodSpec{
					Annotations: map[string]string{
						"prometheus.io/scrape": "true",
					},
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: pointer.Int64(1000),
					},
					Volumes: []corev1.Volume{
						{
							Name: "certs",
							VolumeSource: corev1.VolumeSource{
								Secret: &corev1.SecretVolumeSource{
									SecretName:  "custom-envoy-cert",
									DefaultMode: pointer.Int32(420),
								},
							},
						},
					},
				},
				Container: &egv1a1.KubernetesContainerSpec{
					Env: []corev1.EnvVar{
						{
							Name:  "env_a",
							Value: "env_a_value",
						},
						{
							Name:  "env_b",
							Value: "env_b_value",
						},
					},
					Image: pointer.String("envoyproxy/envoy:v1.2.3"),
					Resources: &corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("400m"),
							corev1.ResourceMemory: resource.MustParse("2Gi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("200m"),
							corev1.ResourceMemory: resource.MustParse("1Gi"),
						},
					},
					SecurityContext: &corev1.SecurityContext{
						Privileged: pointer.Bool(true),
					},
				},
			},
		},
		{
			caseName: "component-level",
			infra:    newTestInfra(),
			deploy:   nil,
			proxyLogging: map[egv1a1.ProxyLogComponent]egv1a1.LogLevel{
				egv1a1.LogComponentDefault: egv1a1.LogLevelError,
				egv1a1.LogComponentFilter:  egv1a1.LogLevelInfo,
			},
			bootstrap: `test bootstrap config`,
		},
		{
			caseName: "enable-prometheus",
			infra:    newTestInfra(),
			telemetry: &egv1a1.ProxyTelemetry{
				Metrics: &egv1a1.ProxyMetrics{
					Prometheus: &egv1a1.ProxyPrometheusProvider{},
				},
			},
		},
		{
			caseName:    "with-concurrency",
			infra:       newTestInfra(),
			deploy:      nil,
			concurrency: pointer.Int32(4),
			bootstrap:   `test bootstrap config`,
		},
		{
			caseName: "custom_with_initcontainers",
			infra:    newTestInfra(),
			deploy: &egv1a1.KubernetesDeploymentSpec{
				Replicas: pointer.Int32(3),
				Strategy: egv1a1.DefaultKubernetesDeploymentStrategy(),
				Pod: &egv1a1.KubernetesPodSpec{
					Annotations: map[string]string{
						"prometheus.io/scrape": "true",
					},
					Labels: map[string]string{
						"foo.bar": "custom-label",
					},
					SecurityContext: &corev1.PodSecurityContext{
						RunAsUser: pointer.Int64(1000),
					},
					Volumes: []corev1.Volume{
						{
							Name: "custom-libs",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
					},
				},
				Container: &egv1a1.KubernetesContainerSpec{
					Image: pointer.String("envoyproxy/envoy:v1.2.3"),
					Resources: &corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("400m"),
							corev1.ResourceMemory: resource.MustParse("2Gi"),
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:    resource.MustParse("200m"),
							corev1.ResourceMemory: resource.MustParse("1Gi"),
						},
					},
					SecurityContext: &corev1.SecurityContext{
						Privileged: pointer.Bool(true),
					},
					VolumeMounts: []corev1.VolumeMount{
						{
							Name:      "custom-libs",
							MountPath: "/lib/filter_foo.so",
						},
					},
				},
				InitContainers: []corev1.Container{
					{
						Name:    "install-filter-foo",
						Image:   "alpine:3.11.3",
						Command: []string{"/bin/sh", "-c"},
						Args:    []string{"echo \"Installing filter-foo\"; wget -q https://example.com/download/filter_foo_v1.0.0.tgz -O - | tar -xz --directory=/lib filter_foo.so; echo \"Done\";"},
						VolumeMounts: []corev1.VolumeMount{
							{
								Name:      "custom-libs",
								MountPath: "/lib",
							},
						},
					},
				},
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			kube := tc.infra.GetProxyInfra().GetProxyConfig().GetEnvoyProxyProvider().GetEnvoyProxyKubeProvider()
			if tc.deploy != nil {
				kube.EnvoyDeployment = tc.deploy
			}

			replace := egv1a1.BootstrapTypeReplace
			if tc.bootstrap != "" {
				tc.infra.Proxy.Config.Spec.Bootstrap = &egv1a1.ProxyBootstrap{
					Type:  &replace,
					Value: tc.bootstrap,
				}
			}

			if tc.telemetry != nil {
				tc.infra.Proxy.Config.Spec.Telemetry = tc.telemetry
			} else {
				tc.infra.Proxy.Config.Spec.Telemetry = &egv1a1.ProxyTelemetry{
					Metrics: &egv1a1.ProxyMetrics{
						Prometheus: &egv1a1.ProxyPrometheusProvider{
							Disable: true,
						},
					},
				}
			}

			if len(tc.proxyLogging) > 0 {
				tc.infra.Proxy.Config.Spec.Logging = egv1a1.ProxyLogging{
					Level: tc.proxyLogging,
				}
			}

			if tc.concurrency != nil {
				tc.infra.Proxy.Config.Spec.Concurrency = tc.concurrency
			}

			r := NewResourceRender(cfg.Namespace, tc.infra.GetProxyInfra())
			dp, err := r.Deployment()
			require.NoError(t, err)

			expected, err := loadDeployment(tc.caseName)
			require.NoError(t, err)

			sortEnv := func(env []corev1.EnvVar) {
				sort.Slice(env, func(i, j int) bool {
					return env[i].Name > env[j].Name
				})
			}

			sortEnv(dp.Spec.Template.Spec.Containers[0].Env)
			sortEnv(expected.Spec.Template.Spec.Containers[0].Env)
			assert.Equal(t, expected, dp)
		})
	}
}

func loadDeployment(caseName string) (*appsv1.Deployment, error) {
	deploymentYAML, err := os.ReadFile(fmt.Sprintf("testdata/deployments/%s.yaml", caseName))
	if err != nil {
		return nil, err
	}
	deployment := &appsv1.Deployment{}
	_ = yaml.Unmarshal(deploymentYAML, deployment)
	return deployment, nil
}

func TestService(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	svcType := egv1a1.ServiceTypeClusterIP
	cases := []struct {
		caseName string
		infra    *ir.Infra
		service  *egv1a1.KubernetesServiceSpec
	}{
		{
			caseName: "default",
			infra:    newTestInfra(),
			service:  nil,
		},
		{
			caseName: "custom",
			infra:    newTestInfra(),
			service: &egv1a1.KubernetesServiceSpec{
				Annotations: map[string]string{
					"key1": "value1",
				},
				Type: &svcType,
			},
		},
	}
	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			provider := tc.infra.GetProxyInfra().GetProxyConfig().GetEnvoyProxyProvider().GetEnvoyProxyKubeProvider()
			if tc.service != nil {
				provider.EnvoyService = tc.service
			}

			r := NewResourceRender(cfg.Namespace, tc.infra.GetProxyInfra())
			svc, err := r.Service()
			require.NoError(t, err)

			expected, err := loadService(tc.caseName)
			require.NoError(t, err)

			assert.Equal(t, expected, svc)
		})
	}
}

func loadService(caseName string) (*corev1.Service, error) {
	serviceYAML, err := os.ReadFile(fmt.Sprintf("testdata/services/%s.yaml", caseName))
	if err != nil {
		return nil, err
	}
	svc := &corev1.Service{}
	_ = yaml.Unmarshal(serviceYAML, svc)
	return svc, nil
}

func TestConfigMap(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	infra := newTestInfra()

	r := NewResourceRender(cfg.Namespace, infra.GetProxyInfra())
	cm, err := r.ConfigMap()
	require.NoError(t, err)

	expected, err := loadConfigmap()
	require.NoError(t, err)

	assert.Equal(t, expected, cm)
}

func loadConfigmap() (*corev1.ConfigMap, error) {
	cmYAML, err := os.ReadFile("testdata/configmap.yaml")
	if err != nil {
		return nil, err
	}
	cm := &corev1.ConfigMap{}
	_ = yaml.Unmarshal(cmYAML, cm)
	return cm, nil
}

func TestServiceAccount(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	infra := newTestInfra()

	r := NewResourceRender(cfg.Namespace, infra.GetProxyInfra())
	sa, err := r.ServiceAccount()
	require.NoError(t, err)

	expected, err := loadServiceAccount()
	require.NoError(t, err)

	assert.Equal(t, expected, sa)
}

func loadServiceAccount() (*corev1.ServiceAccount, error) {
	saYAML, err := os.ReadFile("testdata/serviceaccount.yaml")
	if err != nil {
		return nil, err
	}
	sa := &corev1.ServiceAccount{}
	_ = yaml.Unmarshal(saYAML, sa)
	return sa, nil
}

func TestHorizontalPodAutoscaler(t *testing.T) {
	cfg, err := config.New()
	require.NoError(t, err)

	cases := []struct {
		caseName string
		infra    *ir.Infra
		hpa      *egv1a1.KubernetesHorizontalPodAutoscalerSpec
	}{
		{
			caseName: "default",
			infra:    newTestInfra(),
			hpa: &egv1a1.KubernetesHorizontalPodAutoscalerSpec{
				MaxReplicas: ptr.To[int32](1),
			},
		},
		{
			caseName: "custom",
			infra:    newTestInfra(),
			hpa: &egv1a1.KubernetesHorizontalPodAutoscalerSpec{
				MinReplicas: ptr.To[int32](5),
				MaxReplicas: ptr.To[int32](10),
				Metrics: []autoscalingv2.MetricSpec{
					{
						Resource: &autoscalingv2.ResourceMetricSource{
							Name: corev1.ResourceCPU,
							Target: autoscalingv2.MetricTarget{
								Type:               autoscalingv2.UtilizationMetricType,
								AverageUtilization: ptr.To[int32](60),
							},
						},
						Type: autoscalingv2.ResourceMetricSourceType,
					},
					{
						Resource: &autoscalingv2.ResourceMetricSource{
							Name: corev1.ResourceMemory,
							Target: autoscalingv2.MetricTarget{
								Type:               autoscalingv2.UtilizationMetricType,
								AverageUtilization: ptr.To[int32](70),
							},
						},
						Type: autoscalingv2.ResourceMetricSourceType,
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.caseName, func(t *testing.T) {
			provider := tc.infra.GetProxyInfra().GetProxyConfig().GetEnvoyProxyProvider()
			provider.Kubernetes = egv1a1.DefaultEnvoyProxyKubeProvider()

			if tc.hpa != nil {
				provider.Kubernetes.EnvoyHpa = tc.hpa
			}

			provider.GetEnvoyProxyKubeProvider()

			r := NewResourceRender(cfg.Namespace, tc.infra.GetProxyInfra())
			hpa, err := r.HorizontalPodAutoscaler()
			require.NoError(t, err)

			want, err := loadHPA(tc.caseName)
			require.NoError(t, err)

			assert.Equal(t, want, hpa)
		})
	}
}

func loadHPA(caseName string) (*autoscalingv2.HorizontalPodAutoscaler, error) {
	hpaYAML, err := os.ReadFile(fmt.Sprintf("testdata/hpa/%s.yaml", caseName))
	if err != nil {
		return nil, err
	}

	hpa := &autoscalingv2.HorizontalPodAutoscaler{}
	_ = yaml.Unmarshal(hpaYAML, hpa)
	return hpa, nil
}
