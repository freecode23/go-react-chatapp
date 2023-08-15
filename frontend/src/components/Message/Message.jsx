
import React from 'react';
import './Message.scss';


function Message(props) {
    const msg = JSON.parse(props.message);

    return (
        <div className='Message'>
            {msg.body}
        </div>
    );
}

export default Message;
