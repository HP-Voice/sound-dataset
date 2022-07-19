package main

import "log"

func main() {
	initFlags()

	err := initConfig(*flags.config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("markov initializing")
	err = initMarkov()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("db initializing on %s@%s:%d/%s\n", config.Db.User, config.Db.Host, config.Db.Port, config.Db.Database)
	err = initDb()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("fs initializing in %s\n", config.Fs.Path)
	err = initFs()
	if err != nil {
		log.Fatal(err)
	}

	initRandom()

	log.Printf("api server initializing on %s\n", config.Api.Address)
	go func() {
		err = initApi()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("static server initializing on %s\n", config.Static.Address)
	go func() {
		err = initStatic()
		if err != nil {
			log.Fatal(err)
		}
	}()

	select {}
}
