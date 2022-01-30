package daos

type ServiceCreate struct {
	Name string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Host string `json:"host" validate:"required"`
	Port int64 `json:"port" validate:"required,number"`
	Protocol string `json:"protocol" validate:"required"`
	TimeoutMs int64 `json:"timeout" validate:"required,number"`
}

type ServiceUpdate struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Host string `json:"host"`
	Port int64 `json:"port"`
	Protocol string `json:"protocol"`
	TimeoutMs int64 `json:"timeout"`
}

type DeleteServiceResponse struct {
	Message string `json:"message"`
}