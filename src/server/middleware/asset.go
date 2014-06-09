package middleware

import (
	"github.com/adminibar/boatyard/res/bin"
	"log"
	"net/http"
)

/**
 * Asset Middle ware that attempts to
 * return embedded web resources
 */
type AssetM struct {
	next http.Handler
}

//interface implementation for middleware
func (a *AssetM) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.RequestURI()[1:]
	if path == "" {
		a.next.ServeHTTP(w, r)
	} else {
		log.Printf("Requesting asset '%s'", path)
		err := WriteHTTPAsset(w, path)
		if err != nil {
			a.next.ServeHTTP(w, r)
			return
		}
	}
}

//create the middleware
func NewAssetM(handler http.Handler) *AssetM {
	return &AssetM{next: handler}
}

func WriteHTTPAsset(w http.ResponseWriter, name string) error {
	data, err := res.Asset(name)
	if err != nil {
		return err
	}

	w.Write(data)
	return nil
}
