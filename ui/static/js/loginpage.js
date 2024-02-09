import { checkCookies, getCookie } from "./cookiesCheck.js";
window.onload = () => {
  console.log("aboba");
  checkCookies().then((res) => {
    if (res) {
      let userID = getCookie("userid");
      window.location.href = `/${userID}`;
      return;
    }
  });
};

function validatePassword() {
  const password = document.getElementById("password").value;
  const email = document.getElementById("email").value;
  const regularExpression =
    /^(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,16}$/;
  const pattern =
    /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
  const minNumberofChars = 8;
  const maxNumberofChars = 16;
  if (!pattern.test(email)) {
    alert("Email введен неверно");
    return false;
  }
  if (
    password.length < minNumberofChars ||
    password.length > maxNumberofChars
  ) {
    alert("Длина пароля должная быть от 8 до 16 символов");
    return false;
  }
  if (!regularExpression.test(password)) {
    alert("Пароль должен включать хотя бы одну цифру и специальный символ.");
    return false;
  }
  return true;
}
const form = document.getElementById("loginForm");

form.onsubmit = async (e) => {
  e.preventDefault();
  if (!validatePassword()) {
    return;
  }
  const fd = new FormData(form);
  await fetch("/login", {
    method: "POST",
    body: fd,
  })
    .then(async (response) => {
      console.log(response.status);
      if (response.status === 200) {
        response.json().then((data) => {
          let userid = data.userid;
          let accessToken = data.accessToken;
          let refreshToken = data.refreshToken;
          document.cookie = `accessToken=${accessToken};`;
          document.cookie = `refreshToken=${refreshToken};`;
          document.cookie = `userid=${userid};`;
          window.location.href = "/" + userid;
        });
        return;
      }
      if (response.status === 400) {
        alert("Неверные имя пользователя или пароль");
        return;
      }
    })
    .catch((error) => {
      alert(error);
    });
};
