package response

type KeyTemplates struct {
	Templates []KeyTemplatesWithType `json:"templates"`
}

type KeyTemplatesWithType struct {
	Type      string        `json:"type"`
	Templates []KeyTemplate `json:"templates"`
}

type KeyTemplate struct {
	ID           string `json:"id"`
	Name         string `json:"templateName"`
	Description  string `json:"templateDescription"`
	Color        string `json:"color"`
	Tags         string `json:"tags"`
	PreviewUrl   string `json:"previewUrl"`
	LastModified string `json:"lastModified"`
	Archived     bool   `json:"archived"`
}

type BlockTemplates struct {
	Templates []BlockTemplate `json:"templates"`
}

type BlockTemplate struct {
	ID            string  `json:"id"`
	Name          string  `json:"templateName"`
	Description   string  `json:"templateDescription"`
	Color         string  `json:"color"`
	Tags          string  `json:"tags"`
	PreviewUrl    string  `json:"previewUrl"`
	LastModified  string  `json:"lastModified"`
	Archived      bool    `json:"archived"`
	KeyTemplateId *string `json:"keyTemplateId"`
}

type PodTemplates struct {
	Templates []PodTemplate `json:"templates"`
}

type PodTemplate struct {
	ID              string  `json:"id"`
	Name            string  `json:"templateName"`
	Description     string  `json:"templateDescription"`
	Color           string  `json:"color"`
	Tags            string  `json:"tags"`
	PreviewUrl      string  `json:"previewUrl"`
	LastModified    string  `json:"lastModified"`
	Archived        bool    `json:"archived"`
	KeyTemplateId   *string `json:"keyTemplateId"`
	BlockTemplateId string  `json:"blockTemplateId"`
}
