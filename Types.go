package main

type ReportType struct {
	player_id int
	fleets    map[string]FleetType
	players   map[string]PlayerType
	stars     map[string]StarType
}

type NeptuneResponse struct {
	event string
	order string
	error string
}

type FleetType struct {
	ships          int
	owner_id       int
	destination_id int
	name           string
}

type PlayerType struct {
	id       int
	name     string
	economy  int
	industry int
	science  int
	stars    int
}

type StarType struct {
	id       int
	name     string
	economy  int
	industry int
	science  int
}
