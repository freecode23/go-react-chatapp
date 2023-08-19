import React, { useState, useEffect } from 'react';

import ChatHistory from '../components/ChatHistory';
import ChatInput from '../components/ChatInput';
import Header from '../components/Header';
import UserContext from '../utils/UserContext';

import {initChatroomSocket} from '../socketApi'
import {fetchChatHistory} from '../restApi';
import generator from '../utils/UniqueNameGenerator'


function ChatRoom(props) {

  // 0. init random name for this user 
  const [userValues, setUserValues] = useState({
    userName: generator.generateUniqueName(),
    roomName: props.roomName
  });
  
  // 1. init chat history of this room and connection to backend
  const [chatHistory, setChatHistory] = useState([]);
  const [chatroomSocket, setChatroomSocket] = useState(null);

  // 2. Will only run once as it has en empty dependency
  // side effect
  useEffect(() => {

    // 1. fetch all history
    async function fetchData() {
      const messageObj = await fetchChatHistory(userValues.roomName);
      setChatHistory(messageObj.reverse());
    }
    fetchData();

    // 2. Connect to socket with call back function
    // that adds message to chat history everytime theres a new message received
    // from any other clients
    const initializedSocket = initChatroomSocket(userValues.roomName, (msgEvent) => {

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

    // 3. set the socket to state
    setChatroomSocket(initializedSocket); 
  }, []);

  return (
    <div className="ChatRoom">
      <UserContext.Provider value={userValues}>
      <Header/>
      <ChatHistory chatHistory={chatHistory}/>
      {chatroomSocket ? <ChatInput sendFunc={chatroomSocket.socketSendMsg}/> : <p>Loading...</p>}
      </UserContext.Provider>
    </div>
  );
}

export default ChatRoom;
