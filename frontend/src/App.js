import React, { useState, useEffect } from 'react';

import ChatHistory from './components/ChatHistory';
import ChatInput from './components/ChatInput';
import Header from './components/Header';

import {connect, sendMsg} from './api'


function App() {

  // 1. init chat history
  const [chatHistory, setChatHistory] = useState([]);

  // 2. connect
  useEffect(() => {

    // pass in the callback function when you call connect
    connect((msg) => {

      // get message from sockets and add to history
      setChatHistory(prevChatHistory => [...prevChatHistory, msg]);

    });
  }, []);

  return (
    <div className="App">
      <Header/>
      <ChatInput sendFunc={sendMsg}/>
      <ChatHistory chatHistory={chatHistory}/>
    </div>
  );
}

export default App;
