import axios from 'axios';
const BASE_URL = 'http://localhost:8080';
  
// 1. fetch all the history
async function fetchChatHistory(roomName) {
    try {
        const response = await axios.get(`${BASE_URL}/chatHistory/${roomName}`);

        // return a list of json object of {
            // body: item.body, 
            // userName: item.userName
        // }
        // console.log("rest: response.data:", response.data)
        return response.data.map(item => ({
            body: item.body,
            userName: item.userName
        }));
    } catch (error) {
        console.error('Failed to fetch chat history:', error);
    }
}

export default fetchChatHistory;