import { useState } from 'react'
import { useNavigate } from 'react-router-dom'

const API_URL = import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080'

function Home() {
    const navigate = useNavigate()

    const [hostName, setHostName] = useState('')
    const [creating, setCreating] = useState(false)
    const [joinName, setJoinName] = useState('')
    const [meetingLink, setMeetingLink] = useState('')
    const [joining, setJoining] = useState(false)
    const [error, setError] = useState('')

    const handleCreateMeeting = async () => {
        if (!hostName.trim()) {
            setError('Please enter your name')
            return
        }

        setCreating(true)
        setError('')

        try {
            const response = await fetch(`${API_URL}/api/v1/meetings`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ host_name: hostName.trim() })
            })

            const data = await response.json()

            if (!response.ok) {
                setError(data.error || 'Failed to create meeting')
                return
            }

            navigate(`/room/${data.meeting_id}`, {
                state: {
                    token: data.token,
                    livekitUrl: data.livekit_url,
                    userName: hostName.trim(),
                    isHost: true
                }
            })
        } catch (err) {
            setError('Cannot reach server. Is the backend running?')
        } finally {
            setCreating(false)
        }
    }

    const handleJoinMeeting = async () => {
        if (!joinName.trim()) {
            setError('Please enter your name')
            return
        }
        if (!meetingLink.trim()) {
            setError('Please enter the meeting ID or link')
            return
        }

        let meetingId = meetingLink.trim()
        if (meetingId.includes('/room/')) {
            meetingId = meetingId.split('/room/').pop()
        }

        setJoining(true)
        setError('')

        try {
            const response = await fetch(`${API_URL}/api/v1/meetings/${meetingId}/join`, {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ participant_name: joinName.trim() })
            })

            const data = await response.json()

            if (!response.ok) {
                setError(data.error || 'Failed to join meeting')
                return
            }

            navigate(`/room/${meetingId}`, {
                state: {
                    token: data.token,
                    livekitUrl: data.livekit_url,
                    userName: joinName.trim(),
                    isHost: false
                }
            })
        } catch (err) {
            setError('Cannot reach server. Is the backend running?')
        } finally {
            setJoining(false)
        }
    }

    return (
        <>
            {/* Animated background */}
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
                        <p className="tagline">Meet anyone, anywhere, anytime</p>
                    </div>

                    {error && <div className="error-banner">{error}</div>}

                    <div className="cards-wrapper">
                        {/* Start Meeting */}
                        <div className="card">
                            <div className="card-label">New Session</div>
                            <h2>Start a Meeting</h2>
                            <p className="card-desc">Create a room and invite anyone instantly</p>
                            <div className="input-wrap">
                                <span className="input-icon">👤</span>
                                <input
                                    type="text"
                                    placeholder="Your name"
                                    value={hostName}
                                    onChange={(e) => setHostName(e.target.value)}
                                    onKeyDown={(e) => e.key === 'Enter' && handleCreateMeeting()}
                                />
                            </div>
                            <button
                                className="btn btn-primary"
                                onClick={handleCreateMeeting}
                                disabled={creating}
                            >
                                <span>🚀</span> {creating ? 'Creating...' : 'Start Meeting'}
                            </button>
                        </div>

                        <div className="divider">OR</div>

                        {/* Join Meeting */}
                        <div className="card">
                            <div className="card-label">Have a Link?</div>
                            <h2>Join a Meeting</h2>
                            <p className="card-desc">Enter your name and meeting ID or link</p>
                            <div className="input-wrap">
                                <span className="input-icon">👤</span>
                                <input
                                    type="text"
                                    placeholder="Your name"
                                    value={joinName}
                                    onChange={(e) => setJoinName(e.target.value)}
                                />
                            </div>
                            <div className="input-wrap">
                                <span className="input-icon">🔗</span>
                                <input
                                    type="text"
                                    placeholder="Meeting ID or link"
                                    value={meetingLink}
                                    onChange={(e) => setMeetingLink(e.target.value)}
                                    onKeyDown={(e) => e.key === 'Enter' && handleJoinMeeting()}
                                />
                            </div>
                            <button
                                className="btn btn-secondary"
                                onClick={handleJoinMeeting}
                                disabled={joining}
                            >
                                <span>🔗</span> {joining ? 'Joining...' : 'Join Meeting'}
                            </button>
                        </div>
                    </div>

                    {/* Features strip */}
                    <div className="features">
                        <div className="feat"><div className="feat-dot"></div>No signup required</div>
                        <div className="feat"><div className="feat-dot"></div>End-to-end encrypted</div>
                        <div className="feat"><div className="feat-dot"></div>HD video & audio</div>
                        <div className="feat"><div className="feat-dot"></div>Free forever</div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Home
