$(() => {
    const url = document.getElementById('body').getAttribute('data-url')
    const access_token = `; ${document.cookie}`.match(`;\\s*access_token=([^;]+)`)[1]

    $.get({
        'url': url,
        'beforeSend': (xhr) => {
            xhr.setRequestHeader('Authorization', `Bearer ${access_token}`)
        }
    }).done((data) => {
        document.getElementById("maxDuration").value = data.maxDuration
    })
})