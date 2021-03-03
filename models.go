package hategoas

// SingleItemResponse returns a generic hateoas response.
type SingleItemResponse struct {
	Links          map[string]LinkType     `json:"_links,omitempty"`
	Embedded       map[string]EmbeddedType `json:"_embedded,omitempty"`
	SingleItemData `json:",inline"`
}

type SingleItemData interface{}

// CollectionResponse returns a generic hateoas response containing a collection.
type CollectionResponse struct {
	Collection []SingleItemResponse `json:",inline"`
}

// LinkType is a helpful link description, with additional metadata.
type LinkType struct {
	Href        string      `json:"href"`
	ID          interface{} `json:"id,omitempty"`
	Value       interface{} `json:"value,omitempty"`
	Deprecation string      `json:"deprecation,omitempty"`
	HrefLang    string      `json:"hreflang,omitempty"`
	Method      string      `json:"method,omitempty"`
	Name        string      `json:"name,omitempty"`
	Profile     string      `json:"profile,omitempty"`
	Rel         string      `json:"rel,omitempty"`
	Templated   bool        `json:"templated,omitempty"`
	Title       string      `json:"title,omitempty"`
	Type        string      `json:"type,omitempty"`
}

// EmbeddedType does some stuff.
type EmbeddedType struct {
	SingleItemResponse
	CollectionResponse
}
