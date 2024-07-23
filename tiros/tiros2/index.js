
function lerJson() {
    let divVencedor = document.querySelector("#vencedor")
    let divUsuarios = document.querySelector("#usuarios")
    let divTitulo = document.querySelector("#titulo")

    fetch("./dados.json").then((response) => {
        response.json().then((dados) => {
            if (dados.usuarios.length > 10) {
                divTitulo.innerHTML = `ATUAL VENCEDOR`
                divVencedor.innerHTML =  `${dados.usuarios[0].nome} - ${dados.usuarios[0].tempoStr}`
            }
            dados.usuarios.map(
                (usuario) => {
                    console.log(usuario)
                    if (usuario.id > 10 ) {
                        divUsuarios.innerHTML += 
                            `<tr class = "letra" >
                                <td> ${usuario.nome} </td>
                                <td> ${usuario.tempoStr} </td>
                                <td> ${usuario.cidade} </td>
                                <td> ${usuario.telefone} </td>
                                <td> <button type="button" onclick="excluir(${usuario.id})">DELETAR</button> </td>
                            </tr>`
                    } else {
                        divUsuarios.innerHTML += 
                        `<tr class = "letra" >
                            <td> ${usuario.nome} </td>
                            <td> ${usuario.tempoStr} </td>
                            <td> ${usuario.cidade} </td>
                            <td> ${usuario.telefone} </td>
                            <td> <button disabled type="button" onclick="excluir(${usuario.id})">DELETAR</button> </td>
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
    const cidade = document.querySelector("#cidade")
    const telefone = document.querySelector("#telefone")

    const uri = "http://localhost:8080/salvar?nome="+ nome.value + "&tempo=" + tempo.value + "&cidade=" + cidade.value + "&telefone=" + telefone.value

    fetch(uri).then(()=> {
        window.location.reload()
    })
}

function excluir(id) {
    const req = new XMLHttpRequest()
    const uri = "http://localhost:8080/deletar?id=" + id
    
    window.location.href = uri
    window.location.href = "http://localhost:8080/"
}
