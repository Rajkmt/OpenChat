import { useParams, useLocation, useNavigate } from 'react-router-dom'
import { LiveKitRoom, VideoConference } from '@livekit/components-react'
import '@livekit/components-styles'
import { useState } from 'react'

const API_URL = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080'

function Room() {
    const { meetingId } = useParams()
    const location = useLocation()
    const navigate = useNavigate()

    const [token, setToken] = useState(location.state?.token || null)
    const [livekitUrl, setLivekitUrl] = useState(location.state?.livekitUrl || null)
    const [userName, setUserName] = useState(location.state?.userName || '')
    const [nameInput, setNameInput] = useState('')
    const [joining, setJoining] = useState(false)
    const [error, setError] = useState('')
    const [disconnected, setDisconnected] = useState(false)

    const handleDirectJoin = async () => {
        if (!nameInput.trim()) {
            setError('Please enter your name')
            return
        }

        setJoining(true)
        setError('')

        try {
            const response = await fetch(`${API_URL}/api/v1/meetings/${meetingId}/join`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ participant_name: nameInput.trim() })
            })

            const data = await response.json()

            if (!response.ok) {
                setError(data.error || 'Failed to join meeting')
                return
            }

            setToken(data.token)
            setLivekitUrl(data.livekit_url)
            setUserName(nameInput.trim())
        } catch (err) {
            setError('Cannot reach server. Is the backend running?')
        } finally {
            setJoining(false)
        }
    }

    const handleDisconnect = () => {
        setDisconnected(true)
    }

    if (disconnected) {
        return (
            <>
                <div className="bg-scene">
                    <div className="orb orb1"></div>
                    <div className="orb orb2"></div>
                    <div className="orb orb3"></div>
                </div>
                <div className="grid-lines"></div>
                <div className="home-container">
                    <div className="home-content">
                        <div className="card" style={{ textAlign: 'center' }}>
                            <div style={{ fontSize: '3rem', marginBottom: '1rem' }}>👋</div>
                            <h2>You left the meeting</h2>
                            <p className="card-desc">Meeting ID: {meetingId}</p>
                            <button className="btn btn-primary" onClick={() => navigate('/')}>
                                ← Back to Home
                            </button>
                        </div>
                    </div>
                </div>
            </>
        )
    }

    if (!token) {
        return (
            <>
                <div className="bg-scene">
                    <div className="orb orb1"></div>
                    <div className="orb orb2"></div>
                    <div className="orb orb3"></div>
                </div>
                <div className="grid-lines"></div>
                <div className="home-container">
                    <div className="home-content">
                        <div className="brand">
                            <div className="brand-icon">💬</div>
                            <h1>OpenChat</h1>
                        </div>

                        {error && <div className="error-banner">{error}</div>}

                        <div className="card">
                            <div className="card-label">Join Room</div>
                            <h2>Join Meeting</h2>
                            <p className="card-desc">Meeting ID: <strong>{meetingId}</strong></p>
                            <div className="input-wrap">
                                <span className="input-icon">👤</span>
                                <input
                                    type="text"
                                    placeholder="Enter your name"
                                    value={nameInput}
                                    onChange={(e) => setNameInput(e.target.value)}
                                    onKeyDown={(e) => e.key === 'Enter' && handleDirectJoin()}
                                />
                            </div>
                            <button
                                className="btn btn-primary"
                                onClick={handleDirectJoin}
                                disabled={joining}
                            >
                                <span>🔗</span> {joining ? 'Joining...' : 'Join Meeting'}
                            </button>
                        </div>
                    </div>
                </div>
            </>
        )
    }

    return (
        <div className="room-container">
            <div className="room-header">
                <span className="room-title">OpenChat</span>
                <span className="room-info">
                    Meeting: <strong>{meetingId}</strong> · {userName}
                </span>
            </div>
            <div className="room-video">
                <LiveKitRoom
                    serverUrl={livekitUrl}
                    token={token}
                    connect={true}
                    onDisconnected={handleDisconnect}
                    data-lk-theme="default"
                    style={{ height: 'calc(100vh - 56px)' }}
                >
                    <VideoConference />
                </LiveKitRoom>
            </div>
        </div>
    )
}

export default Room
