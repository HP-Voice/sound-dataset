all:
	cd ./back; go build
	cd ./front; npm run build
	mkdir -p ./build/static
	cp -rf ./back/back ./build/server
	cp -rf ./back/config.json ./build/
	cp -rf ./front/dist/* ./build/static
	cp -rf ./markov/* ./build/
