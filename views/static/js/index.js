$().ready(() => {


    layui.use('element', function () {
        let element = layui.element;
        let $iframe = document.getElementById('iframe');
        src_arr = [
            "page/file_list.html",
            "page/upload.html",
            "page/file_list_depth.html"
        ]
        element.on('tab(docDemoTabBrief)', function (data) {
            // console.log(data);
            let index = data.index
            $iframe.src = src_arr[index]
        });
    })


})




