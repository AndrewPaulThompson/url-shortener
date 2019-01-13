import '../css/app.pcss';

document.addEventListener("DOMContentLoaded", function(){
  const button = document.querySelector('.form__button')
  const container = document.querySelector('.form__container')
  const form = container.querySelector('form')

  if (button != 'undefined') {
    button.addEventListener('click', e => {
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
        const successMessage = document.querySelector('.form__success')
        const title = document.querySelector('.form__title').querySelector('h1')

        successMessage.innerHTML = `Your shortened url for <a href=${data.longUrl}>${data.longUrl}</a> is <a href=${data.shortUrl}>${data.shortUrl}</a>`
        successMessage.style.display = "flex"
        
        title.innerHTML = "Success"
        
        container.style.display = "none"
      })
    })
  }
})
