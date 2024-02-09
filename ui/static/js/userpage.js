import { checkCookies, getCookie } from "./cookiesCheck.js";

window.onload = () => {
  checkCookies().then((res) => {
    if (!res) {
      window.location.href = `/login`;
      return;
    }
  });
};

function deleteAllCookies() {
  const cookies = document.cookie.split(";");
  for (let i = 0; i < cookies.length; i++) {
    const cookie = cookies[i];
    const eqPos = cookie.indexOf("=");
    const name = eqPos > -1 ? cookie.substr(0, eqPos) : cookie;
    document.cookie = name + "=;expires=Thu, 01 Jan 1970 00:00:00 GMT";
  }
}

logoutButton.onclick = async () => {
  const cookies = new FormData();
  cookies.append("accessToken", getCookie("accessToken"));
  cookies.append("refreshToken", getCookie("refreshToken"));
  cookies.append("userid", getCookie("userid"));
  await fetch("/logout", {
    method: "POST",
    body: cookies,
  })
    .then(async (response) => {
      if (response.status == 200) {
        console.log(response.status);
        deleteAllCookies();
        window.location.href = `/login`;
        return;
      }
      if (response.status == 400) {
        deleteAllCookies();
        window.location.href = `/login`;
        return;
      }
    })
    .catch((error) => {
      alert(error);
    });
};
