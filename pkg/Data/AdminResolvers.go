package Data

import (
	"github.com/graphql-go/graphql"
	"log"
)

var listAdminsResolver = func(p graphql.ResolveParams) (interface{}, error) {
	var admins []Admin
	rows, err := db.Query("SELECT id, username, email FROM classmanager.admins")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id       int64
			username string
			email    string
		)
		if err := rows.Scan(&id, &username, &email); err != nil {
			log.Println(err)
		}
		admins = append(admins, Admin{
			ID:       id,
			Username: username,
			Email:    email,
		})
	}
	return admins, nil
}