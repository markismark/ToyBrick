;(function() {
    var tb = window.ToyBrick ={};
    
    var buildParams = function(params) {
        if (typeof params == "undefined"){
            return "";
        }
        if (typeof params == "string"){
            return params;
        }
        var result = [];
        for(var key in params){
            result.push(encodeURI(key) + "=" + params[key])
        }
        return result.join("&");
    }
    
    tb.Ajax = function(requests) {
        var length = requests.length;
        console.log(requests);
        var params = [];
        for (var index = 0; index < length; index++) {
            var request = requests[index];
            request.method = typeof request.method == "undefined" || request.method.trim() == "" ? "GET" : request.method.toUpperCase();
            var data = buildParams(request.data);
            if (request.method == "GET" && data != ""){
                if(request.url.indexOf("?") > 0){
                    request.url += "&" + data;
                }else{
                    request.url += "?" + data;
                }
                data = ""; 
            }
            params.push({
                "method" : request.method.toUpperCase(),
                "url" : request.url,
                "data" : data
            })
        }
        console.log(params);
    }
    
    // var test = [
    //     {
    //         "url" : "http://www.baidu.com/",
    //         "method" : "get",
    //         "data" : {
    //             "name" : "x",
    //             "yu" : "y"
    //         }
    //     },
    //     {
    //         "url" : "http://www.baidu.com/?dysue=dsue",
    //         "method" : "get",
    //         "data" : {
    //             "name" : "x",
    //             "yu" : "y"
    //         }
    //     },
    //     {
    //         "url" : "http://www.baidu.com/?dysue=dsue",
    //         "data" : {
    //             "name" : "x",
    //             "yu" : "y"
    //         }
    //     },
    //     {
    //         "url" : "http://www.baidu.com/?dysue=dsue",
    //         "method" : "post",
    //         "data" : {
    //             "name" : "x",
    //             "yu" : "y"
    //         }
    //     },
        
    // ]
    // tb.Ajax(test);
      
})();