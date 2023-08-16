import React,{ useContext, useState } from 'react';
import UserContext from '../../utils/UserContext';
import './ChatInput.scss';

function ChatInput(props) {
    const userName = useContext(UserContext);
    const [messageValue, setMessageValue] = useState('');

    function handleKeyDown(event) {
        // 1. if user press enter and its not an empty message after trim
        if (event.key === 'Enter'&& messageValue.trim()) {

            // - get the message body
            const messageBody = event.target.value;

            // - init message object
            const messageObj = {
                userName: userName,
                body: messageBody
            };

            // - send to socket backend
            props.sendFunc(messageObj);

            // - reset to nempty after sending
            setMessageValue('');
        }
    }

    function handleChange(event) {
        // Update the state when the input value changes
        setMessageValue(event.target.value); 
    }

    return (
        <div className='chatInput'>
            <input 
            value={messageValue} 
            onKeyDown={handleKeyDown} 
            onChange={handleChange} 
            placeholder={`${userName}: Type a message... Hit enter to send`}/> 
        </div>
    )
}

export default ChatInput;
