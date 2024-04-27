package heatmap

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
)

func UpdateHeatMap() error {
	ds, err := data.Connect()
	if err != nil {
		return err
	}

	if err := getLatestActivities(ds); err != nil {
		return err
	}

	mbtilesFileName, err := buildHeatmap(ds)
	if err != nil {
		return err
	}

	err = updateMapbox(ds, mbtilesFileName)
	if err != nil {
		return err
	}

	return nil
}
