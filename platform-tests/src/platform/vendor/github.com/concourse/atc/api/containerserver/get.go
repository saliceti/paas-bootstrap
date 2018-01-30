package containerserver

import (
	"encoding/json"
	"net/http"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/api/present"
	"github.com/concourse/atc/db"
)

func (s *Server) GetContainer(teamDB db.TeamDB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handle := r.FormValue(":id")

		hLog := s.logger.Session("container", lager.Data{
			"handle": handle,
		})

		container, found, err := teamDB.GetContainer(handle)
		if err != nil {
			hLog.Error("failed-to-lookup-container", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if !found {
			hLog.Debug("container-not-found")
			w.WriteHeader(http.StatusNotFound)
			return
		}

		hLog.Debug("found-container")

		presentedContainer := present.Container(container)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(presentedContainer)
	})
}