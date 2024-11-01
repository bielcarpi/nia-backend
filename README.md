<h1 align="center">NIA Backend</h1>

<p align="center">
  <a href="https://go.dev">
    <img src="https://img.shields.io/badge/Go-1.19+-00ADD8.svg">
  </a>
  <a href="https://cloud.google.com">
    <img src="https://img.shields.io/badge/GCP-App%20Engine-red.svg">
  </a>
  <a href="https://platform.openai.com">
    <img src="https://img.shields.io/badge/OpenAI-API-green.svg">
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg">
  </a>
</p>

<p align="center">
The NIA Backend powers the NIA language learning assistant with real-time audio processing and AI integration. Built with Go and optimized for low-latency communication, it handles all AI interactions and audio streaming between the mobile app and OpenAI's services.
</p>

## Features
- [x] WebSocket Audio Streaming
- [x] OpenAI API Integration
- [x] Authentication Middleware
- [x] Low-latency Processing
- [x] Cloud Deployment
- [ ] Self-hosted AI Models

## Project Structure
```
.
├── config/           # Configuration management
├── pkg/
│   ├── api/         # HTTP/WebSocket handlers
│   ├── openai/      # OpenAI API integration
│   └── util/        # Helper functions
└── migrations/      # Database migrations
```

## Getting Started
### Prerequisites
- Go 1.19+
- OpenAI API key
- Google Cloud SDK

### Installation
```bash
# Clone the repository
git clone https://github.com/bielcarpi/nia_backend.git

# Install dependencies
go mod download

# Run locally
go run main.go
```

### Deployment
```bash
# Deploy to Google Cloud App Engine
gcloud app deploy
```

## Architecture
The backend implements a clean architecture pattern with:
- WebSocket handlers for real-time communication
- OpenAI service integration (GPT-3.5, Whisper, TTS)
- Authentication middleware
- Helper utilities

## Audio Processing
- Real-time streaming via WebSocket
- AAC codec support
- Optimized for low latency
- Future PCM16 support planned
