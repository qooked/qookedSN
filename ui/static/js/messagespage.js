import { checkCookies, getCookie } from "./cookiesCheck.js";

window.onload = async () => {
  await checkCookies().then((res) => {
    if (!res) {
      window.location.href = `/login`;
      return;
    }
  });
};

// Создаем новое соединение WebSocket
