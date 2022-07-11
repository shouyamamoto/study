getButton = document.getElementById("getAlbum")
deleteButton = document.getElementById("deleteAlbum")
albums = document.getElementById("albums")

async function getAlbum() {
    const res = await fetch("http://localhost:8888/")
    const data = await res.json()

    while (albums.firstChild) {
        albums.removeChild(albums.firstChild)
    }

    data.forEach(d => {
        item = document.createElement("li")
        item.innerText = `ID: ${d.ID} アルバム名:${d.title} 作者:${d.artist} 価格:${d.price}`
        albums.appendChild(item)
    })
}

async function deleteAlbum() {
    id = document.getElementById("id")  
    await fetch(`http://localhost:8888/albums/delete/${id.value}`,{method: "DELETE"})
    id.value = ""
}

getButton.addEventListener("click", getAlbum)
deleteButton.addEventListener("click", deleteAlbum)
