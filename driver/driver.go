package driver

import (
	"database/sql"
	"github.com/lib/pq"
	"log"
)

var db *sql.DB

func ConnectDB() *sql.DB {
	//pqUrl := string(os.Getenv("ELEPHANTSQL_URL"))
	//fmt.Println(pqUrl)
	//fmt.Printf("%T", pqUrl)
	//pgUrl, err := pq.ParseURL(pqUrl)

	pgUrl, err := pq.ParseURL("postgres://iuvewaik:gGu0xypDpZZsuImooWDl3FRmItpqcWUq@john.db.elephantsql.com:5432/iuvewaik")
	if err != nil {
		log.Fatal(err)
	}

	db, err = sql.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}
