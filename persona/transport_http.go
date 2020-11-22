package persona

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	kithttp "github.com/go-kit/kit/transport/http"
)

/*MakeHTTPHandler nos permitira ejecutar metodos de persona*/
func MakeHTTPHandler() http.Handler {
	r := chi.NewRouter()

	//Obtener personas por su identificador
	getPersonByHandler := kithttp.NewServer(
		makeGetPersonByIDEndPoint(),
		getPersonByIDRequestDecoder,
		kithttp.EncodeJSONResponse,
	)
	r.Method(http.MethodGet, "/mensaje", getPersonByHandler)

	return r
}

func getPersonByIDRequestDecoder(context context.Context, r *http.Request) (interface{}, error) {
	return getPersonByIDRequest{
		PersonaID: chi.URLParam(r, "id"),
	}, nil
}
