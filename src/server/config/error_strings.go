package config

var DBerrorDescriptions = map[string]string{
  "open": "Could not ping database.",
  "alreadyExists": "Already exists in database.",

}

//formating
//	s := fmt.Sprintf("%s %s %s %s %s", user, password, host, port, _db)
//   log.Printf(s)