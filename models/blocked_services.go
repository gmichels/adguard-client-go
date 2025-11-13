package models

// BlockedServicesAll
type BlockedServicesAll struct {
	BlockedServices []BlockedService `json:"blocked_services"`
	Groups          []ServiceGroup   `json:"groups,omitempty"`
}

// BlockedService
type BlockedService struct {
	IconSvg string   `json:"icon_svg"`
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Rules   []string `json:"rules"`
	GroupId string   `json:"group_id,omitempty"`
}

// BlockedServicesSchedule
type BlockedServicesSchedule struct {
	Schedule Schedule `json:"schedule,omitempty"`
	Ids      []string `json:"ids"`
}

// Schedule - Sets periods of inactivity for filtering blocked services. The schedule contains 7 days (Sunday to Saturday) and a time zone.
type Schedule struct {
	TimeZone  string   `json:"time_zone,omitempty" description:"Time zone name according to IANA time zone database. For example 'Europe/Brussels'. 'Local' represents the system's local time zone."`
	Sunday    DayRange `json:"sun,omitempty"`
	Monday    DayRange `json:"mon,omitempty"`
	Tuesday   DayRange `json:"tue,omitempty"`
	Wednesday DayRange `json:"wed,omitempty"`
	Thursday  DayRange `json:"thu,omitempty"`
	Friday    DayRange `json:"fri,omitempty"`
	Saturday  DayRange `json:"sat,omitempty"`
}

// DayRange - The single interval within a day. It begins at the `start` and ends before the `end`.
type DayRange struct {
	Start uint `json:"start,omitempty" description:"The number of milliseconds elapsed from the start of a day. It must be less than 'end' and is expected to be rounded to minutes. So the maximum value is '86340000' (23 hours and 59 minutes)."`
	End   uint `json:"end,omitempty" description:"The number of milliseconds elapsed from the start of a day. It is expected to be rounded to minutes. The maximum value is 86400000 (24 hours)."`
}

// ServiceGroup represents a group of blocked services
type ServiceGroup struct {
	Id string `json:"id"`
}
