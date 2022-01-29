package models

type Resume struct {
	FullName     string        `json:"fullname,omitempty"`
	Email        string        `json:"email,omitempty"`
	Website      string        `json:"website,omitempty"`
	Summary      string        `json:"summary,omitempty"`
	Technologies []string      `json:"technologies,omitempty"`
	Education    []Education   `json:"education,omitempty"`
	WorkHistory  []WorkHistory `json:"work history,omitempty"`
	Languages    []string      `json:"languages,omitempty"`
}

type Education struct {
	Institution string   `json:"institution,omitempty"`
	Programme   string   `json:"programme,omitempty"`
	Country     string   `json:"country,omitempty"`
	City        string   `json:"city,omitempty"`
	Startyear   string   `json:"startyear,omitempty"`
	Endyear     string   `json:"endyear,omitempty"`
	Grade       string   `json:"grade,omitempty"`
	Noteworthy  []string `json:"noteworthy,omitempty"`
}

type WorkHistory struct {
	Employer         string   `json:"employer,omitempty"`
	Role             string   `json:"role,omitempty"`
	Country          string   `json:"country,omitempty"`
	City             string   `json:"city,omitempty"`
	Start            Start    `json:"start,omitempty"`
	End              End      `json:"end,omitempty"`
	Responsibilities []string `json:"responsibilities,omitempty"`
}

type Start struct {
	Month string `json:"month,omitempty"`
	Year  string `json:"year,omitempty"`
}

type End struct {
	Month string `json:"month,omitempty"`
	Year  string `json:"year,omitempty"`
}
