$(document).ready(function (params) {
    activeMenu()
})
function activeMenu(params) {
    var pathname = window.location.pathname; // Returns path only (/path/example.html)
    var url      = window.location.href;     // Returns full URL (https://example.com/path/example.html)
    var origin   = window.location.origin;  
    $('#container-menu .nav-item').each(function(e,params2) {
        
        if ($(this).children().length > 1) {
            $(this).find('div').children().each(function(e,params) {
                if ($(params).attr('href') == url) {
                    $(params).addClass('active')
                    $(params2).addClass('show')
                    $(params2).find('div').addClass('show')
                }
            })
        }else{
            if ($(this).find('a').attr('href') == url) {
                $(this).find('a').addClass('active')
            }
        }
    })

}