
function lerJson() {
    let divVencedor = document.querySelector("#vencedor")
    let divUsuarios = document.querySelector("#usuarios")
    divUsuarios.innerHTML = ``

    fetch("./dados.json").then((response) => {
        response.json().then((dados) => {
            divVencedor.innerHTML =  `<h2> Nome: ${dados.usuarios[0].nome} | Tempo: ${dados.usuarios[0].tempoStr} </h2>`
            dados.usuarios.map(
                (usuario) => {
                        divUsuarios.innerHTML += 
                        `<tr>
                            <td> ${usuario.nome} </td>
                            <td> ${usuario.tempoStr} </td>
                            <td> ${usuario.cidade} </td>
                            <td> ${usuario.telefone} </td>
                            <td> <button type="button" onclick="excluir(${usuario.id})">DELETAR</button> </td>
                        </tr>`
                }    
            )
        })
    })
}

function salvar() {
    const nome = document.querySelector("#nome")
    const tempo = document.querySelector("#tempo")
    const cidade = document.querySelector("#cidade")
    const telefone = document.querySelector("#telefone")

    const uri = "http://localhost:8080/salvar?nome="+ nome.value + "&tempo=" + tempo.value + "&cidade=" + cidade.value + "&telefone=" + telefone.value


    window.location.href = uri
    window.location.href = "http://localhost:8080/"
}

function excluir(id) {
    const req = new XMLHttpRequest()
    const uri = "http://localhost:8080/deletar?id=" + id
    
    window.location.href = uri
    window.location.href = "http://localhost:8080/"
}
