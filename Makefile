.PHONY: help build up debug down start stop shell
.DEFAULT_GOAL: help

default: help

help:    ## Output available commands
	@echo "Available commands:"
	@echo
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build:   ## Build the docker image
	@docker compose build

up:      ## Start all services in detached mode
	@docker compose up -d

debug:   ## Start api in debug mode
	@docker compose up -d
	@docker attach $$(docker compose ps -q api)

down:    ## Bring all services that are in detached mode down
	@docker compose down

start:   ## Start stopped services
	@docker compose start

stop:    ## Stop all services
	@docker compose stop

shell:   ## Lunch a bash shell in running api container
	@docker compose exec api bash
