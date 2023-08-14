import React from 'react';
import './ChatInput.scss';


function ChatInput(props) {

    function handleKeyDown(event) {
        if (event.key === 'Enter') {
            const message = event.target.value;
            props.sendFunc(message);
        }
    }
    return (
        <div className='chatInput'>
            <input onKeyDown={handleKeyDown} placeholder="Type a message... Hit enter to send"/> 
        </div>
    )
}

export default ChatInput;
