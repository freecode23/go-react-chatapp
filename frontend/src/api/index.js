var socket = new WebSocket('ws://localhost:9000/ws')


// init connect function
let connect = (cb) => {
    console.log("connecting")

    socket.onopen =  () => {
        console.log("successfully connected")
    }

    socket.onmessage = (msg) => {
        console.log("socket message from webs:", msg)
    }

    socket.onclose = (event) => {
        console.log("socket closed connection:", event)
    }

    socket.onerror = (err) => {
        console.log("socket error:", err)
    }
}

let sendMsg = (msg) => {
    console.log("socket sending message:",msg)
    socket.send(msg)
}

export {connect, sendMsg}