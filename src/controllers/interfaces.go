package controllers

type renderer interface {
	Render(id string, data []byte) ([]byte, error)
}
