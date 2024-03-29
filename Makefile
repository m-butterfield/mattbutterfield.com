cloudrunbasecommand := gcloud run deploy --project=mattbutterfield --region=us-central1 --platform=managed
deployservercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/mattbutterfield.com mattbutterfield
deployworkercommand := $(cloudrunbasecommand) --image=gcr.io/mattbutterfield/mattbutterfield.com-worker mattbutterfield-worker

terraformbasecommand := cd infra && terraform
terraformvarsarg := -var-file=secrets.tfvars

export DB_SOCKET=host=localhost dbname=mattbutterfield

build: build-server build-worker

build-server:
	go build -o bin/server cmd/server/main.go

build-worker:
	go build -o bin/worker cmd/worker/main.go

deploy: docker-build docker-push
	$(deployservercommand)
	$(deployworkercommand)

deploy-server: docker-build-server docker-push-server
	$(deployservercommand)

deploy-worker: docker-build-worker docker-push-worker
	$(deployworkercommand)

docker-build:
	docker-compose build

docker-build-server:
	docker-compose build server

docker-build-worker:
	docker-compose build worker

docker-push:
	docker-compose push

docker-push-server:
	docker-compose push server

docker-push-worker:
	docker-compose push worker

reset-db:
	dropdb --if-exists mattbutterfield
	createdb mattbutterfield
	psql -f dump.sql mattbutterfield

migrate:
	go run cmd/migrate/main.go

fmt:
	go fmt ./...
	npx eslint app/static/js/ --fix
	cd infra/ && terraform fmt

run-server: export USE_LOCAL_FS=true
run-server: export SQL_LOGS=true
run-server: export WORKER_BASE_URL=http://localhost:8001/
run-server: export AUTH_TOKEN=1234
run-server:
	go run cmd/server/main.go

run-worker: export SQL_LOGS=true
run-worker:
	go run cmd/worker/main.go

