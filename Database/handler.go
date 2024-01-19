package sqlDB

import (
	"fmt"
	"net/http"
	"strings"
)




func DBHandlerfunc( db *Database, fallback http.Handler) (http.HandlerFunc) {


	return func(w http.ResponseWriter, r *http.Request) {
	
		url := strings.Split(r.URL.Path, "/")
		if len(url) < 2 {
			fallback.ServeHTTP(w, r)
			return
		}
		token := url[len(url)- 1]
		data, err := db.FetchOne("url_shortner", "TOKEN", token)
		if err != nil {

			fmt.Fprintf(w, "Not Found")
			return 
		}
		http.Redirect(w, r, data.Url, http.StatusMovedPermanently)
	}

}