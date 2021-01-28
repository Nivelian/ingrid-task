package model

type OsrmRoute struct {
	Duration float64
	Distance float64
}

type OsrmRoutes struct {
	Routes []*OsrmRoute
	Code   string
}
