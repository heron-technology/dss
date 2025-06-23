package models

import (
	"testing"
	"time"

	dssmodels "github.com/interuss/dss/pkg/models"
)

func TestNewImplicitSubscription(t *testing.T) {
	manager := dssmodels.ManagerFromString("")
	extents := []*dssmodels.Volume4D{{
		SpatialVolume: &dssmodels.Volume3D{
			AltitudeHi: nil,
			AltitudeLo: nil,
			Footprint: &dssmodels.GeoCircle{
				Center: dssmodels.LatLngPoint{
					Lng: 0.0,
					Lat: 0.0,
				},
				RadiusMeter: 5.0,
			},
		},
		EndTime:   &time.Time{},
		StartTime: &time.Time{},
	}}
	uExtent, _ := dssmodels.UnionVolumes4D(extents...)
	cells, _ := uExtent.CalculateSpatialCovering()
	shouldNotifyForConstraints := true
	httpOnly := true
	ussBaseURL := "https://123123.com"

	sut, err := NewImplicitSubscription(manager, extents[0], cells, ussBaseURL, &shouldNotifyForConstraints, httpOnly)

	if err != nil {
		t.Errorf("failed %v", err)
	}

	if sut.NotifyForConstraints != true {
		t.Errorf("NotifyForConstraints should be true")
	}
}
