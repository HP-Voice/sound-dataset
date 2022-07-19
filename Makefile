all:
	cd ./back; go build
	cd ./front; npm run build
	mkdir -p ./build/static
	mv ./back/back ./build/server
	cp ./back/config.json ./build/
	mv ./front/dist/* ./build/static
	cp ./markov/* ./build/
