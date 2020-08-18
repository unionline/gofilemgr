let app = window.App = {}


app.SetPathToLocalStorage = (path) => {
    localStorage.setItem("app.gofilemgr.path", path)
}

app.GetPathFromLocalStorage = () => {
    let path = localStorage.getItem("app.gofilemgr.path")
    if (path == "undefined") {
        return ''
    }
    return path
}

app.SetPathValue = (path) => {
    App.SetPathToLocalStorage(path)
    $('#path').val(path)
}

app.ClearPathValue = () => {
    $('#path').val('')
    App.SetPathToLocalStorage('')
}

// 获取路径Path
app.GetPathValue = () => {
    return $('#path').val()
}

// 获取create_dir
app.GetFlagCreateDir = () => {
    return $('#create_dir').val()
}


app.CheckPathValid = () => {
    let path = App.GetPathValue()
    if (path == "") {
        return true
    }

    if (App.os == "windows") {

        // if (path[path.length - 1] == "\\") {
        //     path = path + '\\'
        // } else if  (path[path.length - 1] == "/") {
        //     path = path + '/'
        // }


        // 正则表达式
        //// Windows 一下三种都支持
        // 右斜杠一条例如：D:\xxx\zzz3\
        reg_win_right1 = /^[a-zA-Z]:(((\\(?! )[^/:*?<>\\""|\\]+)+\\?)|(\\)?)\s*$/
        // 右斜杠两条条例如：D:\\xxx\\zzz3\\
        reg_win_right2 = /^[a-zA-Z]:(((\\\\(?! )[^/:*?<>\\""|\\\\]+)+\\\\?)|(\\\\)?)\\s*$/
        // 左斜杠一条例如：D:/xxx/zzz3/
        reg_win = /^[a-zA-Z]:(((\/(?! )[^/:*?<>\/""|\/]+)+\/?)|(\/)?)\/s*$/


        if (reg_win_right1.test(path) || reg_win.test(path) || reg_win_right2.test(path)) {
            ok = true
        } else {
            ok = false
        }

    } else {
        if (path[path.length - 1] != "/") {
            path = path + '/'
        }

        //// Linux
        // 左斜杠一条例如：/home/fanbi/xxx/zzz3/
        reg_linux = /^(((\/home\/fanbi\/(?! )[^/\/""|\/]+)+\/?)|(\/)?)\/s*$/

        ok = reg_linux.test(path)
    }

    App.SetPathValue(path)

    return ok
}


// 获取路径Path
app.GetPathDepth = () => {
    return $('#depth').val()
}


// init
app.SetPathValue(app.GetPathFromLocalStorage())
// app.GetFileList()
