default:
	docker-compose up --build

clean:
	docker-compose down

.PHONY: clean default
