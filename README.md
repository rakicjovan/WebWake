# WebWake

![Go](https://img.shields.io/badge/Go-1.24-blue)
![Docker](https://img.shields.io/badge/Docker-Supported-blue)
![License](https://img.shields.io/badge/License-MIT-green)

> Simple lightweight HTTP server to send Wake-on-LAN packets.

---

## 🚀 Features

- Lightweight Go application
- Accepts simple HTTP `POST` requests
- Sends standard Wake-on-LAN magic packets
- Docker support (Linux-only with Host Networking)

---

## 📦 Usage

Send an HTTP `POST` request with a JSON body specifying the target device's MAC address.

Example using `curl`:

```bash
curl -X POST "http://localhost:8080/wake" \
  -H "Content-Type: application/json" \
  -d '{"mac": "AA:BB:CC:DD:EE:FF"}'
```

---

## 🛠 Building

If you want to build and run the application manually:

```bash
go build -o webwake main.go
./webwake
```

---

## 🐳 Docker

A `Dockerfile` is provided for containerized deployments.

Because Wake-on-LAN operates over Layer 2 (Ethernet broadcast), the container **must** run with **host networking** enabled.

### Build and Run

```bash
docker build -t webwake .
docker run --network host webwake
```

### Important Notes

- **Host networking** is only available on **Linux**.
- On **macOS** and **Windows** (Docker Desktop), `--network host` behaves differently and will not allow proper Layer 2 broadcasts.
- If you are not using Linux, it is recommended to run the compiled Go binary directly on the host machine instead of inside Docker.

---

## ⚙️ Configuration

| Environment Variable | Description            | Default |
|:---------------------|:------------------------|:--------|
| `PORT`                | HTTP port to listen on   | `8080`  |

Example running on a custom port:

```bash
PORT=9090 ./webwake
```

---

## 📜 License

This project is licensed under the MIT License.  

---
