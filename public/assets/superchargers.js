var markers = [];
var map;


function urlParam(name) {
    var url = window.location.href;
    var results = new RegExp("[\\?&]" + name + "=([^&#]*)").exec(url);
    if (!results) {
        return undefined;
    }
    return decodeURIComponent(results[1]) || undefined;
}

function buildHTML(location) {
  if(location === undefined) {
    return "Unknown location";
  }

  body = "";
  for(var item in location) {
    body += "<li><strong>" + item + ":</strong> <code>" + location[item] + "</code></li>";
  }

  return "<ul>" + body + "</ul>";
}

function updateLocationCount(count) {
  var el = document.getElementById("location-count");
  var suffix = "location";
  if(count != 1) {
    suffix += "s";
  }
  el.innerText = count + " " + suffix;
}

function buildElementForLocation(location) {
  var el = document.createElement("div");
  el.className = "marker";
  if("locationType" in location) {
    if(location.locationType.includes("store")) {
      el.classList.add("store");
    } else if(location.locationType.includes("supercharger")) {
      el.classList.add("supercharger");
    } else if(location.locationType.includes("service")) {
      el.classList.add("service");
    } else if(location.locationType.includes("standard charger") || location.locationType.includes("destination charger")) {
      el.classList.add("charger");
    }
  }
  if("openSoon" in location && location.openSoon) {
    el.classList.add("coming-soon");
  }
  return el;
}

function query(event) {
  if(event) {
    event.preventDefault();
  }

  // Reset the example select menu
  this.example.selectedIndex = 0;

  Pace.start();
  fetch("/graphql", {
    method: "POST",
    body: this.query.value,
    headers: {
      "Accept": "application/json",
      "Content-Type": "application/graphql"
    }
  }).then(function(response) {
    if(response.ok) {
      return response.json().then(function(json) {
        if(json.errors !== undefined) {
          json.errors.forEach(function(error) {
            alert(error.message);
          });
          return;
        }
        // When we get a new result set, clear the markers
        markers.forEach(function(marker) {
          marker.remove();
        });

        var coordinates = json.data.locations.map(function(location) {
          return [location.longitude, location.latitude];
        });

        markers = [];
        json.data.locations.forEach(function(location) {
          var el = buildElementForLocation(location);
          var popup = new mapboxgl.Popup({
            offset: [0, -36]
          }).setHTML(buildHTML(location));

          var marker = new mapboxgl.Marker(el, {
            offset: [-12, -33]
          }).setLngLat([location.longitude, location.latitude])
          .setPopup(popup)
          .addTo(map);

          markers.push(marker);
        });

        var bounds;

        if(coordinates.length == 1) {
          bounds = new mapboxgl.LngLatBounds(coordinates[0], coordinates[0]);
        } else if(coordinates.length > 1) {
          bounds = coordinates.reduce(function(bounds, coord) {
            return bounds.extend(coord);
          }, new mapboxgl.LngLatBounds(coordinates[0], coordinates[0]));
        }

        map.fitBounds(bounds, {
          padding: 100,
          linear: false,
          maxZoom: 3
        });

        updateLocationCount(markers.length);
        Pace.stop();
      });
    } else {
      Pace.stop();
      alert("Network request could not complete");
    }
  })
  .catch(function(error) {
    alert("There was a problem with your request");
    Pace.stop();
  });
}

Pace.on("start", function() {
  document.getElementById("header").classList.add("loading");
});
Pace.on("hide", function() {
  document.getElementById("header").classList.remove("loading");
});

window.onload = function() {
  map = new mapboxgl.Map({
      container: "map",
      style: "mapbox://styles/mapbox/streets-v9",
      zoom: 3,

      // Center of the US
      center: [-98.585522, 39.8333333]
  });
  var form = document.getElementsByTagName("form")[0];
  var queryParam = urlParam("query");
  if(queryParam !== undefined) {
    form.query.value = queryParam;
  }

  form.addEventListener("submit", query);

  form.query.addEventListener("keydown", function(e) {
  	if(e.keyCode == 13 && e.metaKey) {
  		query.apply(this.form);
  	}
  });

  form.example.addEventListener("change", function(e) {
    var index = this.options.selectedIndex;
    if(index) {
      form.query.value = decodeURIComponent(this.options[index].value);
    }
  });

  document.querySelector("form .toggle").addEventListener("click", function(e) {
    e.preventDefault();
    form.classList.toggle("collapsed");
  });

  // Fetch initial placeholder query data
  query.apply(form);

  window.addEventListener("keyup", function(e) {
    // s and / key
  	if(e.keyCode == 83 || e.keyCode == 191) {
  		form.query.focus();
  	}
  });
};
