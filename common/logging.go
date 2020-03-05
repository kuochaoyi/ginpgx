package common

type CreatedLog struct {
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
}

type UpdatedLog struct {
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type DeletedLog struct {
	DeletedAt string `json:"deleted_at"`
	DeletedBy string `json:"deleted_by"`
}
