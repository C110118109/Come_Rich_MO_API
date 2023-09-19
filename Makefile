SCRIPTS = $(shell pwd)/scripts
PROJECT = $(shell pwd)
GO = $(shell which go)
OUTPUTS = $(shell pwd)/deploy

authority:
	GOOS=linux GOARCH=amd64 $(GO) build -o $(OUTPUTS)/authority cmd/lambda/authority.go
	zip -D -j -r $(OUTPUTS)/authority.zip $(OUTPUTS)/authority

clean:
	rm -rf $(OUTPUTS)

task:
	make clean
	make accounts

setup:
	cp $(SCRIPTS)/environment.example $(SCRIPTS)/debug.sh
	cp $(SCRIPTS)/environment.example $(SCRIPTS)/testing.sh
	cp $(SCRIPTS)/environment.example $(SCRIPTS)/production.sh
	cp $(SCRIPTS)/migration.example $(SCRIPTS)/migrationDebug.sh
	cp $(SCRIPTS)/migration.example $(SCRIPTS)/migrationTesting.sh
	cp $(SCRIPTS)/migration.example $(SCRIPTS)/migrationProduction.sh
	cp $(PROJECT)/air.example $(PROJECT)/.air.toml

debug:
	chmod a+x $(SCRIPTS)/debug.sh
	$(SCRIPTS)/debug.sh

testing:
	chmod a+x $(SCRIPTS)/testing.sh
	$(SCRIPTS)/testing.sh

production:
	chmod a+x $(SCRIPTS)/production.sh
	$(SCRIPTS)/production.sh

migrationDebug:
	chmod a+x $(SCRIPTS)/migrationDebug.sh
	$(SCRIPTS)/migrationDebug.sh

migrationTesting:
	chmod a+x $(SCRIPTS)/migrationTesting.sh
	$(SCRIPTS)/migrationTesting.sh

migrationProduction:
	chmod a+x $(SCRIPTS)/migrationProduction.sh
	$(SCRIPTS)/migrationProduction.sh

changeLog:
	git-chglog > ./README.md

update_go_mod:
	#Update tools...
	go mod tidy
	go get -d github.com/gin-gonic/gin
	go get -d github.com/swaggo/swag/cmd/swag
	go get -d github.com/smallnest/gen
	#chmod a+x update.sh (air)
	./scripts/update_air.sh
	#brew install golang-migrate
    #go get -d github.com/golang-migrate/migrate

## --- Reference ---
dblink:
	ssh -i ~/.ssh/aws-h1-debian-arm-dev.pem -NTf -L 5435:database-1.cyzkhz6wawgg.ap-southeast-1.rds.amazonaws.com:5432 admin@13.229.140.156

genEntity:
	## https://github.com/smallnest/gen
	gen --sqltype=postgres \
       	--connstr "postgres://postgres:pg654321@localhost:5435/siir_sb?sslmode=disable" \
       	--database siir_sb  \
       	--json \
       	--gorm \
       	--model=dbtable \
       	--out ./internal/gen/enitity