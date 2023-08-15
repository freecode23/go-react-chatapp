import axios from 'axios';
const BASE_URL = 'http://localhost:8080';
  
// 1. fetch all the history
async function fetchChatHistory() {
    try {
        const response = await axios.get(`${BASE_URL}/chatHistory`);
        
        // Assuming the response data is an array as inferred from the original code
        return response.data.map(item => item.body);
    } catch (error) {
        console.error('Failed to fetch chat history:', error);
    }
}

export default fetchChatHistory;