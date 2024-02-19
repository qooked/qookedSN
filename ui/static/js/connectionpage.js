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

