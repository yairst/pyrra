package slo

import (
	"testing"
	"time"

	"github.com/prometheus/common/model"
	"github.com/stretchr/testify/require"
)

func TestObjective_QueryTotal(t *testing.T) {
	testcases := []struct {
		name      string
		objective Objective
		expected  string
	}{{
		name: "http",
		objective: Objective{
			Window: model.Duration(24 * time.Hour),
			Indicator: Indicator{
				HTTP: &HTTPIndicator{},
			},
		},
		expected: `sum(increase(http_requests_total{}[1d]))`,
	}, {
		name: "http-custom",
		objective: Objective{
			Window: model.Duration(24 * time.Hour),
			Indicator: Indicator{HTTP: &HTTPIndicator{
				Metric:         "prometheus_http_requests_total",
				Selectors:      Selectors{`foo="bar"`},
				ErrorSelectors: Selectors{`status=~"5.."`},
			}},
		},
		expected: `sum(increase(prometheus_http_requests_total{foo="bar"}[1d]))`,
	}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.objective.QueryTotal())
		})
	}
}

func TestObjective_QueryErrors(t *testing.T) {
	testcases := []struct {
		name      string
		objective Objective
		expected  string
	}{{
		name: "http",
		objective: Objective{
			Window: model.Duration(28 * 24 * time.Hour),
			Indicator: Indicator{
				HTTP: &HTTPIndicator{},
			},
		},
		expected: `sum(increase(http_requests_total{code=~"5.."}[4w]))`,
	}, {
		name: "http-custom",
		objective: Objective{
			Window: model.Duration(28 * 24 * time.Hour),
			Indicator: Indicator{
				HTTP: &HTTPIndicator{
					Metric:         "prometheus_http_requests_total",
					Selectors:      Selectors{`job="prometheus"`},
					ErrorSelectors: Selectors{`status=~"5.."`},
				},
			},
		},
		expected: `sum(increase(prometheus_http_requests_total{job="prometheus",status=~"5.."}[4w]))`,
	}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.objective.QueryErrors())
		})
	}
}

func TestObjective_QueryErrorBudget(t *testing.T) {
	testcases := []struct {
		name      string
		objective Objective
		expected  string
	}{{
		name: "http",
		objective: Objective{
			Window: model.Duration(28 * 24 * time.Hour),
			Target: 99.9,
			Indicator: Indicator{
				HTTP: &HTTPIndicator{},
			},
		},
		expected: `(((100 - 99.900)/100) - (sum(increase(http_requests_total{code=~"5.."}[4w])) / sum(increase(http_requests_total{}[4w])))) / ((100 - 99.900)/100)`,
	}, {
		name: "http-custom",
		objective: Objective{
			Window: model.Duration(28 * 24 * time.Hour),
			Target: 95.3,
			Indicator: Indicator{
				HTTP: &HTTPIndicator{
					Metric:         "prometheus_http_request_total",
					Selectors:      Selectors{`job="prometheus"`},
					ErrorSelectors: Selectors{`status=~"5.."`},
				},
			},
		},
		expected: `(((100 - 95.300)/100) - (sum(increase(prometheus_http_request_total{job="prometheus",status=~"5.."}[4w])) / sum(increase(prometheus_http_request_total{job="prometheus"}[4w])))) / ((100 - 95.300)/100)`,
	}, {
		name: "grpc",
		objective: Objective{
			Window: model.Duration(28 * 24 * time.Hour),
			Target: 99.9,
			Indicator: Indicator{
				GRPC: &GRPCIndicator{
					Service: "service",
					Method:  "method",
				},
			},
		},
		expected: `(((100 - 99.900)/100) - (sum(increase(grpc_server_handled_total{grpc_service="service",grpc_method="method",grpc_code=~"Aborted|Unavailable|Internal|Unknown|Unimplemented|DataLoss"}[4w])) / sum(increase(grpc_server_handled_total{grpc_service="service",grpc_method="method"}[4w])))) / ((100 - 99.900)/100)`,
	}, {
		name: "grpc-custom",
		objective: Objective{
			Window: model.Duration(14 * 24 * time.Hour),
			Target: 95.3,
			Indicator: Indicator{
				GRPC: &GRPCIndicator{
					Service:   "awesome",
					Method:    "lightspeed",
					Selectors: Selectors{`job="app"`},
				},
			},
		},
		expected: `(((100 - 95.300)/100) - (sum(increase(grpc_server_handled_total{grpc_service="awesome",grpc_method="lightspeed",job="app",grpc_code=~"Aborted|Unavailable|Internal|Unknown|Unimplemented|DataLoss"}[2w])) / sum(increase(grpc_server_handled_total{grpc_service="awesome",grpc_method="lightspeed",job="app"}[2w])))) / ((100 - 95.300)/100)`,
	}}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.expected, tc.objective.QueryErrorBudget())
		})
	}
}
