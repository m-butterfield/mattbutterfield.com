mapboxgl.accessToken = "pk.eyJ1IjoibWJ1dHRlcmZpZWxkIiwiYSI6ImNsdDllbDFkYjA3dGUycXBqMXkydjd1aWEifQ.jvPe-lNUJFl4x74IYiRZpA";

const map = new mapboxgl.Map({
  container: "heatmap",
  style: "mapbox://styles/mbutterfield/clt9fms1l003l01qqfjccc3i3",
  zoom: 10.74,
  center: [-73.95551, 40.73932]
});

map.on("load", () => {
  map.addSource("heatmap", {
    type: "vector",
    url: "mapbox://mbutterfield.heatmap"
  });
  map.addLayer({
    "id": "heatmap",
    "type": "line",
    "source": "heatmap",
    "source-layer": "heatmap",
    "paint": {
      "line-color": "#EB9360",
      "line-width": 1,
      "line-opacity": 0.8
    },
  });
  map.addControl(new mapboxgl.NavigationControl());
});
