import { useEffect, useState } from "react";
import SendTab from "./SendTab";
import ReceiveTab from "./ReceiveTab";
import { EventsOn } from "../wailsjs/runtime/runtime"
import "./App.css";
import {Brodudp, Listentcp} from "../wailsjs/go/main/App"
function App() {
   const[ipadd,setipadd] = useState("")
    useEffect(()=>{
        Brodudp()
        Listentcp()
        EventsOn("ipadd",(data)=>{
          setipadd(data)
        })
    }
    ,[])
  const [tab, setTab] = useState("send");

  return (
    <div className="app">

      {/* Sidebar */}

      <div className="sidebar">

        <h2>P2P Sharing Platform</h2>
        <p>Your IP Address : {ipadd}</p>
        <button onClick={() => setTab("send")}>
          ➤ Send 
        </button>

        <button onClick={() => setTab("receive")}>
          🡻 Receive 
        </button>

      </div>

      {/* Content */}

      <div className="content">

        {tab === "send" && <SendTab />}
        {tab === "receive" && <ReceiveTab />}

      </div>

    </div>
  );
}

export default App;