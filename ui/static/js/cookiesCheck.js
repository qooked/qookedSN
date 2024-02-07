window.onload = () => console.log(document.cookie);
function getCookie(cookieName) {
  let cookie = {};
  document.cookie.split(";").forEach(function (el) {
    let [key, value] = el.split("=");
    cookie[key.trim()] = value;
  });
  return cookie[cookieName];
}
e.preventDefault();
if (!validatePassword()) {
  return;
}
await fetch("/login", {
  method: "POST",
  body: fd,
})
  .then(async (response) => {
    console.log(response.status);
    if (response.status === 200) {
      response.json().then((data) => {
        let access = true;
        console.log(access);
      });
      return;
    }
    if (response.status === 400) {
      alert("Неверные имя пользователя или пароль");
      return;
    }
  })
  .catch((error) => {
    console.log(error);
  });
