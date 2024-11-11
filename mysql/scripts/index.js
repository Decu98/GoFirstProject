
$(window).on('load', function () {
    $("#BPrimary").on("click", function () {
        $("#TData").removeClass("invisible");
        $.get( "/SQLIndex", function( data ) {
            var obj = JSON.parse(data)
            var rowsToAppend = ""
            for(var i = 0; i < obj.length; i++){
                var ID = obj[i].ID
                var NAME = obj[i].NAME
                var rowToAppend = `<tr><th scope="row">${ID}</th><td>${NAME}</td></tr>`
                rowsToAppend += rowToAppend;
            }
            $("#TBData").html(rowsToAppend)
        })
    })
});