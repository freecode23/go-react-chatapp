import React, { useState, useEffect } from 'react';

import ChatHistory from './components/ChatHistory';
import ChatInput from './components/ChatInput';
import Header from './components/Header';
import UserContext from './utils/UserContext';

import {socketConnect, sendMsg} from './socketApi'
import fetchChatHistory from './restApi';
import generator from './utils/UniqueNameGenerator'


function App() {

  // 0. init random name for this user 
  const [userValues, setUserValues] = useState({
    userName: generator.generateUniqueName(),
    roomName: "roomX"
  });

  // 1. init chat history of this room
  const [chatHistory, setChatHistory] = useState([]);

  // 2. every render
  useEffect(() => {

    // 1.  fetch all history
    async function fetchData() {
      const messageObj = await fetchChatHistory(userValues.roomName);
      setChatHistory(messageObj.reverse());
    }
    fetchData();

    // 2. connect to socket with call back function
    // that adds message to chat history everytime theres a new message received
    // from any other clients
    socketConnect((msgEvent) => {

      // contains body and userName and type
      const jsonData = JSON.parse(msgEvent.data) 

      // get message from sockets and add to history if there's a new incoming message
      const chatUserBody = {
        roomName: jsonData.roomName,
        userName: jsonData.userName,
        body: jsonData.body,
      }
      setChatHistory(prevChatHistory => [...prevChatHistory, chatUserBody]);

    });
  }, []);

  return (
    <div className="App">
      <UserContext.Provider value={userValues}>
      <Header/>
      <ChatHistory chatHistory={chatHistory}/>
      <ChatInput sendFunc={sendMsg}/>
      </UserContext.Provider>
    </div>
  );
}

export default App;
