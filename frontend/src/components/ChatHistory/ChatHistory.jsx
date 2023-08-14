import React from 'react';
import './ChatHistory.scss';
import Message from '../Message'

function ChatHistory(props) {

    // 1. init the jsx
    const messages = props.chatHistory.map(
        chatHist => <Message key={chatHist.timeStamp} message={chatHist.data} />
    );

    return (
        <div className="ChatHistory">
            <h2>Chat History</h2>
            {messages}
        </div>
    );
}


export default ChatHistory