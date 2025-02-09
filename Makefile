### Makefile
build:
	@echo "Building VoiceBot..."
	go build -o voicebot

run: build
	@echo "Starting VoiceBot..."
	./voicebot

docker-build:
	@echo "Building Docker image..."
	docker build -t voicebot .

docker-run: docker-build
	@echo "Running VoiceBot in Docker..."
	docker run -p 8080:8080 voicebot
