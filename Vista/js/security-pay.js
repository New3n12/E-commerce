function securityPay() {
    const types = document.querySelectorAll('.type');
    var animate = document.getElementById("loadpay");
    console.log(animate);
    types.forEach(type => {
        console.log(type.getAttribute("id"));
        if (type.classList.contains('selected')){
            animate.style.visibility = 'visible';
            $.ajax({
            url: "/pasarela/buy/paypal",
            type: "POST",
            data: "tipo="+type.getAttribute("id"),
            success: function (response) {
                animate.style.visibility = 'hidden';
                if (response != "Err"){
                    window.location.href ='https://www.sandbox.paypal.com/checkoutnow?token='+response
                }
            },  error: function (xhr, status, error) {
                console.log(error);
                animate.style.visibility = 'hidden';
            }
           });
        }
    });
}