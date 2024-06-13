package reqs

type SetCategoryRequest struct {
	RepositoryID int64 `json:"repository_id"`
	CategoryName   string `json:"category_name"`
}
