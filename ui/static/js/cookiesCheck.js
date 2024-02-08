function getCookie(cookieName) {
  let cookie = {};
  document.cookie.split(";").forEach(function (el) {
    let [key, value] = el.split("=");
    cookie[key.trim()] = value;
  });
  return cookie[cookieName];
}

async function checkCookies() {
  const body = new FormData();
  body.append("accessToken", getCookie("accessToken"));
  body.append("refreshToken", getCookie("refreshToken"));
  body.append("userid", getCookie("userid"));
  return await fetch("/compare-tokens", { method: "GET", body })
    .then(async (response) => {
      return response.status === 200;
    })
    .catch((error) => false);
}
