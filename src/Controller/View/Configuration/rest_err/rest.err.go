package rest_err

import "net/http"

//Pensar em todos os códigos HTTP que podem aparecer no cotidiano de um servidor.

// Assim declaro atributos de uma nova classe.
type RestErr struct {
	Message string   `j́son:"message"` //Metodo padrão de retorno em JSON
	Err     string   `j́son:"error"`
	Code    int      `j́son:"code"`
	Causes  []Causes `j́son:"causes"` // <-- []NomeObj -> Objeto :: Posso criar atributos que instanciam outros objetos
}

type Causes struct {
	Field   string `j́son:"field"`
	Message string `j́son:"message"`
}

// Primeiro método, declarado como função, passo um Ponteiro para o atributo RestErr, como forma de construtor do objeto, passando todos os atributos necessários
func NewRestErr(message, err string, code int64, causes []causes) *RestErr {
	return &RestErr{
		Message: message, //Message recebe messagem
		Err:     err,
		Code:    code,
		Causes:  Causes,
	}
}

// Método para lidar com Bad Requisitions
func NewBadReqError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadReqValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal_error",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}
