function validatePassword() {
  const password1 = document.getElementById("password1").value;
  const password2 = document.getElementById("password2").value;
  const email = document.getElementById("email").value;
  const regularExpression =
    /^(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,16}$/;
  const pattern =
    /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
  const minNumberofChars = 8;
  const maxNumberofChars = 16;
  if (!pattern.test(email)) {
    alert("Email введен не верно");
    return false;
  }
  if (
    password1.length < minNumberofChars ||
    password1.length > maxNumberofChars
  ) {
    alert("Длина пароля должная быть от 8 до 16 символов");
    return false;
  }
  if (!regularExpression.test(password1)) {
    alert("Пароль должен включать хотя бы одну цифру и специальный символ.");
    return false;
  }
  if (password1 != password2) {
    alert("Пароли не совпадают");
    return false;
  }
  return true;
}

const form = document.getElementById("registrationForm");

form.onsubmit = async (e) => {
  e.preventDefault();
  if (!validatePassword()) {
    return;
  }
  const fd = new FormData(form);
  fd.delete("password2");
  console.log(form);
  await fetch("/register", {
    method: "POST",
    body: fd,
  }).then(async (response) => {
    console.log(response.status);
    if (response.status === 200) {
      window.location.href = "/login";
      return;
    }
    if (response.status === 400) {
      alert("Такой пользователь уже существует");
      return;
    }
  });
};
