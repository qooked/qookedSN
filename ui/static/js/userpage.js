window.onload = () => {
    checkCookies().then((res) => {
      if (!res) {
        window.location.href = `/login`;
        return;
      }
    });
  };