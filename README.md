### README.md
# VoiceBot - A TCP-Based Voice Processing Bot

## Features
- Handles audio streaming over TCP
- Supports UUID and audio message types
- Built with Go and Docker

## Installation & Usage
### Run Locally
```sh
git clone https://github.com/nikhildhole/voicebot.git
cd voicebot
go run cmd/server/main.go
```

### Run with Docker
```sh
docker-compose up --build
```

## API Protocol
VoiceBot communicates using a simple binary protocol:
- **0x01 (UUID)** - Client identifier
- **0x10 (Audio)** - Audio data
- **0x00 (Terminate)** - Disconnect

## Contributing
Pull requests are welcome! For major changes, please open an issue first.

## License
MIT License
