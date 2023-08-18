import React from 'react';
import './ChatHistory.scss';
import Message from '../Message'

function ChatHistory(props) {

    // 1. init the jsx
    // chat history is just an array of strings
    const messagesJsx = props.chatHistory.map(
        (singleChat, index) => 
        <Message key={index} 
        userName={singleChat.userName} 
        message={singleChat.body} />
    );

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messagesJsx}
        </div>
    );
}


export default ChatHistory