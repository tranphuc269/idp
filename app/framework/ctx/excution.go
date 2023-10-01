package ctx

import (
	"idp_system/app/framework/geo"
	"time"
)

type ExecutionContext struct {
	RequestID        string
	RequestStartAt   time.Time
	Location         geo.Location
	FeatureToggleID  string
	ExperimentBucket string
}
