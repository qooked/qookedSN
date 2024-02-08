export function getCookie(cookieName) {
  let cookie = {};
  document.cookie.split(";").forEach(function (el) {
    let [key, value] = el.split("=");
    cookie[key.trim()] = value;
  });
  return cookie[cookieName];
}

export async function checkCookies() {
  const body = new FormData();
  body.append("accessToken", getCookie("accessToken"));
  body.append("refreshToken", getCookie("refreshToken"));
  body.append("userid", getCookie("userid"));
  console.log(body);
  return await fetch("/compare-tokens", { method: "POST", body })
    .then(async (response) => {
      console.log(response);
      return response.status === 200;
    })
    .catch((error) => alert(error));
}
