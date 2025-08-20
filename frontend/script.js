document.addEventListener("DOMContentLoaded", () => {

    // Elements of first view

    const viewLogin = document.getElementById("view-login");
    const viewDashboard = document.getElementById("view-dashboard");
    const btnLogin = document.getElementById("btn-login");
    const btnRegister = document.getElementById("btn-register");
    const btnLogout = document.getElementById("btn-logout");
    
    // Elements of second view

    const usernameSpan = document.getElementById("username-span");
    const avatarImage = document.getElementById("user-avatar");

    // Logic of the app

    if (localStorage.getItem("loggedUser")) {
        showDashboard(localStorage.getItem("loggedUser"));
        window.dispatchEvent(new Event('app:login'));
    }

    btnLogin.addEventListener("click", async () => {
        const username = document.getElementById("username").value.trim();
        const password = document.getElementById("password").value.trim();

        if (!username || !password) {
            alert("Please enter username and password");
            return;
        }

        try {
            const response = await fetch("http://localhost:8080/usuarios/iniciar-sesion", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ username, password })
            });

            if (!response.ok) {
                alert("User or password incorrect");
                return;
            }

            const data = await response.json();
            saveUserSession(data);
            showDashboard(data.usuario.username);
            window.dispatchEvent(new CustomEvent('app:login', { detail: { user: data.usuario } }));

        } catch (error) {
            console.error("Login error:", error);
            alert("Could not connect to the server");
        }
    });

    btnRegister.addEventListener("click", async () => {
        const username = document.getElementById("username").value.trim();
        const password = document.getElementById("password").value.trim();

        if (!username || !password) {
            alert("Please enter username and password");
            return;
        }

        try {
            const response = await fetch("http://localhost:8080/usuarios", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ username, password })
            });

            if (!response.ok) {
                alert("Can't register because the user already exits");
                return;
            }

            alert("User registered successfully");

            const responseLogin = await fetch("http://localhost:8080/usuarios/iniciar-sesion", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify({ username, password })
            });

            const dataLogin = await responseLogin.json();
            saveUserSession(dataLogin);
            showDashboard(dataLogin.usuario.username);

            window.dispatchEvent(new CustomEvent('app:login', { detail: { user: dataLogin.usuario } }));

        } catch (error) {
            console.error("Register error:", error);
            alert("Could not connect to the server");
        }

    });

    btnLogout.addEventListener("click", () => {
        localStorage.removeItem("loggedUser");
        localStorage.removeItem("authToken");
        localStorage.removeItem("userId");
        localStorage.removeItem("avatarUrl");

        viewDashboard.style.display = "none";;
        viewLogin.style.display = "flex";

        document.getElementById("username").value = "";
        document.getElementById("password").value = "";

        window.dispatchEvent(new Event('app:logout'));
    });

    function saveUserSession(data) {
        if (!data || !data.usuario || !data.token) {
            console.error("Invalid user data:", data);
            return;
        }

        const { id, username, avatarPath } = data.usuario;
        const token = data.token;

        localStorage.setItem("authToken", token);
        localStorage.setItem("loggedUser", username);
        localStorage.setItem("userId", id);
        localStorage.setItem("avatarUrl", `http://localhost:8080${avatarPath}`);
    }

    function showDashboard(username) {
        usernameSpan.textContent = username;

        const avatarUrl = localStorage.getItem("avatarUrl");
        if (avatarUrl) {
            avatarImage.src = avatarUrl;
        } else {
            avatarImage.src = "";
            avatarImage.classList.add("fallback");
            avatarImage.textContent = "👤";
        }

        viewLogin.style.display = "none";
        viewDashboard.style.display = "block";
    }

});