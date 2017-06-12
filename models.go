package main

type (
	Geolocation struct {
		Lat  string
		Lon  string
		Name string
	}

	Response struct {
		Status  string   `json:"status"`
		Results []Result `json:"results"`
	}

	Result struct {
		Types             []string           `json:"types"`
		FormattedAddress  string             `json:"formatted_address"`
		AddressComponents []AddressComponent `json:"address_components"`
		Geometry          GeometryData       `json:"geometry"`
	}

	AddressComponent struct {
		LongName  string   `json:"long_name"`
		ShortName string   `json:"short_name"`
		Types     []string `json:"types"`
	}

	GeometryData struct {
		Location     LatLng `json:"location"`
		LocationType string `json:"location_type"`
		Viewport     struct {
			Southwest LatLng `json:"southwest"`
			Northeast LatLng `json:"northeast"`
		} `json:"viewport"`
		Bounds struct {
			Southwest LatLng `json:"southwest"`
			Northeast LatLng `json:"northeast"`
		} `json:"bounds"`
	}

	LatLng struct {
		Lat float64 `json:"lat"`
		Lng float64 `json:"lng"`
	}
)
