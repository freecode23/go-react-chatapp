import React, { useState } from 'react';
import PickRoom from './page/PickRoom';
import ChatRoom from './page/ChatRoom'; 

function App() {
  const [selectedRoom, setSelectedRoom] = useState(null);

  // if no seletected room will go here
  if (!selectedRoom) {
    return <PickRoom onSelectRoom={setSelectedRoom} />;
  }
  
  return <ChatRoom roomName={selectedRoom} />;
}
export default App;

