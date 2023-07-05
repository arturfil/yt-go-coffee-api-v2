package services

type JsonResponse struct {
    Error bool `json:"error"`
    Message string `json:"message"`
    Data interface{} `json:"data,omitresponse"`
}
