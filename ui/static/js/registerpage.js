function validatePassword() {
    const password1 = document.getElementById('password1').value;
    const password2 = document.getElementById('password2').value;
    const regularExpression = /^(?=.*[0-9])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,16}$/;
    const pattern  = /^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$/i;
    const name1 = document.getElementById('name').value;
    const surname = document.getElementById('surname').value;
    const email = document.getElementById('email').value;
    const minNumberofChars = 8;
    const maxNumberofChars = 16;
    if(!pattern.test(email)) {
        alert("Email введен не верно");
        return false;
    }
    if (name1 === '' || surname === '' || password1 === '' || password2 === '' || email === ''){
        alert("Заполните пустые поля");
        return false;
    }
    if(password1 != password2) {
        alert("Пароли не совпадают");
        return false;
    }
    if(password1.length < minNumberofChars || password1.length > maxNumberofChars){
        alert("Длина пароля должная быть от 8 до 16 символов");
        return false;
    }
    if(!regularExpression.test(password1)) {
        alert("Пароль должен включать хотя бы одну цифру и специальный символ.");
        return false;
    }
    return true;
}
function submitForm() {
    var form = document.getElementById("registrationForm");
    var formData = new FormData(form);

    fetch("/register", {
        method: "POST",
        body: formData
    })
    .then(response => response.json())
    .then(data => {
        console.log(data); // Тут можешь обработать ответ от сервера
    })
    .catch(error => {
        console.error("Error:", error);
    });
}

function checker(){
    if (validatePassword() == false){
    }else{
        window.location.href = "/login";
    }
}

button.addEventListener('click', checker);