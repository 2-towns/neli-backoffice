$(() => {
    const url = document.getElementById('table').getAttribute('data-url')
    const access_token = `; ${document.cookie}`.match(`;\\s*access_token=([^;]+)`)[1]

    const table = $('table').DataTable({
        responsive: true,
        'ajax': {
            'url': url,
            'dataSrc': '',
            'type': 'GET',
            'beforeSend': (xhr) => {
                xhr.setRequestHeader('Authorization', `Bearer ${access_token}`);
            },
        },
        'columns': [
            { 'data': 'leaderId' },
            { 'data': 'videoContentId' },
            { 'data': 'name' },
            { 'data': 'description' },
            { 'data': 'creationDate' },
            { 'data': 'duration' },
            {
                'data': 'sharingStatus',
                'render': (data, type, full, meta) => {
                    const disabled = data ? '' : 'disabled'
                    return `
                    <a id="share-${full.videoContentId}" class="share ${disabled}" data-id="${full.videoContentId}">
                        <i id="share-icon-${full.videoContentId}" class="fas fa-share-alt"></i>
                    </a>
                    `
                }
            },
        ],
        "drawCallback": (settings) => {
            const shares = document.getElementsByClassName('share')
            Array.prototype.forEach.call(shares, (element) => {
                element.onclick = () => {
                    if (!element.classList.contains('disabled')) {
                        const id = element.getAttribute('data-id')

                        const icon = document.getElementById(`share-icon-${id}`)
                        icon.classList.add('fa-spin')

                        const table = document.getElementById('table')
                        const url = table.getAttribute('data-share-url').replace('{videoContentId}', id)

                        callShares(url, id)
                    }
                }
            })
        }
    });


    function callShares(url, id) {
        $.ajax(
            {
                'url': url,
                'method': 'GET',
                'success': (data) => {
                    const list = document.getElementById('modal-list')
                    list.innerHTML = ''
                    let element

                    if (data) {
                        data.forEach((share) => {
                            const date = new Date(share.expirationDate * 1000)
                            let month = date.getMonth() + 1
                            if (month < 10) {
                                month = '0' + month
                            }
                            let day = date.getDate();
                            if (day < 10) {
                                day = '0' + day
                            }
                            const dateString = `${date.getFullYear()}-${month}-${day}`
                            const hours = date.getHours() < 10 ? '0' + date.getHours() : date.getHours()
                            const minutes = date.getMinutes() < 10 ? '0' + date.getMinutes() : date.getMinutes()
                            const seconds = date.getSeconds() < 10 ? '0' + date.getSeconds() : date.getSeconds()
                            const time = `${hours}:${minutes}:${seconds}`

                            element = `
                        <li class="list-group-item d-flex justify-content-between align-items-center">
                            ${share.name}
                            <span class="badge badge-primary badge-pill">Expiration: ${dateString} ${time}</span>
                        </li>
                        `
                            list.innerHTML += element
                        })
                    } else {
                        list.innerHTML += "<p style=\"padding: 10px 13px;\">This share doesn't contain any member</p>"
                    }
                    $('#modal').modal()
                },
                'error': (xhr) => { showError(xhr.responseText) },
                'crossDomain': true,
                'beforeSend': (xhr) => { xhr.setRequestHeader('Authorization', `Bearer ${access_token}`) },
                'complete': () => {
                    const icon = document.getElementById(`share-icon-${id}`)
                    icon.classList.remove('fa-spin')
                }
            })
    }
})    