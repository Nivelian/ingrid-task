package model

type Route struct {
	Destination string  `json:"destination"`
	Duration    float64 `json:"duration"`
	Distance    float64 `json:"distance"`
}

type RoutesInfo struct {
	Source string   `json:"source"`
	Routes []*Route `json:"routes"`
}

func (r RoutesInfo) Len() int {
	return len(r.Routes)
}

func (r RoutesInfo) Less(i, j int) bool {
	routeI, routeJ := r.Routes[i], r.Routes[j]
	if routeI.Duration == routeJ.Duration {
		return routeI.Distance < routeJ.Distance
	} else {
		return routeI.Duration < routeJ.Duration
	}
}

func (r RoutesInfo) Swap(i, j int) {
	r.Routes[i], r.Routes[j] = r.Routes[j], r.Routes[i]
}
