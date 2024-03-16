$(() => {

    $(document).ajaxSuccess((_, xhr) => {
        const json = JSON.parse(xhr.responseText)
        const expires = new Date(json.accessTokenExpires).toUTCString()
        document.cookie = `access_token=${json.accessToken};expires=${expires}`
        window.location.href = '/'
    })

});


