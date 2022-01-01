package daos

type ServiceCreate struct {
	Name string `json:"name" validate:"required"`
	Host string `json:"host" validate:"required"`
	Port int64 `json:"port" validate:"required,number"`
	Protocol string `json:"protocol" validate:"required"`
	TimeoutMs int64 `json:"timeout" validate:"required,number"`
	RefreshTimeMs int64 `json:"refresh_time" validate:"number"`
}

type ServiceUpdate struct {
	ID uint `json:"id" validate:"required"`
	Attribute string `json:"attribute" validate:"required"`
	NewValue interface{} `json:"new_value" validate:"required"`
}