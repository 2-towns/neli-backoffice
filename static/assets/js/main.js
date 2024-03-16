const alert = document.getElementById('alert')
const closeAlert = document.getElementById('alert-close')

closeAlert.onclick = () => {
    alert.style.display = 'none'
}

let timeout;

function logout() {
    document.cookie = 'access_token=; Max-Age=0'
    window.location.href = '/login'
}

function showSuccess() {
    if (timeout) {
        clearTimeout(timeout)
    }

    alert.style.display = 'block'
    alert.classList.remove('alert-danger')
    alert.classList.add('alert-success')
    const content = document.getElementById('alert-content')
    content.innerHTML = 'Done successfully'

    timeout = setTimeout(() => {
        alert.style.display = 'none'
    }, 3000)
}

function showError(response) {
    if (timeout) {
        clearTimeout(timeout)
    }

    alert.classList.remove('alert-success')
    alert.classList.add('alert-danger')
    alert.style.display = 'block'
    const content = document.getElementById('alert-content')
    content.innerHTML = JSON.parse(response).message

    timeout = setTimeout(() => {
        alert.style.display = 'none'
    }, 3000)
}

$(() => {

    const forms = document.getElementsByTagName('form')

    setInterval(() => {
        const element = `; ${document.cookie}`.match(`;\\s*access_token=([^;]+)`)
        if (window.location.pathname != "/login" && !element) {
            window.location.href = "/login"
        }
    }, 5000);

    Array.prototype.forEach.call(forms, (form) => {
        form.onsubmit = (event) => {
            event.preventDefault()
            sendAjax(form)
        }
    })

    function toggleButtons(disabled) {
        const buttons = document.getElementsByClassName('btn-primary')
        Array.prototype.forEach.call(buttons, (button) => {
            disabled ? button.setAttribute('disabled', 'disabled') : button.removeAttribute('disabled')
        })
    }

    function sendAjax(form) {
        toggleButtons(true)

        const url = form.getAttribute('data-url')
        const method = form.getAttribute('method')
        const tags = form.getElementsByTagName('input'), params = {}

        Array.prototype.forEach.call(tags, (tag) => {
            params[tag.name] = tag.type === 'number' ? parseInt(tag.value) : tag.value
        })

        console.log(params)

        $.ajax(
            {
                'url': url,
                'method': method,
                'data': JSON.stringify(params),
                'dataType': 'json',
                'success': () => {
                    showSuccess()
                },
                'error': (xhr) => {
                    showError(xhr.responseText)
                },
                'crossDomain': true,
                'beforeSend': (xhr) => {
                    const element = `; ${document.cookie}`.match(`;\\s*access_token=([^;]+)`)
                    if (element) {
                        access_token = element[1]
                        xhr.setRequestHeader('Authorization', `Bearer ${access_token}`);
                    }
                },
                'complete': () => {
                    toggleButtons(false)
                }
            }
        )
    }
})