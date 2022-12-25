var slide_active = 0;
$(document).ready(function (params) {
    setInterval(function(params) {
        slide_active +=1
        if (slide_active == 2) {
            slide_active = 0
        }
        $('.carousel-inner').find('.active').removeClass('active')
        $('.carousel-inner').children().eq(slide_active).addClass('active')
        console.log(slide_active);
    },3000)
    activeMenu()
})
function activeMenu(params) {
    var pathname = window.location.pathname; // Returns path only (/path/example.html)
    var url      = window.location.href;     // Returns full URL (https://example.com/path/example.html)
    var origin   = window.location.origin;  
    $('#container-menu .nav-item').each(function(e,params2) {
        
        if ($(this).children().length > 1) {
            $(this).find('div').children().each(function(e,params) {
                console.log(params);
                if ($(params).attr('href') == url) {
                    $(params).addClass('active')
                    $(params2).addClass('show')
                    $(params2).find('div').addClass('show')
                }
            })
        }else{
            if ($(this).find('a').attr('href') == url) {
                $(this).find('a').addClass('active')
                $(this).find('a').attr('style','background-color: #428fe1; color:white;')

            }
        }
    })

}