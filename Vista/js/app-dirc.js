document.addEventListener("DOMContentLoaded", function () {
    const form = document.querySelector("form");
    
    form.addEventListener("submit", function (event) {
        event.preventDefault();
        
        if (validateForm()) {
            form.submit();
        }
    });
    
    function validateForm() {
        const firstName = document.getElementById("firstname").value;
        const lastName = document.getElementById("lastname").value;
        const email = document.getElementById("email").value;
        const addressOne = document.getElementById("address-one").value;
        const city = document.getElementById("city").value;
        const state = document.getElementById("state").value;
        const zip = document.getElementById("zip").value;
        const country = document.getElementById("country").value;
        const telefono = document.getElementById("telefono").value;

        if (firstName.trim() === "" || lastName.trim() === "" || email.trim() === "" || 
            addressOne.trim() === "" || city.trim() === "" || state.trim() === "" || 
            zip.trim() === "" || country.trim() === "" || telefono.trim() === "") {
            alert("Por favor, complete todos los campos obligatorios.");
            return false;
        }

        if (!validateEmail(email)) {
            alert("Ingrese una dirección de correo electrónico válida.");
            return false;
        }

        if (!validatePhoneNumber(telefono)) {
            alert("Ingrese un número de teléfono válido.");
            return false;
        }

        return true;
    }

    function validateEmail(email) {
        const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
        return emailPattern.test(email);
    }

    function validatePhoneNumber(phoneNumber) {
        const phoneNumberPattern = /^\d{10,}$/;
        return phoneNumberPattern.test(phoneNumber);
    }
});