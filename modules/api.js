const fetch = require('node-fetch')

exports.send = async (method, url , data, addHeader=null) => {
    let headers = {}
    if (addHeader) {
        for (const key in addHeader) {
            const element = addHeader[key];
            headers[key] = element
        }
    }
    
    let response = await fetch(url, {
        method,
        body: method == 'POST' ? JSON.stringify(data) : null,
        headers,
    })

    var data = await response.text()
    data = JSON.parse(data)

    return data
}



