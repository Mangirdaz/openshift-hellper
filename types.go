package main

//reflection is not full, just what we need. Feel free to extend
type ImageStream struct {
	APIVersion string   `json:"apiVersion,omitempty`
	Kind       string   `json:"kind,omitempty"`
	Metadata   struct{} `json:"metadata,omitempty"`
	Items      []Items  `json:"items,omitempty"`
}

type Items struct {
	APIVersion    string        `json:"apiVersion,omitempty"`
	Kind          string        `json:"kind,omitempty"`
	ItemsMetadata ItemsMetadata `json:"metadata,omitempty"`
	Spec          Spec          `json:"spec,omitempty"`
}

type ItemsMetadata struct {
	CreationTimestamp string            `json:"creationTimestamp,omitempty"`
	Name              string            `json:"name,omitempty"`
	Annotations       map[string]string `json:"annotations,omitempty"`
}

type Spec struct {
	Tags []Tags `json:"tags,omitempty"`
}

type Tags struct {
	Name         string       `json:"name,omitempty"`
	Annotations  Annotations  `json:"annotations,omitempty"`
	From         From         `json:"from,omitempty"`
	ImportPolicy ImportPolicy `json:"importPolicy,omitempty"`
}

type Annotations struct {
	Description string `json:"description,omitempty"`
	IconClass   string `json:"iconClass,omitempty"`
	Tags        string `json:"tags,omitempty"`
}

type From struct {
	Kind string `json:"kind,omitempty"`
	Name string `json:"name,omitempty"`
}

type ImportPolicy struct {
	Insecure bool `json:"insecure,omitempty"`
}
