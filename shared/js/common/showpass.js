function  ShowHidePass(id) {
    let x = document.getElementById(id);
    if (x.type === "password") {
        x.type = "text";
    } else {
        x.type = "password";
    }
}