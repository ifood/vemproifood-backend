const express = require('express')

const weatherController = require('./controllers/appController')

const routes = express.Router()

routes.get('/', (req, res) => {
    return res.json({ message: 'Hello Leo!' })
})

routes.get('/playlist', weatherController.getByCityName)

module.exports = routes
