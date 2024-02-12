import { checkCookies, getCookie } from "./cookiesCheck.js";

let statusFriend; // Переменная объявлена в глобальной области видимости

window.onload = async () => {
    await checkFriendship(); // Ждем завершения выполнения асинхронной функции
    
    // После завершения выполнения асинхронной функции, обновляем HTML элемент
    document.getElementById("addFriendButton").innerHTML = statusFriend;
    
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
  await checkFriendship();
  const URL = window.location.href;
  const data = new FormData();
  data.append("userid", getCookie("userid"));
  data.append(
    "friendid",
    window.location.href.split("/")[window.location.href.split("/").length - 1]
  );

  switch (statusFriend) {

    case "Добавить в друзья":
      await fetch("/add-friend", {
        method: "POST",
        body: data,
      })
        .then(async (response) => {
          if (response.status === 200) {
            statusFriend = "Отменить заявку";
            document.getElementById("addFriendButton").disabled = true;
            document.getElementById("addFriendButton").innerHTML = statusFriend;
            return;
          }
        })
        .catch((error) => {
          alert(error);
        });
      break;

    case "Удалить из друзей":
      await fetch("/delete-friend", {
        method: "POST",
        body: data,
      })
        .then(async (response) => {
          if (response.status === 200) {
            statusFriend = "Добавить в друзья";
            document.getElementById("addFriendButton").disabled = false;
            document.getElementById("addFriendButton").innerHTML = statusFriend;
            return;
          }
        })
        .catch((error) => {
          alert(error);
        });
      break;

    case "Принять заявку в друзья":
      await fetch("/accept-friend", {
        method: "POST",
        body: data,
      })
        .then(async (response) => {
          if (response.status === 200) {
            statusFriend = "Удалить из друзей";
            document.getElementById("addFriendButton").disabled = false;
            document.getElementById("addFriendButton").innerHTML = statusFriend;
            return;
          }
        })
        .catch((error) => {
          alert(error);
        });
      break;

    case "Отменить заявку в друзья":
      await fetch("/decliene-friend", {
        method: "POST",
        body: data,
      })
        .then(async (response) => {
          if (response.status === 200) {
            statusFriend = "Добавить в друзья";
            document.getElementById("addFriendButton").disabled = false;
            document.getElementById("addFriendButton").innerHTML = statusFriend;
            return;
          }
        })
        .catch((error) => {
          alert(error);
        });
      break;
  }
};

async function checkFriendship() {
  var userid = getCookie("userid");
  var friendid =
    window.location.href.split("/")[window.location.href.split("/").length - 1];
  if (userid === friendid) {
    document.getElementById("addFriendButton").style.display = "none";
    return statusFriend;
  }
  const data = new FormData();
  data.append("userid", userid);
  data.append("friendid", friendid);
  await fetch("/check-friendship", {
    method: "POST",
    body: data,
  })
    .then(async (response) => {
      console.log(response.status);
      const text = await response.text();
      if (text === "delete friend") {
        statusFriend = "Удалить из друзей";
        return statusFriend;
      }
      if (text === "send invitation") {
        statusFriend = "Добавить в друзья";
        return statusFriend;
      }
      if (text === "accept invitation") {
        statusFriend = "Принять заявку в друзья";
        return;
      }
    })
    .catch((error) => {
      alert(error);
    });
};

