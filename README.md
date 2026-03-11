# OpenChat

A simple video meeting app — create a room, share the link, jump in. No signup, no friction.

Built with Go (backend) and React (frontend), using LiveKit for real-time video.

## Tech

- **Backend** — Go, Gin, LiveKit SDK
- **Frontend** — React, Vite, LiveKit Components
- **Video** — LiveKit Cloud

## How it works

The backend follows a clean layered structure: Router → Controller → Service → Repository. Each layer has one job and talks only to the next one down.

## Run locally

**Backend** — create a `backend/.env` file with your LiveKit credentials:
```
LIVEKIT_URL=your_livekit_url
LIVEKIT_API_KEY=your_api_key
LIVEKIT_API_SECRET=your_api_secret
FRONTEND_URL=http://localhost:5173
PORT=8080
```
```bash
cd backend
go run main.go
```

**Frontend** — create a `frontend/.env` file:
```
VITE_BACKEND_URL=http://localhost:8080
```
```bash
cd frontend
npm install
npm run dev
```

## Deploy

Backend on Railway, frontend on Vercel. Set env vars in each dashboard — no config files needed in production.

## Live demo

[open-chat-ochre.vercel.app](https://open-chat-ochre.vercel.app)
