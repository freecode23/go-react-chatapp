import React, { useState, useEffect } from 'react';
import { fetchChatroomNames } from '../restApi';

function PickRoom({ onSelectRoom }) {
  const [rooms, setRooms] = useState([]);
  const [newRoomName, setNewRoomName] = useState('');

  useEffect(() => {
    // 1. first fetch all existing rooms
    async function fetchRooms() {
      const roomNames = await fetchChatroomNames();
      console.log(roomNames)
      setRooms(roomNames);
    }

    fetchRooms();
  }, []);

  const handleCreateNewRoom = async () => {
    if (newRoomName.trim() === '') return;  // Prevent empty room names

    // Simply adding it to our local state
    setRooms(prevRooms => [...prevRooms, newRoomName]);

    // Clear the input
    setNewRoomName('');

    // Optionally navigate to the new room
    onSelectRoom(newRoomName);
  };

  return (
    <div>
      <div>
        <input 
          type="text" 
          placeholder="Enter new room name..." 
          value={newRoomName} 
          onChange={e => setNewRoomName(e.target.value)}
        />
        <button onClick={handleCreateNewRoom}>Create Room</button>
      </div>
      
      <div>
        {rooms && rooms.map(room => (
          <button key={room} onClick={() => onSelectRoom(room)}>
            {room}
          </button>
        ))}
      </div>
    </div>
  );
}

export default PickRoom;