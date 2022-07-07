button = document.getElementById("getUser")
users = document.getElementById("users")

async function getUsers() {
    const res = await fetch("http://localhost:8081/users")
    const data = await res.json()

    data.forEach(u => {
        item = document.createElement("li")
        item.innerText = `${u.name}は${u.age}歳です。`
        users.appendChild(item)
    })
}

button.addEventListener("click", getUsers)