test: export DB_SOCKET=host=localhost dbname=mattbutterfield_test
test: export UPLOADER_SERVICE_ACCOUNT=ewogICJ0eXBlIjogInNlcnZpY2VfYWNjb3VudCIsCiAgInByb2plY3RfaWQiOiAiY29vbHByb2plY3QiLAogICJwcml2YXRlX2tleV9pZCI6ICIxMjMiLAogICJwcml2YXRlX2tleSI6ICItLS0tLUJFR0lOIFBSSVZBVEUgS0VZLS0tLS1cbk1JSUV2Z0lCQURBTkJna3Foa2lHOXcwQkFRRUZBQVNDQktnd2dnU2tBZ0VBQW9JQkFRRFkzRThvMU5FRmNqTU1cbkhXLzVaZkZKdzI5LzhORXFwVmlOalFJeDk1WHg1S0R0SituV245K09XMHVxc1NxS2xLR2hBZEFvK1E2Ymp4MmNcbnVYVnNYVHU3WHJaVVk1S2x0dmo5NER2VWExd2pOWHM2MDZyL1J4V1RKNThiZmRDK2dMTHhCZkduQjZDd0swWVFcbnhuZnBqTmJrVWZWVnpPME1RRDdVUDBIbDVaY1kwUHV2eGQveUh1T05Rbi9ySUFpZVRISDFwcWdXK3pySC95M2NcbjU5SUdUaEM5UFB0dWdJOWVhOFJTblZqM1BXejFiWDJVa0NEcHk5SVJoOUx6SkxhWVlYOVJVZDcrK2RVTFVsYXRcbkFhWEJoMVU2ZW1VRHpocklzZ0FwakRWdGltT1BibVFXbVgxUzYwbXFRaWtScFZZWjh1K05ERCtMTncrL0Vvdm5cbnhDajJZM3oxQWdNQkFBRUNnZ0VBV0RCem9xTzFJdlZYakJBMmxxSWQxMFQ2aFhtTjNqMWlmeUgrYUFxSytGVmxcbkdqeVdqRGoweFdRY0o5eW5jN2JRNmZTZVRlTkd6UDBNNmt6RFUxK3c2Rmd5WnF3ZG1YV0kyVm1FaXpSandrKy9cbi91TFFVY0w3STU1RHhuN0tVb1pzL3JaUG1RRHhtR0xvdWU2MEdnNnozeUx6VmNLaURjN2NuaHpoZEJnRGM4dmRcblFvck5BbHFHUFJubTNFcUtRNlZRcDZmeVFtQ0F4cnI0NWtzcFJYTkxkZGF0M0FNc3VxSW1Ea3FHS0JtRjNRMXlcbnhXR2U4MUxwaFVpUnF2cWJ5VWxoNmNkU1o4cExCcGM5bTBjM3FXUEtzOXBhcUJJdmdVUGx2T1pNcWVjNng0UzZcbkNoYmRra1RSTG5ic1JyMFlnL25EZUVQbGtoUkJoYXNYcHhwTVVCZ1B5d0tCZ1FEczJheE5rRmpiVTk0dVh2ZDVcbnpuVWhEVnhQRkJ1eHlVSHRzSk5xVzRwL3VqTE5pbUdldDVFL1l0aENuUWVDMlAzWW03YzNmaXo2OGFtTTZoaUFcbk9uVzdIWVBaK2pLRm5lZnBBdGp5T09zNDZBa2Z0RWcwN1Q5WGp3V05QdDgrOGwwRFlhd1BvSmdiTTVpRTBMMk9cbng4VFUxVnM0bVhjK3FsOUY5MEd6STB4M1Z3S0JnUURxWk9PcVd3M2hUbk5UMDdJeHFubWQzZHVnVjlTN2VXNm9cblU5T29VZ0pCNHJZVHBHK3lGcU5xYlJUOGJreDM3aUtCTUVSZXBwcW9uT3FHbTR3dHVSUjZMU0xsZ2NJVTlJd3hcbnlmSDEyVVdxVm1GU0hzZ1pGcU0vY0szd0dldjM4aDFXQklPeDMvZGpLbjdCZGxLVmg4a1d5eDZ1QzhibVYrRTZcbk9vSzB2SkQ2a3dLQmdIQXlTT25ST0JabHF6a2lLVzhjK3VVMlZBVHR6SlN5ZHJXbTBKNHdVUEppZk5CYS9oVldcbmRjcW1BelhDOXh6bnQ1QVZhM3d4SEJPZnlLYUUraWc4Q1Nzak55TlozdmJtcjBYMDRGb1YxbTkxazJUZVhOb2RcbmpNVG9ia1BUaGFObTRlTEpNTjJTUUp1YUhHVEdFUldDMGwzVDE4dCsvenJETURDUGlTTFgxTkF2QW9HQkFOMVRcblZMSllkanZJTXhmMWJtNTlWWWNlcGJLN0hMSEZrUnE2eE1KTVpidEcwcnlyYVpqVXpZdkI0cTRWakhrMlVEaUNcbmxoeDEzdFhXRFpIN01KdEFCemp5ZytBSTdYV1NFUXMyY0JYQUNvczBNNE15YzZsVStlTCtpQStPdW9VT2htcmhcbnFtVDhZWUd1NzYvSUJXVVNxV3V2Y3BIUHB3bDc4NzFpNEdhL0kzcW5Bb0dCQU5Oa0tBY01vZUFiSlFLN2EvUm5cbndQRUpCK2RQZ05ESWFib0FzaDFuWmhWaE41Y3ZkdkNXdUVZZ09HQ1BRTFlRRjB6bVRMY00rc1Z4T1lnZnk4bVZcbmZiTmdQZ3NQNXhtdTZkdzJDT0JLZHRvencwSHJXU1JqQUNkMU40eUd1NzUrd1BDY1gvZ1FhcmNqUmNYWFplRWFcbk50QkxTZmNxUFVMcUQraDdicjlsRUppb1xuLS0tLS1FTkQgUFJJVkFURSBLRVktLS0tLVxuIiwKICAiY2xpZW50X2VtYWlsIjogIjEyMy1hYmNAZGV2ZWxvcGVyLmdzZXJ2aWNlYWNjb3VudC5jb20iLAogICJjbGllbnRfaWQiOiAiMTIzLWFiYy5hcHBzLmdvb2dsZXVzZXJjb250ZW50LmNvbSIsCiAgImF1dGhfdXJpIjogImh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbS9vL29hdXRoMi9hdXRoIiwKICAidG9rZW5fdXJpIjogImh0dHBzOi8vb2F1dGgyLmdvb2dsZWFwaXMuY29tL3Rva2VuIiwKICAiYXV0aF9wcm92aWRlcl94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL29hdXRoMi92MS9jZXJ0cyIsCiAgImNsaWVudF94NTA5X2NlcnRfdXJsIjogImh0dHBzOi8vd3d3Lmdvb2dsZWFwaXMuY29tL3JvYm90L3YxL21ldGFkYXRhL3g1MDkvMTIzYWJjJTQwZGV2ZWxvcGVyLmlhbS5nc2VydmljZWFjY291bnQuY29tIgp9
test:
	dropdb --if-exists mattbutterfield_test && createdb mattbutterfield_test
	go run cmd/migrate/main.go
	go test -v ./app/...

tf-plan:
	$(terraformbasecommand) plan $(terraformvarsarg)

tf-apply:
	$(terraformbasecommand) apply $(terraformvarsarg)

tf-refresh:
	$(terraformbasecommand) apply $(terraformvarsarg) -refresh-only

update-deps:
	go get -u ./...
	go mod tidy
	npm upgrade
	cd infra && terraform init -upgrade && cd -
