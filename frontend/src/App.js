import React, { useState, useEffect } from 'react';

import ChatHistory from './components/ChatHistory';
import ChatInput from './components/ChatInput';
import Header from './components/Header';

import {socketConnect, sendMsg} from './socketApi'
import fetchChatHistory from './restApi';



function App() {

  // 1. init chat history
  const [chatHistory, setChatHistory] = useState([]);

  // 2. connect socket
  useEffect(() => {

    // - fetch all history
    async function fetchData() {
      const messages = await fetchChatHistory();
      setChatHistory(messages);
    }
    fetchData();

    // - connect to socket with call back function
    socketConnect((msgEvent) => {

      const jsonData = JSON.parse(msgEvent.data) // contains type and body

      // get message from sockets and add to history if there's a new incoming message
      setChatHistory(prevChatHistory => [...prevChatHistory, jsonData.body]);

    });
  }, []);

  return (
    <div className="App">
      <Header/>
      <ChatHistory chatHistory={chatHistory}/>
      <ChatInput sendFunc={sendMsg}/>
    </div>
  );
}

export default App;
