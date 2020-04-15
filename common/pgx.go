package common

type DBinfo struct {
	Dev  map[string]string
	Test map[string]string
	Prod map[string]string
}

// func Connect(c string) error {
// 	pgurl := config.pgurl
// 	conn, err := pgx.Connect(context.Background.pgurl)
// 	if err != nil {
// 		(os.Stderr, "Unable to connection to database: %v\n", err)
// 	}
// }
