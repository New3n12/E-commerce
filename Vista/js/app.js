
const headerMenu = document.querySelector('.hm-header');
const headerlogo = document.querySelector('.hm-logo');

window.addEventListener('scroll', () => {
    if (window.pageYOffset > 20) {
        headerMenu.classList.add('header-fixed');
        headerlogo.classList.add('header-fixed');
        
    } else {
        headerMenu.classList.remove('header-fixed');
        headerlogo.classList.remove('header-fixed');
    }
})

/*=========================================
    Tabs
==========================================*/
if (document.querySelector('.hm-tabs')) {

    const tabLinks = document.querySelectorAll('.hm-tab-link');
    const tabsContent = document.querySelectorAll('.tabs-content');

    tabLinks[0].classList.add('active');

    if (document.querySelector('.tabs-content')) {
        tabsContent[0].classList.add('tab-active');
    }


    for (let i = 0; i < tabLinks.length; i++) {

        tabLinks[i].addEventListener('click', () => {


            tabLinks.forEach((tab) => tab.classList.remove('active'));
            tabLinks[i].classList.add('active');

            tabsContent.forEach((tabCont) => tabCont.classList.remove('tab-active'));
            tabsContent[i].classList.add('tab-active');

        });

    }

}

/*=========================================
    MENU
==========================================*/

const menu = document.querySelector('.icon-menu');
const menuClose = document.querySelector('.cerrar-menu');

menu.addEventListener('click', () => {
    document.querySelector('.header-menu-movil').classList.add('active');
})

menuClose.addEventListener('click', () => {
    document.querySelector('.header-menu-movil').classList.remove('active');
})

/**
 * 
 */
const enlaces = document.querySelectorAll('.enlace');

enlaces.forEach(enlace => {
    enlace.addEventListener('click', (event) => {
        event.preventDefault();

        // Remover la clase 'enlace' de todos los enlaces
        enlaces.forEach(enlace => enlace.classList.remove('enlace'));

        // Agregar la clase 'enlace' solo al enlace seleccionado
        event.target.classList.add('enlace');
    });
});


