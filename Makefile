TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=hashicorp.com
NAMESPACE=epilot
NAME=epilot
BINARY=terraform-provider-${NAME}
VERSION=0.0.2
OS_ARCH=darwin_amd64

default: install

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

test: 
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

testacc: 
	TF_VAR_TOKEN="eyJraWQiOiJ2ZFR0MGQrK1RMc2FQZ2tsQ3AzMDVGbEMxc1lOUCtUOXpsaElzMkJ3WERrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIxNzEyMTkwMy1kM2JlLTRhZTktODZiZS04YjhkZDRmYzY0ZTYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfaGh6MnVJQ2xIIiwicGhvbmVfbnVtYmVyX3ZlcmlmaWVkIjp0cnVlLCJjdXN0b206aXZ5X29yZ19pZCI6IjY2IiwiY29nbml0bzp1c2VybmFtZSI6Im4uZ29lbEBlcGlsb3QuY2xvdWQiLCJjdXN0b206aXZ5X3VzZXJfaWQiOiI4MjYwMiIsImF1ZCI6ImdqOXAwanJlaWh0cTAwY3JpNmEwZmUzMDYiLCJldmVudF9pZCI6ImUyYmNkZDI4LWU5NWUtNDk3Ni04Y2Q4LTg0MTlmMmYyNGU3YyIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjcxMTIwMzE1LCJwaG9uZV9udW1iZXIiOiIrOTE5OTcxNTQzNDUzIiwiZXhwIjoxNjcxODA5MTUwLCJpYXQiOjE2NzE4MDU1NTAsImVtYWlsIjoibi5nb2VsQGVwaWxvdC5jbG91ZCJ9.olDoZb38KrL28j7yol9vAaFENogN1necRlNM3fH4m9ycODig-Fp3Xe98JD4fgNXA74K3td5hlMiAfS4DeJDh_Dg-IHNjW3EhM_SWEzt1YmmOTw2vU39q330yWDPRzYEZw7HzPae3uCnjWG_OHw5A73Cl0lF-v9CjjQAUVK5vJAZ2P9DMV-wnPOtc5uatRpdCdjpZT9wtBOgUbbTbs8qsqDM1p05STCMXWTSBa1kMIuX88NaCB6583jHkpqyRnu0eWwH65un2_DVxJT1wyH-mjCm-mtrTuU-ZuFFAvWAIcpZbJT5qgaqF-vpJY8X3Dt1ZKTaf8bmvJ0gel9IFcCdtDQ" TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   