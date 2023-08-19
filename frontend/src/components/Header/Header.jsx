import React,{ useContext }  from 'react';
import './Header.scss';
import UserContext from '../../utils/UserContext';

function Header() {
    const { userName, roomName} = useContext(UserContext);
    return (
        <div className='header'>
            <h2>Gochat: {userName}@{roomName}</h2>
        </div>
    )
}

export default Header