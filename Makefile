
clean:
	rm -rf models/* restapi/*

gen: clean
	swagger generate server --target ./ --name Swagger --spec ./swagger.yml --exclude-main
