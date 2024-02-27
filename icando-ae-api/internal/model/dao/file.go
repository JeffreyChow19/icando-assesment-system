package dao

type PresignedUrlDao struct {
	Url       string `json:"url"`
	ObjectKey string `json:"objectKey"`
}
