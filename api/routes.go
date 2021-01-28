package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Nivelian/ingrid-task/helpers"
	"github.com/Nivelian/ingrid-task/model"
	"net/http"
	"sort"
)

const routesUrl = "http://router.project-osrm.org/route/v1/driving/%v;%v?overview=false"

func getRoute(src, dst string) (*model.Route, error) {
	resp, err := http.Get(fmt.Sprintf(routesUrl, src, dst))
	if err != nil {
		msg := "Error getting info from osrm"
		if resp != nil {
			msg += fmt.Sprintf(". Status code = %v", resp.StatusCode)
		}
		return nil, helpers.LogErr(err, msg)
	}
	defer resp.Body.Close()

	osrmRoutes := new(model.OsrmRoutes)
	if err := json.NewDecoder(resp.Body).Decode(osrmRoutes); err != nil {
		return nil, helpers.LogErr(err, "Failed to decode osrm routes response")
	}

	if code := osrmRoutes.Code; code != "Ok" {
		return nil, helpers.LogErr(errors.New("bad osrm response"), "Osrm responded with status %v", code)
	}

	return &model.Route{
		Destination: dst,
		Distance:    osrmRoutes.Routes[0].Distance,
		Duration:    osrmRoutes.Routes[0].Duration,
	}, nil
}

func GetRoutes(src string, dsts []string) (*model.RoutesInfo, error) {
	result := &model.RoutesInfo{Source: src}
	var routes []*model.Route

	for _, dst := range dsts {
		route, err := getRoute(src, dst)
		if err != nil {
			return nil, err
		}

		routes = append(routes, route)
	}

	result.Routes = routes

	sort.Sort(result)

	return result, nil
}
