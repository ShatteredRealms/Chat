#####################################################################################
#   _____ _           _   _                    _   _____            _               #
#  / ____| |         | | | |                  | | |  __ \          | |              #
# | (___ | |__   __ _| |_| |_ ___ _ __ ___  __| | | |__) |___  __ _| |_ __ ___  ___ #
#  \___ \| '_ \ / _` | __| __/ _ \ '__/ _ \/ _` | |  _  // _ \/ _` | | '_ ` _ \/ __|#
#  ____) | | | | (_| | |_| ||  __/ | |  __/ (_| | | | \ \  __/ (_| | | | | | | \__ \#
# |_____/|_| |_|\__,_|\__|\__\___|_|  \___|\__,_| |_|  \_\___|\__,_|_|_| |_| |_|___/#
#                               _____ _           _                                 #
#                              / ____| |         | |                                #
#                             | |    | |__   __ _| |_                               #
#                             | |    | '_ \ / _` | __|                              #
#                             | |____| | | | (_| | |_                               #
#                              \_____|_| |_|\__,_|\__|                              #
#####################################################################################
#
# Makefile for building, running, and testing
#

# Application versions
BASE_VERSION = 0.0.1

# Gets the directory containing the Makefile
BUILD_DIR = $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

# Root code directory
ROOT_DIR = $(realpath $(BUILD_DIR)/..)

# Directory containing applications
BASE_APP_DIR = $(realpath $(ROOT_DIR)/cmd)

# Binary output directory
BIN_DIR = $(realpath $(ROOT_DIR)/bin)

# Entrypoint for applications
APP_MAIN = $(BASE_APP_DIR)/chat/main.go

GOLANG_PROTOS = api/backend.proto api/chat.proto api/frontend.proto
#   _____                    _
#  |_   _|                  | |
#    | | __ _ _ __ __ _  ___| |_ ___
#    | |/ _` | '__/ _` |/ _ \ __/ __|
#    | | (_| | | | (_| |  __/ |_\__ \
#    \_/\__,_|_|  \__, |\___|\__|___/
#                  __/ |
#                 |___/

build: build-go all-protos

build-go:
	go build -o $(BIN_DIR)/chat $(APP_MAIN)

test:
	ginkgo $(ROOT_DIR)/... -covermode atomic --randomize-all


all-protos: $(GOLANG_PROTOS)

api/%.proto:
	protoc --go_out=$(ROOT_DIR)/pkg/pb --go_opt=paths=source_relative \
	  --go-grpc_out=$(ROOT_DIR)/pkg/pb --go-grpc_opt=paths=source_relative \
	  -I $(ROOT_DIR)/api \
	  -I $(ROOT_DIR)/api/third_party \
	  $(ROOT_DIR)/$@