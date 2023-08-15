import React from 'react';
import './ChatHistory.scss';
import Message from '../Message'

function ChatHistory(props) {

    // 1. init the jsx
    // chat history is just an array of strings
    const messages = props.chatHistory.map(
        (chatHist, index) => <Message key={index} message={chatHist} />
    );

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    );
}


export default ChatHistory