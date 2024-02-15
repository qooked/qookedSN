import { checkCookies, getCookie } from "./cookiesCheck.js";

window.onload = async () => {
  await checkCookies().then((res) => {
    if (!res) {
      window.location.href = `/login`;
      return;
    }
  });
  await getFriendList();
};

async function getFriendList() {
  const URL = window.location.href;
  const data = new FormData();
  var id =
    window.location.href.split("/")[window.location.href.split("/").length - 2];
  data.append("id", id);
  const response = await fetch(URL, {
    method: "POST",
    body: data,
  });
  try {
    const jsonData = await response.json();
    const friendMap = jsonData.friendIDs;

    var friendList = document.getElementById("friendList");
    var keys = Object.keys(friendMap);
    var header3 = document.querySelector("header h3");
    header3.textContent = jsonData.Name + " " + jsonData.Surname;

    for (var i = 0; i < keys.length; i++) {
      var listItem = document.createElement("li");
      var link = document.createElement("a");

      link.href = "/" + keys[i];
      link.textContent = friendMap[keys[i]];

      listItem.appendChild(link);
      friendList.appendChild(listItem);
    }
  } catch (err) {
    console.log(err);
  }
}
