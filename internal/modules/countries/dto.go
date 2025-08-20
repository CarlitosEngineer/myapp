package countries

type CreateCountryDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
type UpdateCountryDTO struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
