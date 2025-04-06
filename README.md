# microUrl

A lightweight URL shortening service written in Go.

[English]

## Description
microUrl is a simple and efficient URL shortening service that converts long URLs into shorter, more manageable links. It uses SHA hashing and base64 encoding to generate compact URLs while maintaining uniqueness.

## Features
- URL shortening with customizable length
- Optional padding support
- SHA-based hashing for URL generation
- Base64 encoding for URL-safe strings
- Simple and efficient implementation

## Installation

### As a Library
```bash
go get github.com/SaatvikAwasthi/microUrl
```

### From Source
```bash
git clone https://github.com/SaatvikAwasthi/microUrl.git
cd microUrl
go mod download
```

## Usage

### As a Library
```go
import "github.com/SaatvikAwasthi/microUrl/pkg/url"

// Basic usage
shortUrl, err := url.Shrink("https://example.com/very/long/url", 8, false)

// With padding
shortUrl, err := url.Shrink("https://example.com/very/long/url", 8, true)
```

### Command Line Interface
```bash
# Basic usage with default length (5)
microUrl https://example.com/very/long/url

# Specify custom length
microUrl -l 8 https://example.com/very/long/url

# Use padding
microUrl -p https://example.com/very/long/url

# Use both custom length and padding
microUrl -l 8 -p https://example.com/very/long/url
```

## Project Structure
```
microUrl/
├── cmd/            # Command line applications
├── pkg/            # Library code
│   ├── url/        # URL shortening logic
│   ├── hash/       # Hashing implementation
│   └── utils/      # Utility functions
└── micro/          # Microservice implementation
```
