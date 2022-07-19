all:
	rm -rf ./build
	cd ./back; go build
	cd ./front; npm run build
	mkdir -p ./build/front
	cp ./back/back ./build/
	cp ./back/config.json ./build/
	cp -r ./front/dist/* ./build/front/
	cp ./markov/* ./build/
