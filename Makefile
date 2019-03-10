rerun:
	for service in auth gateway users workbook; do \
		cd $$service ; \
		go mod tidy ; \
    	go mod vendor ; \
    	cd .. ; \
	done
	make run
run:
	docker-compose up --build