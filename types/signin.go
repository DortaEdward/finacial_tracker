package types

type SigninRequest struct{
  Email string `json:"email" bind:"required"`
  Password string `json:"password" bind:"required"`
}
