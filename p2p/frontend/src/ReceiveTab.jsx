import { useEffect, useState } from "react"
import { EventsOn } from "../wailsjs/runtime/runtime"
import { AllowFile, RejectFile } from "../wailsjs/go/main/App"
import "./ReceiveTab.css"
function ReceiveTab() {

  const [requests, setRequests] = useState([])
  const [progress, setProgress] = useState({})

  useEffect(() => {
    EventsOn("file_request", (data) => {

      setRequests(prev => {
        if (prev.some(r => r.id === data.id)) return prev
        return [...prev, data]
      })

    })
    EventsOn("recv_progress", (data) => {

      setProgress(prev => ({
        ...prev,
        [data.id]: {
          progress: data.progress,
          name: data.name
        }
      }))

    })

  }, [])

  function allow(id) {
    setRequests(req => req.filter(r => r.id !== id))
    AllowFile(id)

  }

  function reject(id) {
    setRequests(req => req.filter(r => r.id !== id))
    RejectFile(id)
  }

  return (

    <div>

      <h2>Incoming Files</h2>

      {requests.length === 0 && <p>No incoming requests</p>}

      {requests.map((r) => (

        <div key={r.id} className="request-card">
          <p><b>{r.name}</b></p>
          <p>{Math.round(r.size / 1024 / 1024)} MB</p>

          <button onClick={() => allow(r.id)}>Accept</button>
          <button onClick={() => reject(r.id)}>Reject</button>

        </div>

      ))}

      <h2>Transfers</h2>

      {Object.keys(progress).map((id) => (

        <div key={id} className="transfer">
          <span>{progress[id].name} : </span>
          <progress value={progress[id].progress} max="100"></progress>
          <span> {progress[id].progress}%</span>

        </div>

      ))}

    </div>

  )

}

export default ReceiveTab