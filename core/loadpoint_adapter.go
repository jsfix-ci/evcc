package core

import (
	"github.com/evcc-io/evcc/core/planner"
	"github.com/evcc-io/evcc/core/soc"
)

var _ planner.Adapter = (*adapter)(nil)

type adapter struct {
	*LoadPoint
}

func (a *adapter) Publish(key string, val interface{}) {
	a.LoadPoint.publish(key, val)
}

func (a *adapter) SocEstimator() *soc.Estimator {
	return a.LoadPoint.socEstimator
}

func (a *adapter) TargetSoC() int {
	return a.LoadPoint.targetSoC
}

func (a *adapter) TargetTime() int {
	return a.LoadPoint.targetTime
}
