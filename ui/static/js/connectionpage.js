import { checkCookies, getCookie } from "./cookiesCheck.js";

window.onload = async () => {
  connectToServer();
  await checkCookies().then((res) => {
    if (!res) {
      window.location.href = `/login`;
      return;
    }
  });
};

const socket = new WebSocket('ws://localhost:8080/ws');



// Функция для обработки входящих сообщений от сервера
socket.onmessage = function(event) {
    const message = event.data;
    displayMessage(message);
};


// Функция для отправки сообщения на сервер
function sendMessage() {
    const messageInput = document.getElementById('messageInput');
    const message = messageInput.value;
    socket.send(message);
    displayMessage('You: ' + message);
    messageInput.value = ''; // Очищаем поле ввода
}

// Функция для отображения сообщений в интерфейсе
function displayMessage(message) {
    const messagesDiv = document.getElementById('messages');
    const messageElement = document.createElement('div');
    messageElement.textContent = message;
    messagesDiv.appendChild(messageElement);
}

// Обработчик нажатия кнопки отправки сообщения
document.getElementById('sendButton').addEventListener('click', sendMessage);
