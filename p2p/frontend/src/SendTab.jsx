import { useEffect, useState } from "react"
import { Sendtcp } from "../wailsjs/go/main/App"
import { EventsOn } from "../wailsjs/runtime/runtime"
import { SelectFile } from "../wailsjs/go/main/App"
import "./SendTab.css"
function SendTab() {

    const [peers, setPeers] = useState([])
    const [selectedPeer, setSelectedPeer] = useState(null)
    const [filePath, setFilePath] = useState("");

    const chooseFile = async () => {
        const path = await SelectFile();

        if (path) {
            setFilePath(path);
            console.log("Selected:", path);
        }
    };

    useEffect(() => {

        EventsOn("peer_list", (list) => {
            setPeers(list)
        })

    }, [])

    return (

        <div>

            <h2>Available Devices</h2>

            {peers.length === 0 && <p>No devices found</p>}

            {peers.map((p, i) => (

                // <div
                //     key={i}
                //     onClick={() => setSelectedPeer(p)}
                //     style={{
                //         padding: "10px",
                //         margin: "5px",
                //         border: "1px solid gray",
                //         cursor: "pointer",
                //         background: selectedPeer === p ? "#ddd" : "white"
                //     }}
                // >
                //     {p}
                // </div>
                <div
                key={i}
                onClick={() => setSelectedPeer(p)}
                className={`device-card ${selectedPeer === p ? "selected" : ""}`}
                >
                {p}
                </div>
            ))}

            <br />

            <button onClick={chooseFile}>Choose a File</button>
            {filePath && (
            <p id="showfname">
                Selected: <b>{filePath.split("\\").pop()}</b>
            </p>
            )}
            <button disabled={!selectedPeer} onClick={() => {
                if(filePath=="" || selectedPeer == null)return
                else{
                Sendtcp(selectedPeer,filePath)
                setFilePath("")
                }}}>
                Send to {selectedPeer || "device"}
            </button>

        </div>
    )
}

export default SendTab