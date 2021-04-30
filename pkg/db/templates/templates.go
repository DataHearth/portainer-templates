package templates

// ** Global **
type Templates struct {
	Version   string        `json:"version"`
	Templates []interface{} `json:"templates"`
}

type Compose struct {
	ID                int        `json:"id,omitempty"`
	Type              int        `json:"type"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	Note              string     `json:"note,omitempty"`
	Categories        []string   `json:"categories,omitempty"`
	Platform          string     `json:"platform,omitempty"`
	Logo              string     `json:"logo,omitempty"`
	Repository        Repository `json:"repository"`
	Env               []Env      `json:"env,omitempty"`
	AdministratorOnly bool       `json:"administrator_only,omitempty"`
	Name              string     `json:"name,omitempty"`
}

type Stack struct {
	ID                int        `json:"id,omitempty"`
	Type              int        `json:"type"`
	Title             string     `json:"title"`
	Description       string     `json:"description"`
	Note              string     `json:"note,omitempty"`
	Categories        []string   `json:"categories,omitempty"`
	Platform          string     `json:"platform,omitempty"`
	Logo              string     `json:"logo,omitempty"`
	Repository        Repository `json:"repository"`
	Env               []Env      `json:"env,omitempty"`
	AdministratorOnly bool       `json:"administrator_only,omitempty"`
	Name              string     `json:"name,omitempty"`
}

type Container struct {
	ID                int       `json:"id,omitempty"`
	Type              int       `json:"type"`
	Title             string    `json:"title"`
	Description       string    `json:"description"`
	Categories        []string  `json:"categories,omitempty"`
	Platform          string    `json:"platform,omitempty"`
	Logo              string    `json:"logo,omitempty"`
	Image             string    `json:"image"`
	Ports             []string  `json:"ports,omitempty"`
	Volumes           []Volumes `json:"volumes"`
	AdministratorOnly bool      `json:"administrator_only,omitempty"`
	Name              string    `json:"name,omitempty"`
	Registry          string    `json:"registry,omitempty"`
	Command           string    `json:"command,omitempty"`
	Env               []Env     `json:"env,omitempty"`
	Network           string    `json:"network,omitempty"`
	Labels            []Label   `json:"labels,omitempty"`
	Privileged        bool      `json:"privileged,omitempty"`
	Interactive       bool      `json:"interactive,omitempty"`
	RestartPolicy     string    `json:"restart_policy,omitempty"`
	Hostname          string    `json:"hostname,omitempty"`
	Note              string    `json:"note,omitempty"`
}

type Volumes struct {
	ID        int    `json:"id"`
	Container string `json:"container"`
	Bind      string `json:"bind,omitempty"`
	ReadOnly  bool   `json:"readonly,omitempty"`
}

type Env struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Label       string   `json:"label"`
	Description string   `json:"description,omitempty"`
	Default     string   `json:"default,omitempty"`
	Preset      string   `json:"preset,omitempty"`
	Select      []Select `json:"select,omitempty"`
}

type Select struct {
	ID      int    `json:"id"`
	Text    string `json:"text"`
	Value   string `json:"value"`
	Default bool   `json:"default,omitempty"`
}

type Repository struct {
	ID        int    `json:"id"`
	URL       string `json:"url"`
	Stackfile string `json:"stackfile"`
}

type Label struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
