package hello

import (
	"fmt"
	"net/http"

	"MadridMas/server/incident"

	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
)

const (
	// TODO(sara): Replace the username and password for environment variables.
	// Temporary allow anyone see this.
	dbCONNECTIONNAME = "madridmas-172613:europe-west1:madridmassql"
	dbUSER           = "madridmas"
	dbPASSWORD       = "madridmas"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	db, err := mysql.DialPassword(dbCONNECTIONNAME, dbUSER, dbPASSWORD)
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not open db: %v", err), 500)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id,description FROM incidents")
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not query db: %v", err), 500)
		return
	}
	defer rows.Close()
	var i incident.Incident
	for rows.Next() {
		if err := rows.Scan(&i.Id, &i.Description); err != nil {
			http.Error(w, fmt.Sprintf("failed to get filed: %v", err), 500)
		}
		fmt.Fprintf(w, "<p><b>%d,</b>", i.Id)
		fmt.Fprintf(w, "%s</p>", i.Description)
	}
}
