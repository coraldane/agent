package funcs

import (
	"gitlab.tarzip.com/open-falcon/common/model"
)

func AgentMetrics() []*model.MetricValue {
	return []*model.MetricValue{GaugeValue("agent.alive", 1)}
}
