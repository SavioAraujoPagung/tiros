
function lerJson() {
    let divUsuarios = document.querySelector("#usuarios")
    divUsuarios.innerHTML = ``

    fetch("dados.json").then((response) => {
        response.json().then((dados) => {
            dados.usuarios.map(
                (usuario) => {
                    if (usuario.active == true) {
                        divUsuarios.innerHTML += 
                        `<tr>
                            <td> ${usuario.nome} </td>
                            <td> ${usuario.tempo} </td>
                            <td> <button type="button" onclick="excluir(${usuario.id})">DELETAR</button> </td>
                        </tr>`
                    }
                }    
            )
        })
    })
}

function salvar() {
    const nome = document.querySelector("#nome")
    const tempo = document.querySelector("#tempo")

    const req = new XMLHttpRequest()
    const uri = "http://localhost:1414/save?nome="+ nome.value + "&tempo=" + tempo.value

    req.open("GET", uri, false)
    req.onreadystatechange = function() {
        if (req.readyState === 4 && req.status === 200) {
          const response = JSON.parse(req.responseText);
          // faça algo com a resposta
        }
    }
    req.send()
}

function excluir(id) {
    const req = new XMLHttpRequest()
    const uri = "http://localhost:1414/delete?id=" + id

    req.open("GET", uri, false)
    req.onreadystatechange = function() {
        if (req.readyState === 4 && req.status === 200) {
          const response = JSON.parse(req.responseText);
          // faça algo com a resposta
        }
    }
    req.send()
}
