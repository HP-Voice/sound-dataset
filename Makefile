all:
	rm -rf ./build
	cd ./back; go build
	cd ./front; npm run build
	mkdir -p ./build
	mv ./back/back ./build/
	cp ./back/config.json ./build/
	mv ./front/dist ./build/front
	cp ./markov/* ./build/
