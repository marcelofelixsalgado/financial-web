compile:
	echo "Compiling for linux Platform"
	GOOS=linux GOARCH=386 go build -o bin/financial-web-386 main.go

build:
	go build -o bin/main main.go

run:
	docker compose up --build

deploy:
	 scp -i "~/.ssh/id_ed25519_aws_financial.pem" ./bin/financial-web-386 ubuntu@ec2-44-211-189-217.compute-1.amazonaws.com:~/financial-web/.
	 scp -r -i "~/.ssh/id_ed25519_aws_financial.pem" ./web ubuntu@ec2-44-211-189-217.compute-1.amazonaws.com:~/financial-web/.