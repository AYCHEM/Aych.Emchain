package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mailchain/mailchain/errs"
	"github.com/mailchain/mailchain/internal/address"
	"github.com/mailchain/mailchain/nameservice"
	"github.com/pkg/errors"
)

// Forward handle forward domain lookups where an address is looked up to find a domain name.
func Forward(resolver nameservice.ForwardLookup, protocol, network string) func(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Address string `json:"address"`

		// The rFC1035 error status, if present
		// Since 0 status belongs to 'No Error', it's safe to use 'omitempty'
		//
		// Required: false
		// example: 3
		Status int `json:"status,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query()["domain-name"]) != 1 {
			errs.JSONWriter(w, http.StatusPreconditionFailed, errors.Errorf("domain-name must be specified exactly once"))
			return
		}

		resolvedAddress, err := resolver.ResolveName(r.Context(), protocol, network, r.URL.Query()["domain-name"][0])
		if nameservice.IsRFC1035Error(err) {
			_ = json.NewEncoder(w).Encode(response{
				Status: nameservice.RFC1035StatusMap[err],
			})

			return
		}

		if err != nil {
			errs.JSONWriter(w, http.StatusInternalServerError, err)
			return
		}

		encAddress, _, err := address.EncodeByProtocol(resolvedAddress, protocol)
		if err != nil {
			errs.JSONWriter(w, http.StatusInternalServerError, errors.WithMessage(err, "failed to encode address"))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response{
			Address: encAddress,
		})
	}
}
