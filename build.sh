#!/usr/bin/env bash
set -e


go get -t ./...

GREEN='\033[0;32m'
NC='\033[0m'



printf "\n\n\n\n\n\n\n\n\n\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}=                        E N S U R E    C L E A N    A N D   T E S T E D                           =${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
./hook.sh true



printf "\n\n\n\n\n\n\n\n\n\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}=                                       B U I L D     A P P                                        =${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}=                                                                                                  =${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
printf "${GREEN}====================================================================================================${NC}\n"
cd CLI/task
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
cd ../..

cd image-transformation
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
cd ..

cd recover
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
cd ..

cd secret
CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o main .
cd ..

