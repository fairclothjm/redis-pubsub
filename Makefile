default:
	docker-compose up

clean:
	docker-compose down

.PHONY: clean default
