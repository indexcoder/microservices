SHELL=cmd.exe
FRONT_END_BINARY=frontApp.exe
BROKER_BINARY=brokerApp
AUTH_BINARY=authApp
LOGGER_BINARY=loggerServiceApp
MAIL_BINARY=mailerApp
LISTENER_BINARY=listenerApp
FRONT_BINARY=frontEndApp

up:
	@echo Starting Docker image...
	docker-compose up -d
	@echo Docker image started

up_build: build_broker build_auth build_logger build_mail build_listener
	@echo Stopping docker image (if running...)
	docker-compose down
	@echo Building (when required) and starting docker image...
	docker-compose up --build -d
	@echo Docker images build and started!

down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

build_broker:
	@echo building broker binary...
	@cd ../broker-service && \
	set "GOOS=linux" && \
	set "GOARCH=amd64" && \
	set "CGO_ENABLED=0" && \
	go build -o $(BROKER_BINARY) ./cmd/api
	@echo Done!

build_logger:
	@echo building logger binary...
	@cd ../logger-service && \
	set "GOOS=linux" && \
	set "GOARCH=amd64" && \
	set "CGO_ENABLED=0" && \
	go build -o $(LOGGER_BINARY) ./cmd/api
	@echo Done!

build_listener:
	@echo building listener binary...
	@cd ../listener-service && \
	set "GOOS=linux" && \
	set "GOARCH=amd64" && \
	set "CGO_ENABLED=0" && \
	go build -o $(LISTENER_BINARY) .
	@echo Done!

build_mail:
	@echo building mail binary...
	@cd ../mail-service && \
	set "GOOS=linux" && \
	set "GOARCH=amd64" && \
	set "CGO_ENABLED=0" && \
	go build -o $(MAIL_BINARY) ./cmd/api
	@echo Done!

build_auth:
	@echo building auth binary...
	@cd ../authentication-service && \
	set "GOOS=linux" && \
	set "GOARCH=amd64" && \
	set "CGO_ENABLED=0" && \
	go build -o $(AUTH_BINARY) ./cmd/api
	@echo Done!

build_front:
	@echo Building front end binary...
	chdir ..\front-end && set CGO_ENABLED=0&& set GOOS=windows&& go build -o ${FRONT_BINARY} ./cmd/web
	@echo Done!

start: build_front
	@echo Starting front end...
	cd ..\front-end && start ${FRONT_END_BINARY}
	@echo Done

stop:
	@echo Stoping front-end...
	@taskkill /IM "${FRONT_END_BINARY}" /F
	@echo "Front end stopped!"