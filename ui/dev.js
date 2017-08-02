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
            visual(result)
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            alert(XMLHttpRequest.responseText);
        }
    });
}

function one(k, v) {
    return "<fieldset><legend>"+k+"</legend>"+v+"</fieldset>";
}

function visual(doc) {
    var html = "";
    for (k in doc) {
        if (doc[k] == "") {
            continue;
        }
        var v = doc[k];
        if (k == "url" || k == "canonical_url") {
            v = "<a href=\""+v+"\">"+v+"</a>";
        } else if (k == "text") {
            v = "<pre>"+v+"</pre>";
        } else if (k == "favicon") {
            v = "<img src=\""+v+"\" />";
        } else if (k == "images") {
            var str = "<ul>";
            for (i in v) {
                str += "<li><img src=\""+v[i]+"\" /></li>";
            }
            str += "</ul>";
            v = str;
        }
        html += one(k, v);
    }
    $("#visual").html(html);

}
