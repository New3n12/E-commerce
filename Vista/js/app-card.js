/* card */
const cartCountElement = document.getElementById('cart-count');
// Actualizar el contador del carrito
function actualizarContadorCarrito(nuevoContador) {
    cartCountElement.textContent = nuevoContador;
}

// Ejemplo: aumentar el contador del carrito en 1
function aumentarContadorCarrito() {
    const contadorActual = parseInt(cartCountElement.textContent);
    const nuevoContador = contadorActual + 1;
    Swal.fire({
        position: 'center',
        icon: 'success',
        title: 'Se Agrego Al Carrito',
        showConfirmButton: false,
        timer: 1500
    }).then(function () {
        actualizarContadorCarrito(nuevoContador);
    });
}

function addCar(id_prod, id_usu, cant) {
    if (id_usu != 0) {
        const data = "caso=0&id=" + id_prod + "&idusu=" + id_usu + "&cant=" + cant;
        $.ajax({
            url: "/operaciones/card/produc",
            type: "POST",
            data: data,
            dataType: 'json',
            success: function (response) {
                var produc = response;
                var nuevoElemento = document.createElement('div');
                nuevoElemento.setAttribute('data-card-id', produc.ID_card);
                nuevoElemento.classList.add('modal__item');
                nuevoElemento.style.margin = "15px";

                // Agrega el contenido del producto al nuevo elemento
                nuevoElemento.innerHTML = `
                        <div class="modal__thumb">
                            <img src="data:image/jpeg;base64,${produc.Img_card}" class="img-card" alt="">
                        </div>
                        <div class="modal__text-product">
                            <p class="desc">${produc.Nombre_card}</p>
                            <p class="desc"><strong>$${produc.Precio_card}</strong></p>
                        </div>
                        <div class="quant">
                            <span onclick="moreCard('${produc.ID_card}')" class="plus">+</span>
                            <span id="${produc.ID_card}" class="qt">${produc.Cantidad_card}</span>
                            <span onclick="lessCard('${produc.ID_card}')" class="minus">-</span>
                        </div>
                    `;

                // Agrega el nuevo elemento a la lista del carrito
                var listaCarrito = document.getElementById('listaCarrito');
                var label = document.createElement("label");
                label.setAttribute("name", produc.ID_card);
                label.style.color = " red";
                label.style.visibility = " hidden";
                label.style.paddingLeft = " 62px";
                label.style.fontSize = "15px";
                label.innerText = "Cantidad Maxima Selecionada";
                listaCarrito.appendChild(label);
                listaCarrito.appendChild(nuevoElemento);
                aumentarContadorCarrito();
                // console.log(response);
            },
            error: function (xhr, status, error) {
                Swal.fire({
                    title: 'El Articulo ya esta en el Carrito!!',
                    text: '¿Desea abrir el Carrito?',
                    icon: 'question',
                    showCancelButton: true,
                    confirmButtonText: 'Sí',
                    cancelButtonText: 'No'
                }).then((result) => {
                    if (result.isConfirmed) {
                        const modal = document.getElementById("jsModalCarrito");
                        if (modal !== null) {
                            modal.classList.add('active');
                        } else {
                            console.error('El modal es null');
                        }
                    } else if (result.isDenied) {

                    }
                });

            }
        });
    } else {

        Swal.fire({
            icon: 'error',
            title: 'Oops...',
            text: 'Debe estar registrado para agregar al carrito!',
            timer: 2000,
            showConfirmButton: false
        }).then(() => {
            window.location.href = "/login";
        });

    }

}
