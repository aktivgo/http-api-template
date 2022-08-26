package user

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"wire-di/internal/domain"
	"wire-di/pkg/logger"
)

type handler struct {
	logger logger.Logger
	svc    domain.UserService
}

func (h *handler) FetchByUsername() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			h.logger.Error(err.Error(), nil)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"body read failed: %s\"}", err)))
			return
		}

		var params struct {
			Username string
		}
		if err = json.Unmarshal(body, &params); err != nil {
			h.logger.Error(err.Error(), nil)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"params unmarshall failed: %s\"}", err)))
			return
		}

		user, err := h.svc.FetchByUsername(ctx, params.Username)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"fetch by username failed: %s\"}", err)))
			return
		}

		response, err := json.Marshal(user)
		if err != nil {
			h.logger.Error(err.Error(), nil)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("{\"error\":\"response marshall failed: %s\"}", err)))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
