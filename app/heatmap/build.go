package heatmap

import (
	"github.com/m-butterfield/mattbutterfield.com/app/data"
	"log"
	"os/exec"
)

func buildHeatmap(ds data.Store) (string, error) {
	coordinates, err := getActivityCoordinates(ds)
	if err != nil {
		return "", err
	}

	log.Print("Saving heatmap geoJSON")
	geoJSONFileName := "./tmp/heatmap.geojson"
	if err = saveGeoJSONResultLocal(coordinates, geoJSONFileName); err != nil {
		return "", err
	}

	// I don't think we need this step...
	//log.Print("Converting geoJSON to EPSG:4326")
	//convertedGeoJSONFileName := "./tmp/heatmap_4326.geojson"
	//cmd := exec.Command("ogr2ogr", "-f", "GeoJSON", convertedGeoJSONFileName, "-t_srs", "EPSG:4326", geoJSONFileName)
	//if err := cmd.Run(); err != nil {
	//	return "", err
	//}

	log.Print("Converting geoJSON to MBTiles")
	mbtilesFileName := "./tmp/heatmap.mbtiles"
	cmd := exec.Command("tippecanoe", "--force", "-o", mbtilesFileName, geoJSONFileName)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	return mbtilesFileName, nil
}
