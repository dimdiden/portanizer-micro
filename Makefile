protogen:
	for service in auth gateway users workbook; do \
		cd ./services/$$service ; \
		sh ./compile.sh; \
		cd ../../ ; \
	done
vendor:
	for service in auth gateway users workbook; do \
		cd ./services/$$service ; \
		go mod tidy ; \
    	go mod vendor ; \
    	cd ../.. ; \
	done
build:
	docker-compose up --build
rebuild:
	make vendor
	make build