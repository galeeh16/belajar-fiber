package request

type CreateContactRequest struct {
	Name      string `json:"name" validate:"required,min=3,max=100,unique_contact_name"`
	Email     string `json:"email,omitempty"`   // nullable -> omitempty
	Address   string `json:"address,omitempty"` // nullable -> omitempty
	Handphone string `json:"no_hp,omitempty"`   // nullable -> omitempty
}
