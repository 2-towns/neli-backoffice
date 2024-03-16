$(() => {
    const url = document.getElementById('table').getAttribute('data-url')
    const access_token = `; ${document.cookie}`.match(`;\\s*access_token=([^;]+)`)[1]

    // Global to access from ajax success
    let email;

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
            { 'data': 'userId' },
            { 'data': 'email' },
            { 'data': 'firstName' },
            { 'data': 'lastName' },
            {
                'data': null,
                'defaultContent': '<a class="edit mr-2"><i class="fas fa-pencil-alt"></i></a>' +
                    '<a class="remove"><i class="fas fa-trash"></i></a>'
            }
        ],
        "drawCallback": (settings) => {
            const edits = document.getElementsByClassName('edit')
            Array.prototype.forEach.call(edits, (element) => { edit(element) })

            const removes = document.getElementsByClassName('remove')
            Array.prototype.forEach.call(removes, (element) => { remove(element) })
        }
    });

    function edit(element) {
        element.onclick = (event) => {
            const lastname = element.parentElement.previousSibling
            document.getElementById('lastname').value = lastname.innerHTML

            const firstname = lastname.previousSibling
            document.getElementById('firstname').value = firstname.innerHTML

            email = firstname.previousSibling
            document.getElementById('email').value = email.innerHTML

            const userId = email.previousSibling.innerHTML
            const form = document.getElementById('edit-form')
            const url = form.getAttribute('url')
            form.setAttribute('data-url', url.replace('user_id', userId))

            $('#edit-modal').modal()
        }
    }

    function remove(element) {
        element.onclick = (event) => {
            const userId = element.parentElement.previousSibling.previousSibling.previousSibling.previousSibling.innerHTML
            const form = document.getElementById('remove-form')
            const url = form.getAttribute('url')
            form.setAttribute('data-url', url.replace('user_id', userId))
            $('#remove-modal').modal()
        }
    }

    $(document).ajaxSuccess((event, xhr, settings) => {
        if (settings.type !== "GET") {
            table.ajax.reload();
        }

        if (settings.type !== "POST") {
            document.getElementById('add-form').reset()
        }

    });

    $(document).ajaxComplete(() => { $('.modal').modal('hide') });

})    