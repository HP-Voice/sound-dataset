package main

import (
	"log"
	"os"
)

func base() {
	initFlags()

	err := initConfig(*flags.config)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("db initializing on %s", config.Db)
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
}

func clean() {
	if *flags.clean {
		log.Printf("cleaning up...")
		err := cleanSamples()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("done")
		os.Exit(0)
	}
}

func service() {
	log.Printf("markov initializing")
	err := initMarkov()
	if err != nil {
		log.Fatal(err)
	}

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
}

func main() {
	base()
	clean()
	service()
	select {}
}
