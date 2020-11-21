package persona

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

/*getPersonByIdRequest estructura para recueperar datos del request*/
type getPersonByIDRequest struct {
	PersonaID string
}

func makeGetPersonByIDEndPoint() endpoint.Endpoint {
	getPersonByID := func(ctx context.Context, request interface{}) (interface{}, error) {
		// rep := request.(getPersonByIDRequest)
		//_, err := s.GetPersonByID(&rep)

		return map[string]string{"mensaje": "cscscsc"}, nil
	}
	return getPersonByID
}
