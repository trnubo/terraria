.PHONY: help testserver build
help:
	@echo "No help"

testserver:
	docker run -d --name terraria \
	  -v $(PWD)/testserver/config:/config \
	  -v $(PWD)/testserver/Worlds:/root/.local/share/Terraria/Worlds \
	  trnubo/terraria:latest -config /config/serverconfig.txt

docker-shell:
	docker run --rm -it --name terraria \
	  -v $(PWD)/testserver/config:/config \
	  -v $(PWD)/testserver/Worlds:/root/.local/share/Terraria/Worlds \
	  --entrypoint bash trnubo/terraria:latest

build:
	( cd TerrariaServerWrapper; make snapshot; )
	( cd tshock; docker build -t trnubo/terraria:latest .; )
