$(document).ready(function() {
    var url = GetQueryString("url");
    if (url != null) {
        $('#text').val(url);
        testCrawler(url);
    }
    $('#submit').click(function() {
        var url = $('#text').val().trim();
        if (url != "") {
            testCrawler(url);
        }
    });
});

function testCrawler(url) {
    $('#editor_holder').html("<h4>loading...</h4>");
    $.ajax({
        url: "/api/?url="+url, cache: false,
        success: function(result) {
            $('#editor_holder').jsonview(result);
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.responseText);
        }
    });
}
