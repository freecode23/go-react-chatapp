
import React from 'react';
import './Message.scss';


function Message(props) {
    const msg = props.message;

    return (
        <div className='Message'>
            {msg}
        </div>
    );
}

export default Message;
