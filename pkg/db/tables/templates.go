package tables

// ** Global **
type Templates struct {
	Version   string        `json:"version"`
	Templates []interface{} `json:"templates"`
}

type TemplatesArray struct {
	Container []Container
	Compose   []Compose
	Stack     []Stack
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
