import '../css/app.pcss';

document.addEventListener("DOMContentLoaded", function(){
  const button = document.querySelector('.form__button')
  const container = document.querySelector('.form__container')
  const input = document.querySelector('.form__input-url')
  const form = container.querySelector('form')

  if (button != 'undefined') {
    button.addEventListener('click', send)
  }

  /**
   * Send formdata to /create endpoint
   * @param {event} e 
   */
  function send(e) {
    e.preventDefault()
    let formData = new FormData(form)
    
    fetch("/create", {
        body: formData,
        method: "POST"
    })
    .then(data => {
      return data.json()
    })
    .then(data => {
      input.value = data.shortUrl
      button.innerHTML = "Copy"
      button.removeEventListener('click', send)
      button.addEventListener('click', copy)
    })
  }

  /**
   * Copy input text to clipboard
   * @param {event} e 
   */
  function copy(e) {
    e.preventDefault()
    input.select()
    document.execCommand("copy");
  }
})
