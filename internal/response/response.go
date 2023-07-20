package response

type PaginatedResponse struct {
    Page       int         `json:"page"`
    Result     interface{} `json:"result"`
    TotalPage  int         `json:"total_page"`
    TotalItems int         `json:"total_result"`
}
