package prioritas

type Prioritas struct {
	PrioritasId string `json:"prioritas_id" gorm:"primaryKey"`
	Name        string `json:"name"`
}
