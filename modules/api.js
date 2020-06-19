let baseURL = 'https://apifactory.telkom.co.id:8243/'

const fetch = require('node-fetch')

exports.send = async (method, url , data, addHeader=null) => {
    let headers = {}
    if (addHeader) {
        for (const key in addHeader) {
            const element = addHeader[key];
            headers[key] = element
        }
    }
    try {
        let response = await fetch(baseURL + url, {
            method,
            body: method == 'POST' ? JSON.stringify(data) : null,
            headers,
        })

        var data = await response.text()
        data = JSON.parse(data)

        return data
    } catch (error) {
        throw error.message
    }
}


