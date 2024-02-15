import { checkCookies, getCookie } from "./cookiesCheck.js";

window.onload = async () => {
  await checkFriendStatus();
  await checkCookies().then((res) => {
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

addFriendButton.onclick = async () => {
  const URL = window.location.href;
  const data = new FormData();
  data.append("userid", getCookie("userid"));
  data.append(
    "friendid",
    window.location.href.split("/")[window.location.href.split("/").length - 1]
  );
  await fetch("/change-friend-status", {
    method: "POST",
    body: data,
  })
    .then(async (response) => {
      if (response.status == 200) {
        document.getElementById("addFriendButton").innerHTML =
          await response.text();
        return;
      }
      if (response.status == 500) {
        alert();
        return;
      }
    })
    .catch((error) => alert(error));
};

async function checkFriendStatus() {
  if (
    getCookie("userid") ===
    window.location.href.split("/")[window.location.href.split("/").length - 1]
  ) {
    document.getElementById("addFriendButton").style.display = "none";
    return;
  }
  const URL = window.location.href;
  const data = new FormData();
  data.append("userid", getCookie("userid"));
  data.append(
    "friendid",
    window.location.href.split("/")[window.location.href.split("/").length - 1]
  );
  await fetch("/check-friend-status", {
    method: "POST",
    body: data,
  })
    .then(async (response) => {
      if (response.status == 200) {
        document.getElementById("addFriendButton").innerHTML =
          await response.text();
        return;
      }
      if (response.status == 400) {
        alert();
        return;
      }
    })
    .catch((error) => alert(error));
}

friendListButton.onclick = () => {
  var id =
    window.location.href.split("/")[window.location.href.split("/").length - 1];
  window.location.href = "/" + id + "/firend-list";
};
